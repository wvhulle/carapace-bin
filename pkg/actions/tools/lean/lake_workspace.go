package lean

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	toml "github.com/pelletier/go-toml"
)

type lakeTargetKind string

const (
	lakeLibraryKind    lakeTargetKind = "library"
	lakeExecutableKind lakeTargetKind = "executable"
	lakeExternLibKind  lakeTargetKind = "extern library"
	lakeCustomKind     lakeTargetKind = "target"
	lakeInputFileKind  lakeTargetKind = "input file"
	lakeInputDirKind   lakeTargetKind = "input directory"
	lakeScriptKind     lakeTargetKind = "script"
)

type lakeTarget struct {
	Name string
	Kind lakeTargetKind
}

type lakePackageConfig struct {
	Name    string
	Targets []lakeTarget
}

func (config lakePackageConfig) targetsOfKind(kinds ...lakeTargetKind) []string {
	names := make([]string, 0)
	for _, target := range config.Targets {
		for _, kind := range kinds {
			if target.Kind == kind {
				names = append(names, target.Name)
			}
		}
	}
	return names
}

type lakeManifestPackage struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	URL      string  `json:"url"`
	Dir      string  `json:"dir"`
	InputRev *string `json:"inputRev"`
}

type lakeManifest struct {
	PackagesDir string                `json:"packagesDir"`
	Packages    []lakeManifestPackage `json:"packages"`
}

type lakeWorkspace struct {
	Dir      string
	Root     lakePackageConfig
	Manifest lakeManifest
}

func findLakeWorkspaceDir(dir string) (string, error) {
	current, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	for {
		for _, configName := range []string{"lakefile.lean", "lakefile.toml"} {
			if _, err := os.Stat(filepath.Join(current, configName)); err == nil {
				return current, nil
			}
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", errors.New("could not find lakefile.lean or lakefile.toml")
		}
		current = parent
	}
}

var lakeLeanPackagePattern = regexp.MustCompile(`(?m)^\s*package\s+([^\s(:]+)`)
var lakeLeanTargetPattern = regexp.MustCompile(`(?m)^\s*(?:@\[[^\]]*\]\s*)?(lean_lib|lean_exe|extern_lib|target|input_file|input_dir|script)\s+([^\s(:]+)`)

var lakeLeanDeclarationKinds = map[string]lakeTargetKind{
	"extern_lib": lakeExternLibKind,
	"input_dir":  lakeInputDirKind,
	"input_file": lakeInputFileKind,
	"lean_exe":   lakeExecutableKind,
	"lean_lib":   lakeLibraryKind,
	"script":     lakeScriptKind,
	"target":     lakeCustomKind,
}

func unquoteLeanName(name string) string {
	name = strings.Trim(name, `"`)
	name = strings.ReplaceAll(name, "«", "")
	return strings.ReplaceAll(name, "»", "")
}

func parseLakefileLean(content string) lakePackageConfig {
	config := lakePackageConfig{}
	if match := lakeLeanPackagePattern.FindStringSubmatch(content); match != nil {
		config.Name = unquoteLeanName(match[1])
	}
	for _, match := range lakeLeanTargetPattern.FindAllStringSubmatch(content, -1) {
		config.Targets = append(config.Targets, lakeTarget{
			Name: unquoteLeanName(match[2]),
			Kind: lakeLeanDeclarationKinds[match[1]],
		})
	}
	return config
}

type lakeTomlTarget struct {
	Name string `toml:"name"`
}

type lakeTomlConfig struct {
	Name       string           `toml:"name"`
	LeanLibs   []lakeTomlTarget `toml:"lean_lib"`
	LeanExes   []lakeTomlTarget `toml:"lean_exe"`
	InputFiles []lakeTomlTarget `toml:"input_file"`
	InputDirs  []lakeTomlTarget `toml:"input_dir"`
}

func parseLakefileToml(content []byte) (lakePackageConfig, error) {
	var tomlConfig lakeTomlConfig
	if err := toml.Unmarshal(content, &tomlConfig); err != nil {
		return lakePackageConfig{}, err
	}
	config := lakePackageConfig{Name: tomlConfig.Name}
	for _, group := range []struct {
		targets []lakeTomlTarget
		kind    lakeTargetKind
	}{
		{tomlConfig.LeanLibs, lakeLibraryKind},
		{tomlConfig.LeanExes, lakeExecutableKind},
		{tomlConfig.InputFiles, lakeInputFileKind},
		{tomlConfig.InputDirs, lakeInputDirKind},
	} {
		for _, target := range group.targets {
			config.Targets = append(config.Targets, lakeTarget{Name: target.Name, Kind: group.kind})
		}
	}
	return config, nil
}

func loadLakePackageConfig(packageDir string) (lakePackageConfig, error) {
	if content, err := os.ReadFile(filepath.Join(packageDir, "lakefile.lean")); err == nil {
		return parseLakefileLean(string(content)), nil
	}
	content, err := os.ReadFile(filepath.Join(packageDir, "lakefile.toml"))
	if err != nil {
		return lakePackageConfig{}, err
	}
	return parseLakefileToml(content)
}

func loadLakeManifest(workspaceDir string) lakeManifest {
	manifest := lakeManifest{}
	if content, err := os.ReadFile(filepath.Join(workspaceDir, "lake-manifest.json")); err == nil {
		_ = json.Unmarshal(content, &manifest)
	}
	if manifest.PackagesDir == "" {
		manifest.PackagesDir = filepath.Join(".lake", "packages")
	}
	return manifest
}

func loadLakeWorkspace(dir string) (lakeWorkspace, error) {
	workspaceDir, err := findLakeWorkspaceDir(dir)
	if err != nil {
		return lakeWorkspace{}, err
	}
	root, err := loadLakePackageConfig(workspaceDir)
	if err != nil {
		return lakeWorkspace{}, err
	}
	return lakeWorkspace{
		Dir:      workspaceDir,
		Root:     root,
		Manifest: loadLakeManifest(workspaceDir),
	}, nil
}

func (workspace lakeWorkspace) dependencyDir(dependency lakeManifestPackage) string {
	if dependency.Type == "path" && dependency.Dir != "" {
		if filepath.IsAbs(dependency.Dir) {
			return dependency.Dir
		}
		return filepath.Join(workspace.Dir, dependency.Dir)
	}
	return filepath.Join(workspace.Dir, workspace.Manifest.PackagesDir, dependency.Name)
}

func (workspace lakeWorkspace) dependencyConfigs() map[string]lakePackageConfig {
	configs := make(map[string]lakePackageConfig)
	for _, dependency := range workspace.Manifest.Packages {
		if config, err := loadLakePackageConfig(workspace.dependencyDir(dependency)); err == nil {
			configs[dependency.Name] = config
		}
	}
	return configs
}

func actionWithLakeWorkspace(f func(workspace lakeWorkspace) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		workspace, err := loadLakeWorkspace(c.Dir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(workspace)
	})
}
