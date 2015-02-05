package main

import (
    "log"
    "os"
    "github.com/mitchellh/cli"
    commands "github.com/d-smith/go-examples/cli/commands"
)





func main() {
    c := cli.NewCLI("mycli", "1.0.0")
    c.Args = os.Args[1:]
    
    fooCommand := &commands.FooCommand{
				HelpText: "foo blah blah blah flibby dibby do",
				SynopsisText: "do some foo",
			}
    
    barCommand := &commands.BarCommand {
    	HelpText: "bar bar barness",
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

    exitStatus, err := c.Run()
    if err != nil {
        log.Println(err)
    }

    os.Exit(exitStatus)
}
