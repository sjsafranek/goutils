package shell

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunScript executes shell script with supplied arguments.
func RunScript(script string, args ...string) {
	cmd := exec.Command(script, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if nil != err {
		fmt.Println(stderr.String())
		return
	}
	fmt.Println(stdout.String())
}

