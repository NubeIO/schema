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

	If   *Condition `json:"if,omitempty"`
	Then *Condition `json:"then,omitempty"`
	Else *Condition `json:"else,omitempty"`
}

type UIProperty struct {
	Widget  string                 `json:"ui:widget,omitempty"`
	Options map[string]interface{} `json:"ui:options,omitempty"`
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

type ConditionalStructure struct {
	If   Condition `json:"if,omitempty"`
	Then Condition `json:"then,omitempty"`
}

type Condition struct {
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
}

type SchemaBuilder struct {
	schema       Schema
	uiOrder      []string
	uiProperties map[string]UIProperty
	conditions   []ConditionalStructure
}

func (b *SchemaBuilder) AddUIProperty(propertyName string, uiProp UIProperty) *SchemaBuilder {
	if b.uiProperties == nil {
		b.uiProperties = make(map[string]UIProperty)
	}
	b.uiProperties[propertyName] = uiProp
	return b
}

// NewSchemaBuilder creates a new SchemaBuilder.
func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{
		schema: Schema{
			Type:       "object",
			Properties: make(map[string]Property),
		},
		uiOrder:    []string{},
		conditions: make([]ConditionalStructure, 0),
	}
}

func (b *SchemaBuilder) AddConditionalStructure(ifProps, thenProps map[string]Property, thenRequired []string) *SchemaBuilder {
	condition := ConditionalStructure{
		If: Condition{
			Properties: ifProps,
		},
		Then: Condition{
			Properties: thenProps,
			Required:   thenRequired,
		},
	}
	b.schema.AllOf = append(b.schema.AllOf, condition)
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

func (b *SchemaBuilder) AddNestedProperties(parentName string, nestedProperties map[string]Property) *SchemaBuilder {
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

// AddArrayProperty adds an array property with nested items to the schema.
func (b *SchemaBuilder) AddArrayProperty(name string, itemProperties map[string]Property) *SchemaBuilder {
	b.schema.Properties[name] = Property{
		Type:  "array",
		Items: &Property{Type: "object", Properties: itemProperties},
	}
	b.uiOrder = append(b.uiOrder, name)
	return b
}

func NewNumberWithWidget(title string, required bool, min, max, multipleOf *float64, defaultValue float64, widget string, options map[string]interface{}) Property {
	prop := NewNumber(title, required, min, max, defaultValue)
	if multipleOf != nil {
		prop.MultipleOf = multipleOf
	}
	return prop
}

// AddProperty adds a property to the schema.
func (b *SchemaBuilder) AddProperty(name string, property Property) *SchemaBuilder {
	b.schema.Properties[name] = property
	b.uiOrder = append(b.uiOrder, name)
	return b
}

// AddUIOrder sets the UI order of properties.
func (b *SchemaBuilder) AddUIOrder(order []string) *SchemaBuilder {
	b.uiOrder = order
	return b
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

func (b *SchemaBuilder) SetProperty(name string, property Property) *SchemaBuilder {
	b.schema.Properties[name] = property
	return b
}

func (b *SchemaBuilder) AddCondition(condition ConditionalStructure) *SchemaBuilder {
	b.conditions = append(b.conditions, condition)
	return b
}

func (b *SchemaBuilder) Build() Schema {
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

func buildAndDump(builder *SchemaBuilder) {
	schema := builder.Build()
	jsonBytes, err := json.MarshalIndent(schema, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
