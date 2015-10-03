package main

import (
	commands "github.com/d-smith/go-examples/cli/commands"
	"github.com/mitchellh/cli"
	"os"
)

//Can't reuse CLI in a repl because only the first command is run via
// sync.Once
func makeCLI() *cli.CLI {
	c := cli.NewCLI("cli", "1.0.0")

	fooCommand := &commands.FooCommand{
		HelpText:     "foo blah blah blah flibby dibby do",
		SynopsisText: "do some foo",
	}

	barCommand := &commands.BarCommand{
		HelpText:     "bar bar barness",
		SynopsisText: "when you want bar",
	}

	c.Commands = map[string]cli.CommandFactory{
		"foo": func() (cli.Command, error) {
			return fooCommand, nil
		},

		"bar": func() (cli.Command, error) {
			return barCommand, nil
		},
	}

	return c
}

func main() {
	ui := &cli.BasicUi{Reader: os.Stdin, Writer: os.Stdout}
	c := makeCLI()
	c.Args = os.Args[1:]

	_, err := c.Run()
	if err != nil {
		ui.Error("Error executing command" + err.Error())
	}

}
