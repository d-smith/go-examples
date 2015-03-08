package commands

import (
	"log"
)

//BarCommand is an example command
type BarCommand struct {
	HelpText string
	SynopsisText string
	RunArgs []string
}

//Help returns the help text for the Bar command
func (c *BarCommand) Help() string {
	return c.HelpText
}

//Run executes the command logic
func (c *BarCommand) Run(args []string) int {
	c.RunArgs = args
	
	log.Println("Doing some bar...")
	
	for i := range args {
		log.Printf("Arg: %s\n", args[i])
	}
	
	log.Println("Ok, done with the bar")
	
	return 0
}

//Synopsis returns a concise summary of the command
func (c *BarCommand) Synopsis() string {
	return c.SynopsisText
}