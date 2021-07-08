package entities

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nelsonlpco/createmockbuilder/domain/utils"
)

func Test_ShouldBeReturnAvalidClassMapFromJsonData(t *testing.T) {
	json, name, err := utils.OpenJson("../../mocks/people.json")
	name = strings.Title(name)
	expectedParams := []string{"id", "active", "age", "name"}

	expectedClassName := "peopleBuilder"

	if err != nil {
		panic(err)
	}

	classModel := NewBuilderClass(name, json)

	if classModel.name != "PeopleBuilder" {
		t.Errorf("epected %v but receveid %v", expectedClassName, classModel.name)
	}

	if len(classModel.params) != 4 {
		t.Errorf("expected %v but receveid %v", 4, len(classModel.params))
	}

	for _, param := range classModel.params {
		result := false
		for _, expectedParam := range expectedParams {
			result = param.Name == expectedParam
			if result {
				break
			}
		}

		if !result {
			t.Errorf("param %v not found", param.Name)
		}
	}

	if classModel.params[0].ValueType != reflect.Float64 {
		t.Errorf("expected %v but receveid %v", "number", classModel.params[0].ValueType)
	}

	if classModel.params[0].Value.(float64) != 1 {
		t.Errorf("expected %v but receveid %v", 1, classModel.params[0].Value)
	}
}
