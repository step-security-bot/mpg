# Usage

## Command Line

```sh
mcornick@macbook-air mpg % ./mpg --help
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

## Configuration File

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
