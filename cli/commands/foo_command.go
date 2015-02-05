package commands

import (
	"log"
)

type FooCommand struct {
	HelpText string
	SynopsisText string
	RunArgs []string
}

func (c *FooCommand) Help() string {
	return c.HelpText
}

func (c *FooCommand) Run(args []string) int {
	c.RunArgs = args
	
	log.Println("Doing some foo...")
	
	for i := range args {
		log.Printf("Arg: %s\n", args[i])
	}
	
	log.Println("Ok, done with the foo")
	
	return 0
}

func (c *FooCommand) Synopsis() string {
	return c.SynopsisText
}