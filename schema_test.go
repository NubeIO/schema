package schema

import (
	"encoding/json"
	"os"
	"testing"
)

func PrintJOSN(x interface{}) {
	ioWriter := os.Stdout
	w := json.NewEncoder(ioWriter)
	w.SetIndent("", "    ")
	w.Encode(x)
}

// https://rjsf-team.github.io/react-jsonschema-form/
// put json in here to test

func Test_simple(t *testing.T) {
	builder := NewSchemaBuilder("HEY")

	builder.NewString("exampleString", "Example String", true, 3, 10, "default value")
	builder.NewNumber("exampleNumber", "Example Num", true, nil, nil, 11)

	ui := UI{}
	ui.AddUIOrder([]string{"exampleString", "exampleNumber"})
	out := &Generated{
		Schema: builder.Build(),
		UI:     ui,
	}
	buildAndDump(out)
}

func Test_all(t *testing.T) {
	builder := NewSchemaBuilder("New Demo")

	builder.NewString("exampleString", "Example String", true, 3, 10, "default value")

	// Adding a number property
	builder.NewNumber("exampleString", "Example String", true, nil, nil, 11)

	// Adding an enum string property
	builder.AddProperty("exampleEnumString", NewEnumString("Example Enum String", []string{"a", "b"}, []string{"Option A", "Option B"}))

	// Adding a ratio property
	minRatio, maxRatio := 0.0, 1.0
	builder.AddProperty("ratio", NewNumber("Ratio", false, &minRatio, &maxRatio, 0.5))

	// Adding a number property with range widget and multipleOf
	multipleOfTen := 10.0
	builder.AddProperty("integerRangeSteps", NewNumberWithWidget("Integer Range Steps", false, nil, nil, &multipleOfTen, 50, "range", nil))

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

	ui := UI{}
	ui.AddUIOrder([]string{"exampleString", "exampleNumber", "exampleEnumString", "ratio", "integerRangeSteps", "nestedObject", "tasks"})
	ui.AddUIProperty("integerRangeSteps", UIProperty{Widget: "range"})
	ui.AddUIProperty("pass1", UIProperty{Widget: "password", Options: map[string]interface{}{"help": "Hint: Make it strong!"}})
	out := &Generated{
		Schema: builder.Build(),
		UI:     ui,
	}
	buildAndDump(out)
}

// test case for IF THEN
func Test_if(t *testing.T) {

	builder := NewSchemaBuilder("IF/THEN").
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

	out := &Generated{
		Schema: builder.Build(),
	}
	buildAndDump(out)

}
