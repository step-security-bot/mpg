# Installation

I use this repository to experiment with
[goreleaser](https://goreleaser.com/), so there are a ridiculous number
of builds and packages available for such a trivial app. Don't judge.

## Homebrew
```bash
brew install mcornick/tap/mpg
```

## Scoop
```bash
scoop bucket add mcornick https://github.com/mcornick/scoop-bucket.git
scoop install mpg
```

## Docker
```bash
docker run --rm ghcr.io/mcornick/mpg
```

## Binaries and RPM/DEB/APK packages

[Here.]

## Signatures

Binary checksums and container manifests are signed with
[Cosign](https://docs.sigstore.dev/cosign/overview/).
The certificate needed to verify signatures can be found in each release.
