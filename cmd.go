package cmd

import (
	"log"
	"os"
	"strings"
)

type Cmd struct {
	greeting    string            // Greeting
	usage       string            // Usage
	flags       []*Flag           // Flags
	flagsValues map[string]string // Flags values
	command     *Command          // Active command
	commands    []*Command        // Supported commands
	arguments   []string          // Arguments
}

// Create new Cmd instance and return pointer to Cmd
func NewCmd() *Cmd {
	return &Cmd{
		flags:       make([]*Flag, 0),
		flagsValues: make(map[string]string),
		commands:    make([]*Command, 0),
		arguments:   make([]string, 0),
	}
}

// Set greeting and return pointer to Cmd
func (c *Cmd) SetGreeting(greeting string) *Cmd {
	c.greeting = greeting
	return c
}

// Set usage and return pointer to Cmd
func (c *Cmd) SetUsage(usage string) *Cmd {
	c.usage = usage
	return c
}

// Register flags and return pointer to Cmd
func (c *Cmd) RegisterFlag(name string, alias string, desc string, required bool, defaultValue string) *Cmd {
	c.flags = append(c.flags, NewFlag(name, alias, desc, required, defaultValue))
	return c
}

// Add flag and return pointer to Cmd
func (c *Cmd) AddFlag(flag *Flag) *Cmd {
	c.flags = append(c.flags, flag)
	return c
}

// Create Flag helper function return pointer to Flag
func (c *Cmd) NewFlag() *Flag {
	return &Flag{}
}

// Register command and return pointer to Cmd
func (c *Cmd) RegisterCommand(name string, alias string, desc string, required bool, f func(cmd *Cmd)) *Cmd {
	c.commands = append(c.commands, NewCommand(name, alias, desc, f))
	return c
}

// Add command and return pointer Cmd
func (c *Cmd) AddCommand(command *Command) *Cmd {
	c.commands = append(c.commands, command)
	return c
}

// Create Command helper function return pointer to Command
func (c *Cmd) NewCommand() *Command {
	return &Command{}
}

// Parse arguments
func (c *Cmd) Parse(args []string) {
	if len(args) > 0 {
		c.parse(args)
	} else {
		c.parse(os.Args[1:])
	}
}

// Get Flag value by name return string
func (c *Cmd) GetFlagValue(name string) string {
	if v, ok := c.flagsValues[name]; ok {
		return v
	}

	if c.hasFlag(name) {
		flag := c.getFlag(name)
		if len(flag.defaultValue) > 0 {
			return flag.defaultValue
		}
	}

	return ""
}

// Get all flags values return map[string]string
func (c *Cmd) GetFlagsValues() map[string]string {
	return c.flagsValues
}

// Get active command return Command
func (c *Cmd) GetCommand() *Command {
	return c.command
}

// Parse arguments
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

// Check that flag exists by name return bool
func (c *Cmd) hasFlag(name string) bool {
	for _, flag := range c.flags {
		if flag.name == name || flag.alias == name {
			return true
		}
	}

	return false
}

// Get flag by name return pointer to Flag
func (c *Cmd) getFlag(name string) *Flag {
	for _, flag := range c.flags {
		if flag.name == name || flag.alias == name {
			return flag
		}
	}

	return nil
}

// Check that command exists by name return bool
func (c *Cmd) hasCommand(name string) bool {
	for _, command := range c.commands {
		if command.name == name || command.alias == name {
			return true
		}
	}

	return false
}

// Get command by name return pointer to Command
func (c *Cmd) getCommand(name string) *Command {
	for _, command := range c.commands {
		if command.name == name || command.alias == name {
			return command
		}
	}

	return nil
}
