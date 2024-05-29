package main

import (
	"fmt"
	"os/exec"

	"github.com/christianhturner/goblin-logger/internal/loggerCmd"
)

var cmd = exec.Command("ls", "./")

func main() {
	lc, err := loggerCmd.New(*cmd, &loggerCmd.LoggerOpts{PollFreq: 0, Schedule: 0})
	if err != nil {
		fmt.Errorf("An error occured: %v", err)
	}
	out, errout, err := lc.Run()
	if err != nil {
		fmt.Errorf("An error occured %v", err)
	}
	fmt.Printf("stdout:\n%s", out)
	fmt.Printf("stderr:\n%s", errout)
}
