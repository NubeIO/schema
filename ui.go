package schema

type UIProperty struct {
	Widget  string                 `json:"ui:widget,omitempty"`
	Options map[string]interface{} `json:"ui:options,omitempty"`
}

// AddUIOrder sets the UI order of properties.
func (b *Builder) AddUIOrder(order []string) *Builder {
	b.uiOrder = order
	return b
}

func (b *Builder) AddUIProperty(propertyName string, uiProp UIProperty) *Builder {
	if b.uiProperties == nil {
		b.uiProperties = make(map[string]UIProperty)
	}
	b.uiProperties[propertyName] = uiProp
	return b
}
