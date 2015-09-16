package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/hunje/gotoy/tool"
	"os/exec"
)

var targetRepository string
var forceUpdate bool

func init() {
	flag.StringVar(&targetRepository, "t", "", "target repository")
	flag.BoolVar(&forceUpdate, "f", false, "set to force update")
	flag.Parse()
}

func GetBranchName() string {
	commands := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}
	return tool.ExecCommand(commands)
}

func main() {
	if targetRepository == "" {
		flag.PrintDefaults()
		return
	}

	commands := []string{"push"}

	currentBranch := GetBranchName()

	if len(currentBranch) <= 0 {
		fmt.Println("here is not under git")
		return
	}

	if forceUpdate == true {
		commands = append(commands, "-f")
	}

	commands = append(commands, targetRepository)
	commands = append(commands, currentBranch)

	fmt.Println(commands)
	var output bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("git", commands...)
	cmd.Stdout = &output
	cmd.Stderr = &errOut
	cmd.Run()

	fmt.Println(output.String())
	fmt.Println(errOut.String())
}
