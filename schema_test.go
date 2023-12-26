package schema

import (
	"testing"
)

// https://rjsf-team.github.io/react-jsonschema-form/
// put json in here to test

// test case for IF THEN
func Test_basic(t *testing.T) {
	builder := NewSchemaBuilder()

	builder.AddProperty("exampleString", NewString("Example String", true, 3, 10, "default value"))

	// Adding a number property
	builder.AddProperty("exampleNumber", NewNumber("Example Number", false, nil, nil, 42))

	// Adding an enum string property
	builder.AddProperty("exampleEnumString", NewEnumString("Example Enum String", []string{"a", "b"}, []string{"Option A", "Option B"}))

	// Adding a ratio property
	minRatio, maxRatio := 0.0, 1.0
	builder.AddProperty("ratio", NewNumber("Ratio", false, &minRatio, &maxRatio, 0.5))

	// Adding a number property with range widget and multipleOf
	multipleOfTen := 10.0
	builder.AddProperty("integerRangeSteps", NewNumberWithWidget("Integer Range Steps", false, nil, nil, &multipleOfTen, 50, "range", nil))
	builder.AddUIProperty("integerRangeSteps", UIProperty{Widget: "range"})

	// Adding nested properties

	nestedProps := CreateNestedProperties(
		NestedProperty{"email", NewEmail("Email")},
		NestedProperty{"name", NewString("test", true, 0, 100, "Aidan")},
	)
	builder.AddNestedProperties("nestedObject", nestedProps)

	// Adding array property with nested items
	builder.AddArrayProperty("tasks", map[string]Property{
		"title": NewString("Task Title", true, 0, 0, ""),
		"done":  NewNumber("Done", false, nil, nil, 0),
	})

	builder.AddProperty("uri", NewURI("URI"))

	// password
	builder.AddProperty("pass1", NewPassword("Password", 8))
	builder.AddUIProperty("pass1", UIProperty{Widget: "password", Options: map[string]interface{}{"help": "Hint: Make it strong!"}})
	builder.AddUIOrder([]string{"exampleString", "exampleNumber", "exampleEnumString", "ratio", "integerRangeSteps", "nestedObject", "tasks"})
	buildAndDump(builder)
}

// test case for IF THEN
func Test_if(t *testing.T) {

	builder := NewSchemaBuilder().
		SetProperty("holidayType", Property{Type: "string", Enum: []string{"snow", "beach"}})

	newCon := make(map[string]Property)
	aa := 8.0
	newCon["a"] = NewNumber("want a new aa", false, nil, &aa, 1)
	newCon["b"] = NewNumber("want a new bb", false, nil, &aa, 1)

	builder.AddCondition(ConditionalStructure{
		If: Condition{
			Properties: map[string]Property{"holidayType": {Const: "beach"}},
		},
		Then: Condition{
			Properties: newCon,
		},
	})

	newCon2 := make(map[string]Property)

	newCon2["a"] = NewString("want a new aa str", false, 1, 100, "hey")
	newCon2["bb"] = newCon["b"]

	// Dynamic condition for 'snow' holiday
	builder.AddCondition(ConditionalStructure{
		If: Condition{
			Properties: map[string]Property{"holidayType": {Const: "snow"}},
		},
		Then: Condition{
			Properties: newCon2,
		},
	})
	buildAndDump(builder)

}
