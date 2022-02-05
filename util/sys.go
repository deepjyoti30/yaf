package util

import (
	"os/exec"
	"strings"
)

// runSysctlCmd runs the passed sysctl command
// and returns the ourput or error
func runSysctlCmd(args string) (string, error) {
	sysctlCmdStr := "sysctl"
	cmd := exec.Command(sysctlCmdStr, strings.Split(args, " ")...)
	stdout, err := cmd.Output()
	return string(stdout), err
}
