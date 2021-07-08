package usecases

import (
	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/services/classservices"
	typescriptservice "github.com/nelsonlpco/createmockbuilder/domain/services/typescriptservices"
)

func ConvertJsonToTs(path string) ([]entities.TypescriptBuilder, error) {
	goClasses, err := classservices.MapJsonFromClasses(path)

	if err != nil {
		return nil, err
	}

	result := typescriptservice.MapTypescriptBuilderFromGoClasses(goClasses)

	return result, nil
}
