with import <nixpkgs> {}; {
  goEnv = stdenv.mkDerivation {
    name = "gobetting";
    buildInputs = [ stdenv go openssl gcc6 ];
    shellHook =
      ''
        export GOPATH=~/go_workspace
      '';
  };
}
