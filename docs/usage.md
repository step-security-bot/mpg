# Usage

## Command Line

```bash
$ ./mpg --help
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
suffix; use the command line options as keys; and be located in one
of these directories:

- `$HOME/Library/Application Support/mpg` (macOS)
- `$XDG_CONFIG_HOME/mpg` (Unix)
- `/etc/mpg` (macOS or Unix)
- `%AppData%\mpg` (Windows)

For example, these files replicate the default configuration of mpg:

=== "config.yaml"

    ```yaml
    ---
    length: 16
    upper: true
    lower: true
    digit: true
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
