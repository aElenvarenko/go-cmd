package cmd

type Command struct {
	name      string
	alias     string
	desc      string
	f         func(cmd *Cmd)
	arguments []*Argument
}

func NewCommand(name, alias, desc string, f func(cmd *Cmd)) *Command {
	return &Command{
		name:      name,
		alias:     alias,
		desc:      desc,
		f:         f,
		arguments: make([]*Argument, 0),
	}
}

func (c *Command) SetName(name string) *Command {
	c.name = name
	return c
}

func (c *Command) SetAlias(alias string) *Command {
	c.alias = alias
	return c
}

func (c *Command) SetDesc(desc string) *Command {
	c.desc = desc
	return c
}

func (c *Command) SetFunc(f func(cmd *Cmd)) *Command {
	c.f = f
	return c
}

func (c *Command) SetArguments(args []*Argument) *Command {
	c.arguments = args
	return c
}

func (c *Command) SetArgumentsValues(values []string) {
	if len(c.arguments) > 0 && len(values) > 0 {
		for i, arg := range c.arguments {
			if len(values) > i {
				arg.value = values[i]
			}
		}
	}
}

func (c *Command) AddArgument(argument string) *Command {
	c.arguments = append(c.arguments, NewArgument(argument, ""))
	return c
}

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
