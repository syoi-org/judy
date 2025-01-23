{ ... }:

{
  projectRootFile = "flake.nix";

  programs.gofmt.enable = true;
  programs.mdformat = {
    enable = true;
    settings.number = true;
  };
  programs.nixfmt.enable = true;
  programs.protolint.enable = true;
  programs.yamlfmt.enable = true;
}
