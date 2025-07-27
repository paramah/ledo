{
  description = "Ledo flake";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs, ... }: 
    let
      system = builtins.currentSystem or (if builtins ? currentSystem then builtins.currentSystem else "aarch64-darwin");

      pkgs = import nixpkgs { inherit system; };

      goVersion = "1.24"; # zmień jeśli chcesz inną wersję

      goPkgs = {
        "1.24" = pkgs.go_1_24;
        "1.23" = pkgs.go_1_23;
        "1.22" = pkgs.go_1_22;
        "1.21" = pkgs.go_1_21;
      };

      selectedGo = if goPkgs ? ${goVersion}
        then goPkgs.${goVersion}
        else pkgs.go;

      printInfo = ''
        echo -e "\033[32m🚀 LeDo | Using Go version ${goVersion}\033[0m"
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
          echo -e "\033[34m🛠️  Environment: DEVELOPMENT\033[0m"
        '';
      };

      packages.${system}.default = pkgs.buildGoModule {
        pname = "ledo";
        version = "0.1.0";

        src = ./.;

        vendorHash = "sha256-Xn7icXrEKQuJAGiSyReYGoNdPAsIziEq1KHvXc6HEPU="; 

        modVendor = true;

        
        buildFlags = [ "-v -mod=vendor" ];

        ldflags = [
          "-s" "-w"
          "-X main.version=0.1.0"
        ];

        meta = {
          description = "Ledo built via Flake";
          license = pkgs.lib.licenses.mit;
        };
      };
    };
}

