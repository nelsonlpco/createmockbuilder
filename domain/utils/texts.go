package utils

import (
	"strings"
)

func ToCamelCase(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.Title(name)

	return strings.ReplaceAll(name, " ", "")
}
