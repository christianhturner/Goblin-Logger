package main

import (
	"fmt"
	"os/exec"

	"github.com/christianhturner/goblin-logger/internal/loggerCmd"
)

var cmd = exec.Command("ls", "./")

func main() {
	fmt.Print("Hello World")
	loggerCmd.New(*cmd, loggerCmd.LoggerOpts{PollFreq: 32, EnablePolling: true})
}
