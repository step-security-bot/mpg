---
title: mpg
---
__mpg__ is a command-line tool that generates strings of random characters that can be used as reasonably secure passwords.

Passwords are, by default, chosen from the union of three character classes: upper-case letters, lower-case letters, and digits.

Options can be given to omit any one or two of these character classes. For instance, you can omit uppercase letters and digits with `mpg --upper=false --digit=false`. This will return a password composed only of lowercase letters.

Passwords are guaranteed to contain at least one of each selected character class. The default length is 16. __mpg__ will create a password of any length greater than or equal to the number of character classes selected.

If the password length is less than or equal to the total number of characters selected, __mpg__ will not repeat characters within the generated password.

## Installation

I use this repository to experiment with [goreleaser](https://goreleaser.com/), so there are a ridiculous number of builds and packages available for such a trivial app. Don't judge.

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

        Signatures are an ongoing experiment. They might be broken. Do not depend on them (or, really, on mpg) for mission-critical things.

    Binary checksums and container manifests are signed with [Cosign](https://docs.sigstore.dev/cosign/overview/). The certificate needed to verify binary checksums can be found in each release.

    The signatures are created with Cosign's "keyless" mode, which requires setting `COSIGN_EXPERIMENTAL=1` when using Cosign versions prior to 2.0.0:

    ```bash
    env COSIGN_EXPERIMENTAL=1 cosign verify ghcr.io/mcornick/mpg
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

Values given on the command line take precedence over values given in the configuration file or in environment variables.

## Contributing

Bug reports and pull requests are welcome on [GitHub](https://github.com/mcornick/mpg).

This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](https://www.contributor-covenant.org/) code of conduct.

__mpg__ is available as open source under the terms of the MIT License.
