package cmd

type Argument struct {
	name  string // Name
	value string // Value
}

// Create new Argument instance and return pointer to Argument
func NewArgument(name, value string) *Argument {
	return &Argument{
		name:  name,
		value: value,
	}
}

// Set argument name return pointer to Argument
func (a *Argument) SetName(name string) *Argument {
	a.name = name
	return a
}

// Set argument value return pointer to Argument
func (a *Argument) SetValue(value string) *Argument {
	a.value = value
	return a
}

// Get argument name return string
func (a *Argument) GetName() string {
	return a.name
}

// Get argument value return string
func (a *Argument) GetValue() string {
	return a.value
}
