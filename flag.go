package cmd

type Flag struct {
	name         string
	alias        string
	desc         string
	required     bool
	defaultValue string
}

func NewFlag(name, alias, desc string, required bool, defaultValue string) *Flag {
	return &Flag{
		name:         name,
		alias:        alias,
		desc:         desc,
		required:     required,
		defaultValue: defaultValue,
	}
}

func (f *Flag) SetName(name string) *Flag {
	f.name = name
	return f
}

func (f *Flag) SetAlias(alias string) *Flag {
	f.alias = alias
	return f
}

func (f *Flag) SetDesc(desc string) *Flag {
	f.desc = desc
	return f
}

func (f *Flag) SetRequired(required bool) *Flag {
	f.required = required
	return f
}

func (f *Flag) SetDefaultValue(defaultValue string) *Flag {
	f.defaultValue = defaultValue
	return f
}
