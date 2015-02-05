package commands

import (
	"log"
)

type BarCommand struct {
	HelpText string
	SynopsisText string
	RunArgs []string
}

func (c *BarCommand) Help() string {
	return c.HelpText
}

func (c *BarCommand) Run(args []string) int {
	c.RunArgs = args
	
	log.Println("Doing some bar...")
	
	for i := range args {
		log.Printf("Arg: %s\n", args[i])
	}
	
	log.Println("Ok, done with the bar")
	
	return 0
}

func (c *BarCommand) Synopsis() string {
	return c.SynopsisText
}