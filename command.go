package cmd

type Command struct {
	name      string         // Name
	alias     string         // Alias
	desc      string         // Description
	f         func(cmd *Cmd) // Function
	arguments []*Argument    // Arguments
}

// Create new Command instance return pointer to Command
func NewCommand(name, alias, desc string, f func(cmd *Cmd)) *Command {
	return &Command{
		name:      name,
		alias:     alias,
		desc:      desc,
		f:         f,
		arguments: make([]*Argument, 0),
	}
}

// Set command name return pointer to Command
func (c *Command) SetName(name string) *Command {
	c.name = name
	return c
}

// Set command alias return pointer to Command
func (c *Command) SetAlias(alias string) *Command {
	c.alias = alias
	return c
}

// Set command description return pointer to Command
func (c *Command) SetDesc(desc string) *Command {
	c.desc = desc
	return c
}

// Set command execute function return pointer to Command
func (c *Command) SetFunc(f func(cmd *Cmd)) *Command {
	c.f = f
	return c
}

// Set command arguments return pointer to Command
func (c *Command) SetArguments(args []*Argument) *Command {
	c.arguments = args
	return c
}

// Set command arguments values
func (c *Command) SetArgumentsValues(values []string) {
	if len(c.arguments) > 0 && len(values) > 0 {
		for i, arg := range c.arguments {
			if len(values) > i {
				arg.value = values[i]
			}
		}
	}
}

// Add argument to command return pointer to Command
func (c *Command) AddArgument(argument string) *Command {
	c.arguments = append(c.arguments, NewArgument(argument, ""))
	return c
}

// Get command argument by name return pointer to Argument
func (c *Command) GetArgument(name string) *Argument {
	if len(c.arguments) > 0 {
		for _, arg := range c.arguments {
			if arg.GetName() == name {
				return arg
			}
		}
	}

	return nil
}
