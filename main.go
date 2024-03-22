package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	// "mvdan.cc/sh/v3/syntax"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const TabSize = 2

type playground struct {
	app.Compo

	ShellInput      string
	FormattedOutput string
}

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

func (p *playground) onFormat(ctx app.Context, e app.Event) {
	fmt.Printf("Formatting input:\n%s\n", p.ShellInput)
	formatted, err := formatShellScript(p.ShellInput)
	if err != nil {
		p.FormattedOutput = "Whoopsi something went wrong"
		return
	}
	p.FormattedOutput = formatted
}

// This is meant to insert "tabs" instead of tabbing out of the textarea
func (p *playground) onKeyDown(ctx app.Context, e app.Event) {
	if e.Get("key").String() == "Tab" {
		e.PreventDefault()
		// Preserve current caret position
		// Not sure why Start and End seperatly as i expect them to be the same :shrug:
		start := ctx.JSSrc().Get("selectionStart").Int()
		end := ctx.JSSrc().Get("selectionEnd").Int()

		// If we received a tab we need to insert the specified amount of spaces at the right position.
		value := ctx.JSSrc().Get("value").String()
		value = value[:start] + strings.Repeat(" ", TabSize) + value[end:]
		ctx.JSSrc().Set("value", value)

		// update caret to reflect inserted spaces for the tab
		ctx.JSSrc().Set("selectionStart", start+TabSize)
		ctx.JSSrc().Set("selectionEnd", end+TabSize)
	}
}

func (p *playground) Render() app.UI {
	return app.Div().Body(
		app.Textarea().ID("shellinput").Name("shellinput").Text(p.ShellInput).Cols(80).Rows(20).OnInput(p.ValueTo(&p.ShellInput)).OnKeyDown(p.onKeyDown),
		app.Input().ID("run-format").Value("run formatter").Type("submit").OnClick(p.onFormat),
		app.Textarea().ID("formattedoutput").Name("formattedoutput").Text(p.FormattedOutput).Cols(80).Rows(20),
	)
}

func main() {
	fmt.Println("Shell playground starting up...")
	app.Route("/", &playground{})

	// If this is in the browser start the eventloop in wasm
	app.RunWhenOnBrowser()

	// Otherwise serve
	http.Handle("/", &app.Handler{
		Name:        "Playground",
		Description: "The shell playground!",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started!")
}
