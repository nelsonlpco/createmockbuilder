package usecases

import (
	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/services/classservices"
	javascriptservice "github.com/nelsonlpco/createmockbuilder/domain/services/javascriptservices"
)

func ConvertJsonToJs(path string) ([]entities.ParsedClass, error) {
	goClasses, err := classservices.MapJsonFromClasses(path)

	if err != nil {
		return nil, err
	}

	result := javascriptservice.MapJavascriptBuilderFromGoClasses(goClasses)

	return result, nil
}
