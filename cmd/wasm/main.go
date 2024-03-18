package main

import (
	"fmt"
	"strings"
	"syscall/js"
	// "mvdan.cc/sh/v3/syntax"
)

func formatShellScript(input string) (string, error) {
	// yes this is dumb... but just to show we are removing any indentation
	// and trailing whitespace for now as "formatting"
	// ToDo actually use sh/v3/syntax to format the script.
	trimmed := ""
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		trimmed += strings.TrimSpace(line)
		trimmed += "\n"
	}
	return trimmed, nil
}

func formatWrapper() js.Func {
	formatterFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Internal Error: Invalid number of arguments passed - only pass in one string to format"
		}
		inputScript := args[0].String()
		fmt.Printf("input scripts received for formatting:\n%s\n", inputScript)
		formatted, err := formatShellScript(inputScript)
		if err != nil {
			fmt.Printf("unable to format script: %s\n", err)
			return err.Error()
		}
		return formatted
	})
	return formatterFunc
}

func main() {
	fmt.Println("shell playground loaded")
	js.Global().Set("formatShell", formatWrapper())
	<-make(chan struct{})
}
