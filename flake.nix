{
  description = "Basic environment for nixconfig-go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }@inputs:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.git
          ];
          
          shellHook = ''
            export GOPATH=$PWD/.gopath
            export PATH=$GOPATH/bin:$PATH
          '';
        };
      }
    );
}