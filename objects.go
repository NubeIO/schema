package schema

// AddArrayProperty adds an array property with nested items to the schema.
func (b *Builder) AddArrayProperty(name string, itemProperties map[string]Property) *Builder {
	b.schema.Properties[name] = Property{
		Type:  "array",
		Items: &Property{Type: "object", Properties: itemProperties},
	}
	b.uiOrder = append(b.uiOrder, name)
	return b
}

type NestedProperty struct {
	Key      string
	Property Property
}

// CreateNestedProperties creates a map of properties for nested usage.
func CreateNestedProperties(properties ...NestedProperty) map[string]Property {
	nestedProps := make(map[string]Property)
	for _, np := range properties {
		nestedProps[np.Key] = np.Property
	}
	return nestedProps
}

func (b *Builder) AddNestedProperties(parentName string, nestedProperties map[string]Property) *Builder {
	if _, ok := b.schema.Properties[parentName]; !ok {
		b.schema.Properties[parentName] = Property{
			Type:       "object",
			Properties: make(map[string]Property),
		}
	}

	for name, prop := range nestedProperties {
		b.schema.Properties[parentName].Properties[name] = prop
	}

	return b
}
