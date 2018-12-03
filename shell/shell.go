package shell

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

// RunScript executes shell script with supplied arguments.
func RunScript(script string, args ...string) (string, error) {
	cmd := exec.Command(script, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if nil != err {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}
