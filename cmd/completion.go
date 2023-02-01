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

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Long: fmt.Sprintln(`To load completions:

Bash:

  $ source <(mpg completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ mpg completion bash > /etc/bash_completion.d/mpg
  # macOS:
  $ mpg completion bash > $(brew --prefix)/etc/bash_completion.d/mpg

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ mpg completion zsh > "${fpath[1]}/_mpg"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ mpg completion fish | source

  # To load completions for each session, execute once:
  $ mpg completion fish > ~/.config/fish/completions/mpg.fish

PowerShell:

  PS> mpg completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> mpg completion powershell > mpg.ps1
  # and source this file from your PowerShell profile.`),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				panic(err)
			}
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			if err != nil {
				panic(err)
			}
		case "fish":
			err := cmd.Root().GenFishCompletion(os.Stdout, true)
			if err != nil {
				panic(err)
			}
		case "powershell":
			err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
