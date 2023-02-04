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
	"log"
	"os"

	"github.com/mcornick/linenoise"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	length  int
	upper   bool
	lower   bool
	digit   bool
	version = "dev"
	rootCmd = &cobra.Command{
		Version: version,
		Use:     "mpg",
		Short:   "A password generator",
		Long: `mpg is a command-line tool that generates strings of
random characters that can be used as reasonably secure passwords.`,
		Run: func(cmd *cobra.Command, args []string) {
			p := linenoise.Parameters{
				Length: viper.GetInt("length"),
				Upper:  viper.GetBool("upper"),
				Lower:  viper.GetBool("lower"),
				Digit:  viper.GetBool("digit"),
			}
			result, err := linenoise.Noise(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/mpg/")
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(configDir + "/mpg")
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		}
	}
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "length")
	rootCmd.Flags().BoolVarP(&upper, "upper", "U", true, "include uppercase")
	rootCmd.Flags().BoolVarP(&lower, "lower", "u", true, "include lowercase")
	rootCmd.Flags().BoolVarP(&digit, "digit", "d", true, "include digits")
	viper.SetEnvPrefix("mpg")
	err = viper.BindEnv("length")
	if err != nil {
		panic(err)
	}
	err = viper.BindEnv("upper")
	if err != nil {
		panic(err)
	}
	err = viper.BindEnv("lower")
	if err != nil {
		panic(err)
	}
	err = viper.BindEnv("digit")
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("length", rootCmd.Flags().Lookup("length"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("upper", rootCmd.Flags().Lookup("upper"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("lower", rootCmd.Flags().Lookup("lower"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("digit", rootCmd.Flags().Lookup("digit"))
	if err != nil {
		panic(err)
	}
}
