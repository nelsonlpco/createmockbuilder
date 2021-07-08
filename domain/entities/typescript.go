package entities

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/nelsonlpco/createmockbuilder/domain/utils"
)

const ClassTemplate = "class %v {\n%v\n%v\n%v\n%v\n%v\n}\n"
const ConstructorTemplate = "\tconstructor() {\n%v\t}\n"
const ParamsTemplate = "\t%v: %v;\n"
const ParamsWithValueTemplate = "\t\tthis.%v = %v;\n"
const WithMethodTemplate = "\twith%v(value: %v) {\n\t\tthis.%v = value;\n\t\treturn this;\n\t}\n"
const WithDefaultValuesTemplate = "\twithDefaultValues() {\n%v\n\t\treturn this;\n\t}\n"
const BuildMethodTemplate = "\tbuild() {\n\t\treturn {\n%v\n\t\t}\n\t}\n"

type TypescriptBuilder struct {
	Name         string
	BuilderClass string
}

func NewTypescriptBuilder(classBuilder ClassBuilder) *TypescriptBuilder {
	params := makeParams(classBuilder, false)
	constructor := fmt.Sprintf(ConstructorTemplate, makeParams(classBuilder, true))
	withMethods := makeWithMethod(classBuilder)
	withDefaultValues := makeWithDefaultMethod(classBuilder)
	buildMethods := makeBuildMethod(classBuilder)

	builder := TypescriptBuilder{
		Name:         classBuilder.Name(),
		BuilderClass: fmt.Sprintf(ClassTemplate, classBuilder.Name(), params, constructor, withMethods, withDefaultValues, buildMethods),
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
		return "any[]"
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

func makeParams(model ClassBuilder, withValue bool) (result string) {
	result = ""

	for _, param := range model.Params() {
		fieldName := utils.ToCamelCase(param.Name)
		fieldType := getValueType(param.ValueType)
		fieldValue := getDefaultValueByType(fieldType)

		if withValue {
			result += fmt.Sprintf(ParamsWithValueTemplate, fieldName, fieldValue)
		} else {
			result += fmt.Sprintf(ParamsTemplate, fieldName, fieldType)
		}
	}

	return
}

func makeWithMethod(model ClassBuilder) (result string) {
	result = ""

	for _, param := range model.Params() {
		fieldName := utils.ToCamelCase(param.Name)
		result += fmt.Sprintf(WithMethodTemplate, fieldName, getValueType(param.ValueType), fieldName)
	}

	return
}

func makeWithDefaultMethod(model ClassBuilder) (result string) {
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

func makeBuildMethod(model ClassBuilder) (result string) {
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
