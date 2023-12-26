package schema

type UI struct {
	UiOrder      []string              `json:"ui:order"`
	UiProperties map[string]UIProperty `json:"uiProperties"`
}

type UIProperty struct {
	Widget  string                 `json:"ui:widget,omitempty"`
	Options map[string]interface{} `json:"ui:options,omitempty"`
}

// AddUIOrder sets the UI order of properties.
func (b *UI) AddUIOrder(order []string) {
	b.UiOrder = order
	return
}

func (b *UI) AddUIProperty(propertyName string, uiProp UIProperty) {
	if b.UiProperties == nil {
		b.UiProperties = make(map[string]UIProperty)
	}
	b.UiProperties[propertyName] = uiProp

}
