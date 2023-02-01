# mpg

mpg is a command-line tool that generates strings of
random characters that can be used as reasonably secure passwords.

Passwords are, by default, chosen from the union of three character
classes: upper-case letters, lower-case letters, and digits.

Options can be given to omit any one or any two of these character
classes. For instance, you can omit uppercase letters and digits by
passing `--upper=false --digit=false` to `mpg`. This will
return a password composed only of lowercase letters.

Passwords are guaranteed to contain at least one of each selected
character class. The default length is 16. mpg will create
a password of any length greater than or equal to the number of
character classes selected.

If the password length is less than or equal to the total number of
characters selected, mpg will not repeat characters within
the generated password.

## Installation

Releases and container images are available at
https://github.com/mcornick/mpg

Binaries and DEB/RPM packages are signed with [this GPG
key](https://github.com/mcornick.gpg).
Containers and manifests are signed with [this Cosign
key](https://mcornick.github.io/cosign.pub).

Once the binary is built, `mpg completion [shell]` will generate
completion files for bash, zsh, fish, or powershell.

## Usage

### Command Line

```sh
mpg is a command-line tool that generates strings of
random characters that can be used as reasonably secure passwords.

Usage:
  mpg [flags]
  mpg [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command

Flags:
      --digit        include digits (default true)
  -h, --help         help for mpg
      --length int   length (default 16)
      --lower        include lowercase (default true)
      --upper        include uppercase (default true)

Use "mpg [command] --help" for more information about a command.

$ mpg
h6ECtbDZPnRddHV7
$ mpg --length=8
XdWod8f8
$ mpg --length=64
QhESpeyPDidxV9kFNCrJqeMa4XUYbET4B3s5oGA8kYsV6XwDKHrCL7wojGZm9gj5
$ mpg --length=0
2017/10/31 16:04:47 Invalid length - must be an integer greater than 2
$ mpg --lower=false
387HNFDEUW4YGMZA
$ mpg --upper=false --length=8
hcsym6tj
$ mpg --lower=false --upper=false --length=32
92992759356835354563826487673794
$ mpg --lower=false --upper=false --digit=false
2021/10/02 21:38:36 must include at least one of --lower, --upper and/or --digit
exit status 1
```

### Configuration File

mpg supports an optional configuration file. It must be in a
format supported by [viper](https://github.com/spf13/viper), such as
YAML, TOML, or JSON; be named `config.XXXX` with the proper format
suffix; be located in `$HOME/Library/Application Support/mpg`
(macOS) or `$XDG_CONFIG_HOME/mpg` (other Unix) or
`/etc/mpg`; and use the command line options as keys.

For example, this data in `config.yaml` would replicate the default
configuration of mpg:

```yaml
---
length: 16
upper: true
lower: true
digit: true
```

Configuration is also possible using environment variables. For example,
these environment variables would replicate the default configuration of
mpg:

```sh
export MPG_LENGTH=16
export MPG_UPPER=1
export MPG_LOWER=1
export MPG_DIGIT=1
```

(To disable uppercase, lowercase, or digits, set the appropriate
variable to 0.)

Values given on the command line take precedence over values given in
the configuration file or in environment variables.

## Contributing

This project is intended to be a safe, welcoming space for
collaboration, and contributors are expected to adhere to the
[Contributor Covenant](https://www.contributor-covenant.org/) code of
conduct.

mpg is available as open source under the terms of the MIT
License.
