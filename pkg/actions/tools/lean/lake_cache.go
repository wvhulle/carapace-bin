package lean

import (
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	toml "github.com/pelletier/go-toml"
)

type lakeSystemConfig struct {
	Cache struct {
		Services []struct {
			Name string `toml:"name"`
			Kind string `toml:"kind"`
		} `toml:"service"`
	} `toml:"cache"`
}

// ActionLakeCacheServices completes remote cache services of the Lake system configuration
// Reads `LAKE_CONFIG` or `~/.lake/config.toml`.
//
//	reservoir (default)
//	my-s3 (s3)
func ActionLakeCacheServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		values := []string{"reservoir", "default"}
		configPath := c.Getenv("LAKE_CONFIG")
		if configPath == "" {
			if home, err := os.UserHomeDir(); err == nil {
				configPath = filepath.Join(home, ".lake", "config.toml")
			}
		}
		if content, err := os.ReadFile(configPath); err == nil {
			var config lakeSystemConfig
			if err := toml.Unmarshal(content, &config); err == nil {
				for _, service := range config.Cache.Services {
					values = append(values, service.Name, service.Kind)
				}
			}
		}
		return carapace.ActionValuesDescribed(values...)
	}).Tag("cache services")
}
