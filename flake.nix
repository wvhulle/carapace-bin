{
  description = "Multi-shell multi-command argument completer";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forEachSystem = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      packages = forEachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          carapacePackage = pkgs.buildGoModule (finalAttrs: {
            pname = "carapace";
            version = "1.8.1-${self.shortRev or self.dirtyShortRev or "dirty"}";

            src = self;

            vendorHash = "sha256-ZFpssZN8htDmiW2vYgRrK5cT44hMPNLaNmGwZEwNtFw=";

            proxyVendor = true;

            ldflags = [
              "-s"
              "-w"
              "-X main.version=${finalAttrs.version}"
            ];

            subPackages = [ "./cmd/carapace" ];

            tags = [ "release" ];

            preBuild = ''
              GOOS= GOARCH= go generate ./...
            '';

            meta = {
              description = "Multi-shell multi-command argument completer";
              homepage = "https://carapace.sh/";
              license = nixpkgs.lib.licenses.mit;
              mainProgram = "carapace";
            };
          });
        in
        {
          carapace = carapacePackage;
          default = carapacePackage;
        }
      );

      checks = forEachSystem (system: {
        carapace = self.packages.${system}.carapace;
      });
    };
}
