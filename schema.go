package schema

import (
	"encoding/json"
	"fmt"
)

type Schema struct {
	Type       string                 `json:"type"`
	Properties map[string]Property    `json:"properties"`
	AllOf      []ConditionalStructure `json:"allOf,omitempty"`
	UI         map[string]interface{} `json:"ui,omitempty"`
}

type Property struct {
	Type        string              `json:"type,omitempty"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	Enum        interface{}         `json:"enum,omitempty"`
	Items       *Property           `json:"items,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Default     interface{}         `json:"default,omitempty"`
	Minimum     *float64            `json:"minimum,omitempty"`
	Maximum     *float64            `json:"maximum,omitempty"`
	MinLength   *int                `json:"minLength,omitempty"`
	MaxLength   *int                `json:"maxLength,omitempty"`
	Required    []string            `json:"required,omitempty"`
	EnumNames   []string            `json:"enumNames,omitempty"`
	MultipleOf  *float64            `json:"multipleOf,omitempty"`
	Pattern     string              `json:"pattern,omitempty"`
	Format      string              `json:"format,omitempty"`
	Const       string              `json:"const,omitempty"`
}

type Builder struct {
	schema       Schema
	uiOrder      []string
	uiProperties map[string]UIProperty
	conditions   []ConditionalStructure
}

// NewSchemaBuilder creates a new SchemaBuilder.
func NewSchemaBuilder() *Builder {
	return &Builder{
		schema: Schema{
			Type:       "object",
			Properties: make(map[string]Property),
		},
		uiOrder:    []string{},
		conditions: make([]ConditionalStructure, 0),
	}
}

func NewNumberWithWidget(title string, required bool, min, max, multipleOf *float64, defaultValue float64, widget string, options map[string]interface{}) Property {
	prop := NewNumber(title, required, min, max, defaultValue)
	if multipleOf != nil {
		prop.MultipleOf = multipleOf
	}
	return prop
}

func (b *Builder) Build() Schema {
	// Include UI Order if specified
	if len(b.uiOrder) > 0 {
		if b.schema.UI == nil {
			b.schema.UI = make(map[string]interface{})
		}
		b.schema.UI["ui:order"] = b.uiOrder
	}

	// Add UI properties to the schema
	for propName, uiProp := range b.uiProperties {
		if b.schema.UI == nil {
			b.schema.UI = make(map[string]interface{})
		}
		b.schema.UI[propName] = uiProp
	}
	b.schema.AllOf = b.conditions
	return b.schema
}

func buildAndDump(builder *Builder) {
	schema := builder.Build()
	jsonBytes, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
