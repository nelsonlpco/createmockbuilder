package javascriptservice

import (
	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/entities/javascript"
)

func MapJavascriptBuilderFromGoClasses(builders []entities.ClassBuilder) []entities.ParsedClass {
	result := make([]entities.ParsedClass, 0)

	for _, builder := range builders {
		result = append(result, *javascript.NewJavascriptBuilder(builder))
	}

	return result
}
