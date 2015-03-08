package commands

import (
	"log"
)

//FooCommand defines the foo command and related runtime documentation
type FooCommand struct {
	HelpText string
	SynopsisText string
	RunArgs []string
}

//Help returns a string describing the inputs associated with the command
func (c *FooCommand) Help() string {
	return c.HelpText
}

//Run executes the command logic
func (c *FooCommand) Run(args []string) int {
	c.RunArgs = args
	
	log.Println("Doing some foo...")
	
	for i := range args {
		log.Printf("Arg: %s\n", args[i])
	}
	
	log.Println("Ok, done with the foo")
	
	return 0
}

//Synopsis returns a concise summary of the foo command
func (c *FooCommand) Synopsis() string {
	return c.SynopsisText
}