package classservices

import (
	"reflect"

	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/utils"
)

func MapJsonFromClasses(path string) ([]entities.ClassBuilder, error) {
	jsonData, name, err := utils.OpenJson(path)

	if err != nil {
		return nil, err
	}

	return mapClasses(name, jsonData), nil
}

func mapClasses(name string, data map[string]interface{}) (result []entities.ClassBuilder) {
	result = append(result, *entities.NewBuilderClass(name, data))

	for key, value := range data {
		if reflect.ValueOf(value).Kind() == reflect.Map {
			toMerge := mapClasses(key, value.(map[string]interface{}))
			result = append(result, toMerge...)
		}
	}

	return
}
