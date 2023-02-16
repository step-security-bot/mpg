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
	"bytes"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

var (
	reUpper = regexp.MustCompile(`[[:upper:]]`)
	reLower = regexp.MustCompile(`[[:lower:]]`)
	reDigit = regexp.MustCompile(`[[:digit:]]`)
)

func executeWithArgs(args []string) string {
	testCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			result, err := generatePassword(cmd, args)
			if err != nil {
				cmd.Println(err)
			}
			cmd.Println(result)
		},
	}
	testCmd.Flags().IntVarP(&length, "length", "l", 16, "length")
	testCmd.Flags().BoolVarP(&upper, "upper", "U", true, "include uppercase")
	testCmd.Flags().BoolVarP(&lower, "lower", "u", true, "include lowercase")
	testCmd.Flags().BoolVarP(&digit, "digit", "d", true, "include digits")
	buff := bytes.NewBufferString("")
	testCmd.SetOut(buff)
	testCmd.SetArgs(args)
	err := testCmd.Execute()
	if err != nil {
		panic(err)
	}
	stdout, err := io.ReadAll(buff)
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(string(stdout), "\r\n")
}

func TestWithoutFlags(t *testing.T) {
	result := executeWithArgs([]string{})
	if len(result) != 16 {
		t.Fatalf("expected length of 16, got %d", len(result))
	}
	if !reUpper.MatchString(result) {
		t.Fatalf("expected at least one uppercase letter, got %s", result)
	}
	if !reLower.MatchString(result) {
		t.Fatalf("expected at least one lowercase letter, got %s", result)
	}
	if !reDigit.MatchString(result) {
		t.Fatalf("expected at least one digit, got %s", result)
	}
}

func TestWithLength(t *testing.T) {
	result := executeWithArgs([]string{"--length=12"})
	if len(result) != 12 {
		t.Fatalf("expected length of 12, got %d", len(result))
	}
}
func TestWithUpperFalse(t *testing.T) {
	result := executeWithArgs([]string{"--upper=false"})
	if reUpper.MatchString(result) {
		t.Fatalf("expected no uppercase letters, got %s", result)
	}
}
func TestWithLowerFalse(t *testing.T) {
	result := executeWithArgs([]string{"--lower=false"})
	if reLower.MatchString(result) {
		t.Fatalf("expected no lowercase letters, got %s", result)
	}
}
func TestWithDigitFalse(t *testing.T) {
	result := executeWithArgs([]string{"--digit=false"})
	if reDigit.MatchString(result) {
		t.Fatalf("expected no digits, got %s", result)
	}
}

func TestWithUpperFalseAndLowerFalse(t *testing.T) {
	result := executeWithArgs([]string{"--upper=false", "--lower=false"})
	if reUpper.MatchString(result) {
		t.Fatalf("expected no uppercase letters, got %s", result)
	}
	if reLower.MatchString(result) {
		t.Fatalf("expected no lowercase letters, got %s", result)
	}
}
func TestWithUpperFalseAndDigitFalse(t *testing.T) {
	result := executeWithArgs([]string{"--upper=false", "--digit=false"})
	if reUpper.MatchString(result) {
		t.Fatalf("expected no uppercase letters, got %s", result)
	}
	if reDigit.MatchString(result) {
		t.Fatalf("expected no digits, got %s", result)
	}
}
func TestWithLowerFalseAndDigitFalse(t *testing.T) {
	result := executeWithArgs([]string{"--lower=false", "--digit=false"})
	if reLower.MatchString(result) {
		t.Fatalf("expected no lowercase letters, got %s", result)
	}
	if reDigit.MatchString(result) {
		t.Fatalf("expected no digits, got %s", result)
	}
}

func TestWithUpperFalseAndLowerFalseAndDigitFalse(t *testing.T) {
	result := executeWithArgs([]string{"--upper=false", "--lower=false", "--digit=false"})
	expected := "no character classes were selected"
	if result != expected {
		t.Fatalf("expected Error: %s, got %s", expected, result)
	}
}
func TestWithShortLength(t *testing.T) {
	result := executeWithArgs([]string{"--length=0"})
	expected := "length must be at least 3"
	if result != expected {
		t.Fatalf("expected Error: %s, got %s", expected, result)
	}
}

func TestBtoITrue(t *testing.T) {
	result := btoi(true)
	if result != 1 {
		t.Fatalf("expected 1, got %d", result)
	}
}

func TestBtoIFalse(t *testing.T) {
	result := btoi(false)
	if result != 0 {
		t.Fatalf("expected 0, got %d", result)
	}
}
