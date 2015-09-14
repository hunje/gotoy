package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

var targetRepository string

const WHITE_SPACES string = " \t\n"

func init() {
	flag.StringVar(&targetRepository, "t", "", "target repository")
	flag.Parse()
}

func getCurrentBranchName() string {
	cmd := exec.Command("git", "get-branch-name")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return strings.TrimRight(out.String(), WHITE_SPACES)
}

func main() {
	if targetRepository == "" {
		flag.PrintDefaults()
		return
	}

	currentBranch := getCurrentBranchName()

	if len(currentBranch) <= 0 {
		fmt.Println("here is not under git")
		return
	}

	fmt.Println("git", "push", targetRepository, currentBranch)
	// cmd := exec.Command("git", "push", targetRepository, currentBranch)
	var output bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("git", "push", targetRepository, currentBranch)
	cmd.Stdout = &output
	cmd.Stderr = &errOut
	cmd.Run()

	fmt.Println(output.String())
	fmt.Println(errOut.String())
}
