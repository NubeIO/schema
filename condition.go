package schema

type ConditionalStructure struct {
	If   Condition `json:"if,omitempty"`
	Then Condition `json:"then,omitempty"`
}

type Condition struct {
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
}

func (b *Builder) AddConditionalStructure(ifProps, thenProps map[string]Property, thenRequired []string) *Builder {
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

func (b *Builder) AddCondition(condition ConditionalStructure) *Builder {
	b.conditions = append(b.conditions, condition)
	return b
}
