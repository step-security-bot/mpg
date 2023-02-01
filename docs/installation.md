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

Containers and manifests are signed with [this Cosign
key](https://mcornick.github.io/cosign.pub).

## Binaries and RPM/DEB/APK packages
[Here.](https://github.com/mcornick/mpg/releases) Checksums from that
page are signed with [my GPG key](https://github.com/mcornick.gpg).
