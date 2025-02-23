package cmd

type Argument struct {
	name  string
	value string
}

func NewArgument(name, value string) *Argument {
	return &Argument{
		name:  name,
		value: value,
	}
}

func (a *Argument) GetName() string {
	return a.name
}

func (a *Argument) GetValue() string {
	return a.value
}
