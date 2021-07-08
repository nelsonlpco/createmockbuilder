package typescriptservice

import "github.com/nelsonlpco/createmockbuilder/domain/entities"

func MapTypescriptBuilderFromGoClasses(builders []entities.ClassBuilder) []entities.TypescriptBuilder {
	result := make([]entities.TypescriptBuilder, 0)

	for _, builder := range builders {
		result = append(result, *entities.NewTypescriptBuilder(builder))
	}

	return result
}
