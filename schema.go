package schema

import (
	"encoding/json"
	"fmt"
)

type Schema struct {
	Title      string                 `json:"title"`
	Type       string                 `json:"type"`
	Properties map[string]Property    `json:"properties"`
	AllOf      []ConditionalStructure `json:"allOf,omitempty"`
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
func NewSchemaBuilder(title string) *Builder {
	return &Builder{
		schema: Schema{
			Title:      title,
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
	b.schema.AllOf = b.conditions
	return b.schema
}

func buildAndDump(generated *Generated) {

	jsonBytes, err := json.MarshalIndent(generated, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
