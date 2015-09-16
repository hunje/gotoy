package tool

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecStringCommand(commands string) string {
	cmds := strings.Split(commands, " ")
	if len(cmds) < 1 {
		return ""
	}
	return ExecCommand(cmds)
}

func ExecCommand(commands []string) string {
	cmd := exec.Command(commands[0], commands[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out.String())
}
