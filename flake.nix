{
  description = "SYOI Online Judge";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      gomod2nix,
      treefmt-nix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        treefmtEval = treefmt-nix.lib.evalModule nixpkgs.legacyPackages.${system} ./treefmt.nix;
      in
      {
        packages.default = gomod2nix.legacyPackages.${system}.buildGoApplication {
          pname = "judy";
          name = "judy";
          src = ./.;
          modules = ./gomod2nix.toml;
        };
        devShells.default = pkgs.mkShell {
          packages =
            with pkgs;
            [
              air
              go
              cobra-cli
              ent-go
              isolate
              just
              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
            ]
            ++ [ gomod2nix.packages.${system}.default ];
        };
        formatter = treefmtEval.config.build.wrapper;
      }
    );
}
