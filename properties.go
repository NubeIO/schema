package schema

// AddProperty adds a property to the schema.
func (b *Builder) AddProperty(name string, property Property) *Builder {
	b.schema.Properties[name] = property
	b.uiOrder = append(b.uiOrder, name)
	return b
}

func (b *Builder) NewNumber(name, title string, required bool, min, max *float64, defaultValue float64) *Builder {
	n := NewNumber(title, required, min, max, defaultValue)
	b.AddProperty(name, n)
	return b
}

func (b *Builder) NewString(name, title string, required bool, minLength, maxLength int, defaultValue string) *Builder {
	n := NewString(title, required, minLength, maxLength, defaultValue)
	b.AddProperty(name, n)
	return b
}

func (b *Builder) NewBool(name, title string, required bool, defaultValue bool) *Builder {
	n := NewBool(title, required, defaultValue)
	b.AddProperty(name, n)
	return b
}

// NewBool creates a new bool property.
func NewBool(title string, required bool, defaultValue bool) Property {
	prop := Property{
		Type:    "boolean",
		Title:   title,
		Default: defaultValue,
	}
	if required {
		prop.Required = []string{title}
	}
	return prop
}

// NewString creates a new string property.
func NewString(title string, required bool, minLength, maxLength int, defaultValue string) Property {
	prop := Property{
		Type:    "string",
		Title:   title,
		Default: defaultValue,
	}
	if minLength > 0 {
		prop.MinLength = &minLength
	}
	if maxLength > 0 {
		prop.MaxLength = &maxLength
	}
	if required {
		prop.Required = []string{title}
	}
	return prop
}

// NewNumber creates a new number property.
func NewNumber(title string, required bool, min, max *float64, defaultValue float64) Property {
	prop := Property{
		Type:    "number",
		Title:   title,
		Default: defaultValue,
		Minimum: min,
		Maximum: max,
	}
	if required {
		prop.Required = []string{title}
	}
	return prop
}

// NewEnumString creates a new string enum property.
func NewEnumString(title string, options, enumNames []string) Property {
	return Property{
		Type:      "string",
		Title:     title,
		Enum:      options,
		EnumNames: enumNames,
	}
}

// NewEnumNumber creates a new number enum property.
func NewEnumNumber(title string, options []float64, enumNames []string) Property {
	return Property{
		Type:      "number",
		Title:     title,
		Enum:      options,
		EnumNames: enumNames,
	}
}

func NewEmail(title string) Property {
	return Property{
		Type:   "string",
		Format: "email",
		Title:  title,
	}
}

// NewURI creates a new property for URI.
func NewURI(title string) Property {
	return Property{
		Type:   "string",
		Format: "uri",
		Title:  title,
	}
}

// NewPassword creates a new password property with specified constraints.
func NewPassword(title string, minLength int) Property {
	// Regex pattern for at least one uppercase letter, one number, and one special character
	passwordPattern := `^(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z\d]).+$`

	return Property{
		Type:        "string",
		Title:       title,
		MinLength:   &minLength,
		Pattern:     passwordPattern,
		Description: "Must contain at least one uppercase letter, one number, and one special character.",
	}
}

func (b *Builder) SetProperty(name string, property Property) *Builder {
	b.schema.Properties[name] = property
	return b
}
