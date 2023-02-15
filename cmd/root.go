// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Mark Cornick
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/mcornick/linenoise"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	length  int
	upper   bool
	lower   bool
	digit   bool
	version = "dev"
	rootCmd = &cobra.Command{
		Version: version,
		Use:     "mpg",
		Short:   "Mark's Password Generator",
		Long: `mpg is a command-line tool that generates strings of random characters
that can be used as reasonably secure passwords.`,
		Run: func(cmd *cobra.Command, args []string) {
			p := linenoise.Parameters{
				Length: viper.GetInt("length"),
				Upper:  viper.GetBool("upper"),
				Lower:  viper.GetBool("lower"),
				Digit:  viper.GetBool("digit"),
			}
			if !p.Upper && !p.Lower && !p.Digit {
				cobra.CheckErr(fmt.Errorf("no character classes were selected"))
			}
			minimumLength := btoi(p.Upper) + btoi(p.Lower) + btoi(p.Digit)
			if p.Length < minimumLength {
				cobra.CheckErr(fmt.Errorf("length must be at least %d", minimumLength))
			}
			result, err := linenoise.Noise(p)
			cobra.CheckErr(err)
			fmt.Println(result)
		},
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func init() {
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/mpg/")
		viper.AddConfigPath(configDir + "/mpg")
	}
	viper.SetEnvPrefix("mpg")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			cobra.CheckErr(err)
		}
	}
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "length")
	rootCmd.Flags().BoolVarP(&upper, "upper", "U", true, "include uppercase")
	rootCmd.Flags().BoolVarP(&lower, "lower", "u", true, "include lowercase")
	rootCmd.Flags().BoolVarP(&digit, "digit", "d", true, "include digits")
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "",
		fmt.Sprintf("config file (default %s/mpg/config.yaml)", configDir))
	for _, key := range []string{"length", "upper", "lower", "digit"} {
		cobra.CheckErr(viper.BindPFlag(key, rootCmd.Flags().Lookup(key)))
	}
}
