package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var targetRepository string

func init() {
	flag.StringVar(&targetRepository, "t", "", "target repository")
	flag.Parse()
}

func main() {
	if targetRepository == "" {
		flag.PrintDefaults()
		return
	}

	cmd := exec.Command("git", "push", targetRepository, "`git get-branch-name`")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
	}
}
