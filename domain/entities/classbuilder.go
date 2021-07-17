package entities

import (
	"fmt"
	"reflect"

	"github.com/nelsonlpco/createmockbuilder/domain/utils"
)

const Suffix = "Builder"

type ParamModel struct {
	Name      string
	ValueType reflect.Kind
	Value     interface{}
}

type ClassBuilder struct {
	name   string
	params []ParamModel
}

type ParsedClass struct {
	Name         string
	BuilderClass string
}

func (c *ClassBuilder) Name() string {
	return c.name
}

func (c *ClassBuilder) Params() []ParamModel {
	return c.params
}

func (c *ClassBuilder) createClass(name string, data map[string]interface{}) {
	c.name = fmt.Sprintf("%v%v", utils.ToCamelCase(name), Suffix)
	c.params = []ParamModel{}

	for key, value := range data {
		valueType := reflect.ValueOf(value).Kind()

		param := ParamModel{
			Name:      key,
			ValueType: valueType,
			Value:     value,
		}
		c.params = append(c.params, param)
	}
}

func NewBuilderClass(name string, data map[string]interface{}) *ClassBuilder {
	newClass := new(ClassBuilder)
	newClass.createClass(name, data)

	return newClass
}
