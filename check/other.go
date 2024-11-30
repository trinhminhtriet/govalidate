package check

import (
	"bytes"
	"os/exec"
)

func runCmd(cmd string, arg ...string) (string, error) {
	c := exec.Command(cmd, arg...)
	out, err := c.CombinedOutput()
	return string(bytes.TrimSpace(out)), err
}
