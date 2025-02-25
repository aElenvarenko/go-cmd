package cmd

type Flag struct {
	name         string // Name
	alias        string // Alias
	desc         string // Description
	required     bool   // Required
	defaultValue string // Default value
}

// Create new Flag instance return pointer to Flag
func NewFlag(name, alias, desc string, required bool, defaultValue string) *Flag {
	return &Flag{
		name:         name,
		alias:        alias,
		desc:         desc,
		required:     required,
		defaultValue: defaultValue,
	}
}

// Set flag name return pointer to Flag
func (f *Flag) SetName(name string) *Flag {
	f.name = name
	return f
}

// Set flag alias return pointer to Flag
func (f *Flag) SetAlias(alias string) *Flag {
	f.alias = alias
	return f
}

// Set flag description return pointer to Flag
func (f *Flag) SetDesc(desc string) *Flag {
	f.desc = desc
	return f
}

// Set flag required return pointer to Flag
func (f *Flag) SetRequired(required bool) *Flag {
	f.required = required
	return f
}

// Set flag defaultValue return pointer to Flag
func (f *Flag) SetDefaultValue(defaultValue string) *Flag {
	f.defaultValue = defaultValue
	return f
}

// Get flag name return string
func (f *Flag) GetName() string {
	return f.name
}

// Get flag alias return string
func (f *Flag) GetAlias() string {
	return f.alias
}

// Get flag description return string
func (f *Flag) GetDesc() string {
	return f.desc
}

// Get flag required return bool
func (f *Flag) GetRequired() bool {
	return f.required
}

// Get flag default value return string
func (f *Flag) GetDefaultValue() string {
	return f.defaultValue
}
