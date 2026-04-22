package main

// BuildKit secret ids (docker buildx: --secret id=<id>,src=...)
const (
	secretIDNetrc           = "netrc"
	secretIDNixAccessTokens = "nix_access_tokens"
)

// Mount paths inside the Nix run container (must match llb.AddSecret targets).
const (
	pathNixNetrc = "/etc/nix/.netrc"
	// RHS of nix.conf "access-tokens = …" (e.g. github.com=ghp_…), one line, space-separated hosts if needed.
	// Single path under /run (no subdirs) so the container always has a valid parent.
	pathNixAccessTokensFile = "/run/nix_access_tokens"
)
