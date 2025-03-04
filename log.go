package logger

import "fmt"

var Version string = "0.1"

func Log(mess string) {
	fmt.Println("[LOG] " + mess)
}
