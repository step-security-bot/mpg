---
title: mpg
---
__mpg__ is a command-line tool that generates strings of random characters that can be used as reasonably secure passwords.

Passwords are, by default, chosen from the union of three character classes: upper-case letters, lower-case letters, and digits.

Options can be given to omit any one or two of these character classes. For instance, you can omit uppercase letters and digits with `mpg --upper=false --digit=false`. This will return a password composed only of lowercase letters.

Passwords are guaranteed to contain at least one of each selected character class. The default length is 16. __mpg__ will create a password of any length greater than or equal to the number of character classes selected.

If the password length is less than or equal to the total number of characters selected, __mpg__ will not repeat characters within the generated password. For instance, in the default configuration of __mpg__, 62 characters are selected (26 uppercase, 26 lowercase, and 10 digits) and so a length of 62 or less will not repeat any character.

## Installation

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

    I maintain container images on [GitHub](https://github.com/mcornick/mpg/pkgs/container/mpg) and [Docker Hub](https://hub.docker.com/repository/docker/mcornick/mpg).

    ```bash
    docker run --rm ghcr.io/mcornick/mpg
    # or, for Docker Hub
    docker run --rm mcornick/mpg
    ```

=== "Binaries and Linux packages"

    I maintain binary releases on GitHub [here](https://github.com/mcornick/mpg/releases). As of version 1.0.21, releases are built for macOS (universal), Linux (i386, amd64, arm64, and armv6) and Windows (i386, amd64). Linux packages are built in RPM, DEB, APK, and Arch Linux pkg.tar.zst formats.

=== "YUM Repository"

    RPM packages are also available from my Gemfury repository.

    ```
    # /etc/yum.repos.d/mcornick.repo
    [fury]
    name=mcornick yum repo
    baseurl=https://yum.fury.io/mcornick/
    enabled=1
    gpgcheck=0
    ```

=== "APT Repository"

    DEB packages are also available from my Gemfury repository.

    ```
    # /etc/apt/sources.list.d/mcornick.list
    deb [trusted=yes] https://apt.fury.io/mcornick/ /
    ```

=== "Signatures"

    Binary checksums are signed with my [GPG key](https://github.com/mcornick.gpg).

    !!! Note

        Container signatures are an ongoing experiment. They might be broken. Do not depend on them (or, really, on mpg) for mission-critical things.

    Container manifests are signed with [Cosign](https://docs.sigstore.dev/cosign/overview/). The signatures are created with Cosign's "keyless" mode, which requires setting `COSIGN_EXPERIMENTAL=1` when using Cosign versions prior to 2.0.0:

    ```bash
    env COSIGN_EXPERIMENTAL=1 cosign verify ghcr.io/mcornick/mpg
    # or, for Docker Hub
    env COSIGN_EXPERIMENTAL=1 cosign verify mcornick/mpg
    ```

## Usage

### Command Line

```bash
$ mpg --help
mpg is a command-line tool that generates strings of
random characters that can be used as reasonably secure passwords.

Usage:
  mpg [flags]
  mpg [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command
  man         Generate man pages

Flags:
  -d, --digit        include digits (default true)
  -h, --help         help for mpg
  -l, --length int   length (default 16)
  -u, --lower        include lowercase (default true)
  -U, --upper        include uppercase (default true)
  -v, --version      version for mpg

Use "mpg [command] --help" for more information about a command.
```

### Configuration File

__mpg__ supports an optional configuration file. It must be in a format supported by [viper](https://github.com/spf13/viper), such as YAML, TOML, or JSON; be named `config.XXXX` with the proper format suffix; use the command line options as keys; and be located in one of these directories:

- `$HOME/Library/Application Support/mpg` (macOS)
- `$XDG_CONFIG_HOME/mpg` (Unix)
- `/etc/mpg` (macOS or Unix)
- `%AppData%\mpg` (Windows)

For example, these files replicate the default configuration of __mpg__:

=== "config.yaml"

    ```yaml
    ---
    digit: true
    length: 16
    lower: true
    upper: true
    ```

=== "config.toml"

    ```toml
    digit = true
    length = 16
    lower = true
    upper = true
    ```

=== "config.json"

    ```json
    {
      "digit": true,
      "length": 16,
      "lower": true,
      "upper": true
    }
    ```

### Environment Variables

Configuration is also possible using environment variables. For example, these environment variables would replicate the default configuration of __mpg__:

```sh
export MPG_DIGIT=1
export MPG_LENGTH=16
export MPG_LOWER=1
export MPG_UPPER=1
```

(To disable uppercase, lowercase, or digits, set the appropriate variable to 0.)

The order of precedence for configuration is:

1. Values given on the command line
1. Values given as environment variables
1. Values in the configuration file

## Contributing to mpg

If you think you have a problem, improvement, or other contribution towards the betterment of __mpg__, please file an issue or, where appropriate, a pull request.

Keep in mind that I'm not paid to write Go code, so I'm doing this in my spare time, which means it might take me a while to respond.

When filing a pull request, please explain what you're changing and why. Please use standard Go formatting (`go fmt` is your friend.) Please limit your changes to the specific thing you're fixing and isolate your changes in a topic branch that I can merge without pulling in other stuff.

__mpg__ uses [Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md) style. Please follow this convention. Scopes are not required in commit messages.

__mpg__ uses the MIT license. Please indicate your acceptance of the MIT license by using [git commit --signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

__mpg__ is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!
