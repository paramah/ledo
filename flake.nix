{
  description = "Golang full flake: devshell + build + release";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs, ... }: 
    let
      system = builtins.currentSystem;
      pkgs = import nixpkgs { inherit system; };

      goVersion = "1.22"; # zmień jeśli chcesz inną wersję

      goPkgs = {
        "1.22" = pkgs.go_1_22;
        "1.21" = pkgs.go_1_21;
        "1.20" = pkgs.go_1_20;
      };

      selectedGo = if goPkgs ? ${goVersion}
        then goPkgs.${goVersion}
        else pkgs.go;

      printInfo = ''
        echo -e "\033[32m🚀 Projekt Golang | Używasz Go ${goVersion}\033[0m"
      '';

    in
    {
      devShells.${system}.default = pkgs.mkShell {
        name = "golang-devshell";

        packages = with pkgs; [
          selectedGo
          gopls           # Go Language Server (dla IDE / autouzupełnianie)
          golangci-lint   # Linter
          delve           # Debugger
          git
          jq
        ];

        shellHook = ''
          ${printInfo}
          echo -e "\033[34m🛠️  Środowisko: DEVELOPMENT\033[0m"
        '';
      };

      packages.${system}.my-go-app = pkgs.buildGoModule {
        pname = "my-go-app";
        version = "0.1.0";

        src = ./.;

        vendorSha256 = null; # ustawi się automatycznie po pierwszym buildzie

        buildFlags = [ "-v" ];

        ldflags = [
          "-s" "-w"
          "-X main.version=0.1.0"
        ];

        meta = {
          description = "Example Go app built via Flake";
          license = pkgs.lib.licenses.mit;
        };
      };
    };
}

