// Package main provides a CLI tool to set git user credentials in a particular directory.
//
// The tool accepts an optional command line argument "suffix".
// If provided, the suffix is converted to uppercase and used to form the environment variable names
// for the git name and email (e.g., "GIT_NAME_SUFFIX" and "GIT_EMAIL_SUFFIX").
// If the suffix is not provided, "GIT_NAME" and "GIT_EMAIL" are used as the environment variable names.
// The values of these environment variables are then used to set the git name and email.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	suffix := flag.String("suffix", "", "Suffix for the git name and email environment variables")
	flag.Parse()

	if flag.NArg() > 0 {
		*suffix = flag.Arg(0)
	}

	nameEnvVar := "GIT_NAME"
	emailEnvVar := "GIT_EMAIL"

	if *suffix != "" {
		upperSuffix := strings.ToUpper(*suffix)
		nameEnvVar += "_" + upperSuffix
		emailEnvVar += "_" + upperSuffix
	}

	name := os.Getenv(nameEnvVar)
	email := os.Getenv(emailEnvVar)

	if name == "" || email == "" {
		fmt.Printf("Name or email not set in environment variables: %s, %s\n", nameEnvVar, emailEnvVar)
		os.Exit(1)
	}

	exec.Command("git", "config", "user.name", name).Run()
	exec.Command("git", "config", "user.email", email).Run()

	fmt.Printf("Git name set to: %s\n", name)
	fmt.Printf("Git email set to: %s\n", email)
}
