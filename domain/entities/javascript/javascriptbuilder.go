package javascript

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/utils"
)

const ClassTemplate = "class %v {\n%v\n%v\n%v\n%v\n}\n"
const ConstructorTemplate = "\tconstructor() {\n%v\t}\n"
const ParamsWithValueTemplate = "\t\tthis.%v = %v;\n"
const WithMethodTemplate = "\twith%v(%v) {\n\t\tthis.%v = %v;\n\t\treturn this;\n\t}\n"
const WithDefaultValuesTemplate = "\twithDefaultValues() {\n%v\n\t\treturn this;\n\t}\n"
const BuildMethodTemplate = "\tbuild() {\n\t\treturn {\n%v\n\t\t}\n\t}\n"

func NewJavascriptBuilder(classBuilder entities.ClassBuilder) *entities.ParsedClass {
	constructor := fmt.Sprintf(ConstructorTemplate, makeParams(classBuilder))
	withMethods := makeWithMethod(classBuilder)
	withDefaultValues := makeWithDefaultMethod(classBuilder)
	buildMethods := makeBuildMethod(classBuilder)

	builder := entities.ParsedClass{
		Name:         classBuilder.Name(),
		BuilderClass: fmt.Sprintf(ClassTemplate, classBuilder.Name(), constructor, withMethods, withDefaultValues, buildMethods),
	}

	return &builder
}

func getDefaultValueByType(kind string) interface{} {
	switch kind {
	case "number":
		return 0
	case "boolean":
		return false
	case "string":
		return "\"\""
	case "any[]":
		return "[]"
	default:
		return "{}"
	}
}

func getValueType(kind reflect.Kind) string {
	switch kind {
	case reflect.Slice, reflect.Array:
		return "[]"
	case reflect.Float32, reflect.Float64, reflect.Int:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.String:
		return "string"
	default:
		return "any"
	}
}

func makeParams(model entities.ClassBuilder) (result string) {
	result = ""

	for _, param := range model.Params() {
		fieldName := utils.ToCamelCase(param.Name)
		fieldType := getValueType(param.ValueType)
		fieldValue := getDefaultValueByType(fieldType)

		result += fmt.Sprintf(ParamsWithValueTemplate, fieldName, fieldValue)
	}

	return
}

func makeWithMethod(model entities.ClassBuilder) (result string) {
	result = ""

	for _, param := range model.Params() {
		fieldName := utils.ToCamelCase(param.Name)
		result += fmt.Sprintf(WithMethodTemplate, fieldName, fieldName, fieldName, fieldName)
	}

	return
}

func makeWithDefaultMethod(model entities.ClassBuilder) (result string) {
	params := ""

	for _, param := range model.Params() {
		fieldName := utils.ToCamelCase(param.Name)
		fieldType := getValueType(param.ValueType)
		fieldValue := param.Value

		if fieldType == "string" {
			fieldValue = fmt.Sprintf("\"%v\"", fieldValue)
		}

		if fieldType == "any" || fieldType == "any[]" {
			fieldValue = getDefaultValueByType(fieldType)
		}

		params += fmt.Sprintf(ParamsWithValueTemplate, fieldName, fieldValue)
	}

	result = fmt.Sprintf(WithDefaultValuesTemplate, params)

	return
}

func makeBuildMethod(model entities.ClassBuilder) (result string) {
	jsonParams := ""

	for _, param := range model.Params() {
		template := "\t\t\t%v: this.%v,\n"
		if strings.Contains(param.Name, "-") || strings.Contains(param.Name, "_") {
			template = "\t\t\t\"%v\": this.%v,\n"
		}
		jsonParams += fmt.Sprintf(template, param.Name, utils.ToCamelCase(param.Name))
	}

	result = fmt.Sprintf(BuildMethodTemplate, jsonParams)

	return
}
