package typescriptservice

import "github.com/nelsonlpco/createmockbuilder/domain/entities"

func MapTypescriptBuilderFromGoClasses(builders []entities.ClassBuilder) []entities.ParsedClass {
	result := make([]entities.ParsedClass, 0)

	for _, builder := range builders {
		result = append(result, *entities.NewTypescriptBuilder(builder))
	}

	return result
}
