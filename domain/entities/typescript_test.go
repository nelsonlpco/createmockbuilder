package entities

import (
	"fmt"
	"testing"
)

func Test_ShouldBeReturnAValidTypescriptStringClass(t *testing.T) {
	params := "\tId: number;\n"
	constructor := fmt.Sprintf(ConstructorTemplate, "\t\tthis.Id = 0;\n")
	withId := fmt.Sprintf(WithMethodTemplate, "Id", "number", "Id")
	withMethods := fmt.Sprintf("%v", withId)
	withDefaultValues := fmt.Sprintf(WithDefaultValuesTemplate, "\t\tthis.Id = 1;\n")
	build := fmt.Sprintf(BuildMethodTemplate, "\t\t\tid: this.Id,\n")

	expectedValue := fmt.Sprintf(ClassTemplate, "PeopleBuilder", params, constructor, withMethods, withDefaultValues, build)

	mock := map[string]interface{}{
		"id": 1,
	}

	userModel := NewBuilderClass("people", mock)

	ts := NewTypescriptBuilder(*userModel)

	if ts.BuilderClass != expectedValue {
		t.Errorf("Expected:\n%v \n receveid:\n%v\n", expectedValue, ts.BuilderClass)
	}
}
