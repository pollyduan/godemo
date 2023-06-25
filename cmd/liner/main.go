package main

import (
	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	line.SetMultiLineMode(true)
	defer line.Close()
	for {
		if cmd, err := line.Prompt("-> "); err == nil {
			if cmd == "exit" {
				break
			}
			line.AppendHistory(cmd)
			println(cmd)
		} else {
			if err != liner.ErrPromptAborted {
				println(err.Error())
			}
			break
		}
	}

}
