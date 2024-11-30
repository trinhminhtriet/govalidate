package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/trinhminhtriet/govalidate/check"
)

type checker interface {
	Check() (ok bool, skip bool)
	Summary() string
	Resolution() string
}

var (
	ignoreCGO       bool
	ignoreEditors   bool
	ignorePathCheck bool
)

func main() {
	flag.Usage = printHelp
	flag.BoolVar(&ignoreCGO, "ignore-cgo", false, "")
	flag.BoolVar(&ignoreEditors, "ignore-editors", false, "")
	flag.BoolVar(&ignorePathCheck, "ignore-path", false, "")
	flag.Parse()

	var exit int
	checks := []checker{
		&check.GoChecker{}, // checks go and go version
	}
	// Skip path check.
	if ignorePathCheck == false {
		checks = append(checks, &check.PathChecker{}) // checks $GOPATH/bin is in $PATH
	}
	for _, c := range checks {
		exit += runCheck(false, c)
	}
	// Optional checks.
	var optionals []checker
	if !ignoreCGO {
		optionals = append(optionals, &check.CGOChecker{})
	}
	if !ignoreEditors {
		// TODO(jbd): Add Goland.
		optionals = append(optionals,
			&check.VimChecker{},
			&check.VSCodeChecker{},
		)
	}
	for _, c := range optionals {
		exit += runCheck(true, c)
	}

	if exit > 0 {
		os.Exit(1)
	}
}

func runCheck(optional bool, c checker) int {
	var exit int

	ok, skip := c.Check()
	if skip {
		return exit
	}
	if ok {
		color.New(color.FgHiGreen).Print("[✔]")
	} else {
		if !optional {
			exit = 1
			color.New(color.FgRed).Print("[✗]")
		} else {
			color.New(color.FgYellow).Print("[!]")
		}
	}
	fmt.Print(" ")
	fmt.Println(c.Summary())
	if !ok {
		printWithTabs(c.Resolution())
	}
	return exit
}

func printHelp() {
	fmt.Print(`govalidate:
Validates your Go installation and dependencies.

Usage:
  govalidate [option]

Options:
  -ignore-cgo       Ignore checking CGO support. 
  -ignore-editors   Ignore checking editor plugin support.
  -ignore-path      Ignore checking path.
  -help             Display this user guide.
`)
	os.Exit(1)
}

func printWithTabs(msg string) {
	lines := strings.Split(msg, "\n")
	for _, l := range lines {
		fmt.Printf("    %v\n", l)
	}
}
