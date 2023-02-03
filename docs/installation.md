# Installation

I use this repository to experiment with
[goreleaser](https://goreleaser.com/), so there are a ridiculous number
of builds and packages available for such a trivial app. Don't judge.

=== "Homebrew"

    I maintain a [Homebrew](https://brew.sh/) tap.

    ```bash
    brew install mcornick/tap/mpg
    ```

=== "Scoop"

    I maintain a [Scoop](https://scoop.sh/) bucket.

    ```powershell
    scoop bucket add mcornick https://github.com/mcornick/scoop-bucket.git
    scoop install mpg
    ```

=== "Container Image"

    I maintain container images on GitHub [here](https://github.com/mcornick/mpg/pkgs/container/mpg).

    ```bash
    docker run --rm ghcr.io/mcornick/mpg
    ```

=== "Binaries and RPM/DEB/APK packages"

    I maintain binary releases on GitHub [here](https://github.com/mcornick/mpg/releases).

=== "Signatures"

    !!! Note

        Signatures are an ongoing experiment. They might be broken.
        Do not depend on them (or, really, on mpg) for mission-critical things.

    Binary checksums and container manifests are signed with
    [Cosign](https://docs.sigstore.dev/cosign/overview/).
    The certificate needed to verify binary checksums can be found in each release.
