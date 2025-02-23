package cmd

import (
	"fmt"
	"os"
)

func (c *Cmd) print(message string) {
	if len(message) > 0 {
		os.Stdout.Write([]byte(message))
	}
}

func (c *Cmd) printGreeting() {
	if c.greeting != "" {
		c.print(fmt.Sprintf("%s\n\n", c.greeting))
	}
}

func (c *Cmd) printUsage() {
	if c.usage != "" {
		c.print(fmt.Sprintf("Usage:\n\n%s\n\n", c.usage))
	}
}

func (c *Cmd) printFlags() {
	if len(c.flags) > 0 {
		c.print("The flags are:\n\n")

		for _, flag := range c.flags {
			c.print(fmt.Sprintf(
				"\t-%s\t%s\t%s\n",
				flag.name,
				flag.alias,
				flag.desc,
			))
		}

		c.print("\n")
	}
}

func (c *Cmd) printCommands() {
	if len(c.commands) > 0 {
		c.print("The commands are:\n\n")

		for _, command := range c.commands {
			c.print(fmt.Sprintf(
				"\t%s\t%s\t%s\n",
				command.name,
				command.alias,
				command.desc,
			))
		}

		c.print("\n")
	}
}

func (c *Cmd) printHelp() {
	c.printGreeting()
	c.printUsage()
	c.printFlags()
	c.printCommands()
}

func (c *Cmd) printNoArgs() {
	c.print("no arguments passed\n")
}
