package cmd

import (
	"log"
	"os"
	"strings"
)

type Cmd struct {
	greeting    string
	usage       string
	flags       []*Flag
	flagsValues map[string]string
	command     *Command
	commands    []*Command
	arguments   []string
}

func NewCmd() *Cmd {
	return &Cmd{
		flags:       make([]*Flag, 0),
		flagsValues: make(map[string]string),
		commands:    make([]*Command, 0),
		arguments:   make([]string, 0),
	}
}

func (c *Cmd) SetGreeting(greeting string) *Cmd {
	c.greeting = greeting
	return c
}

func (c *Cmd) SetUsage(usage string) *Cmd {
	c.usage = usage
	return c
}

func (c *Cmd) RegisterFlag(name string, alias string, desc string, required bool, defaultValue string) *Cmd {
	c.flags = append(c.flags, NewFlag(name, alias, desc, required, defaultValue))
	return c
}

func (c *Cmd) AddFlag(flag *Flag) *Cmd {
	c.flags = append(c.flags, flag)
	return c
}

func (c *Cmd) NewFlag() *Flag {
	return &Flag{}
}

func (c *Cmd) RegisterCommand(name string, alias string, desc string, required bool, f func(cmd *Cmd)) *Cmd {
	c.commands = append(c.commands, NewCommand(name, alias, desc, f))
	return c
}

func (c *Cmd) AddCommand(command *Command) *Cmd {
	c.commands = append(c.commands, command)
	return c
}

func (c *Cmd) NewCommand() *Command {
	return &Command{}
}

func (c *Cmd) Parse() {
	c.parse(os.Args[1:])
}

func (c *Cmd) GetFlagValue(name string) string {
	if v, ok := c.flagsValues[name]; ok {
		return v
	}

	return ""
}

func (c *Cmd) GetFlagsValues() map[string]string {
	return c.flagsValues
}

func (c *Cmd) GetCommand() *Command {
	return c.command
}

func (c *Cmd) parse(args []string) {
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			clean := strings.TrimLeft(args[i], "-")

			if c.hasFlag(clean) {
				value := args[i+1]
				flag := c.getFlag(clean)
				c.flagsValues[flag.name] = value
				i++
			} else if c.hasCommand(clean) {
				command := c.getCommand(clean)
				c.command = command
			} else {
				if c.command == nil {
					log.Panicf("unsupported argument \"%s\"\n", clean)
				} else {
					c.arguments = append(c.arguments, args[i])
				}
			}
		}

		if c.command != nil {
			c.command.SetArgumentsValues(c.arguments)
			c.command.f(c)
		} else {
			c.printHelp()
		}
	} else {
		if len(c.flags) > 0 || len(c.commands) > 0 {
			c.printHelp()
		} else {
			c.printNoArgs()
		}
	}
}

func (c *Cmd) hasFlag(name string) bool {
	for _, flag := range c.flags {
		if flag.name == name || flag.alias == name {
			return true
		}
	}

	return false
}

func (c *Cmd) getFlag(name string) *Flag {
	for _, flag := range c.flags {
		if flag.name == name || flag.alias == name {
			return flag
		}
	}

	return nil
}

func (c *Cmd) hasCommand(name string) bool {
	for _, command := range c.commands {
		if command.name == name || command.alias == name {
			return true
		}
	}

	return false
}

func (c *Cmd) getCommand(name string) *Command {
	for _, command := range c.commands {
		if command.name == name || command.alias == name {
			return command
		}
	}

	return nil
}
