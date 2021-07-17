package ui

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nelsonlpco/createmockbuilder/domain/entities"
	"github.com/nelsonlpco/createmockbuilder/domain/usecases"
)

func CliCreateMockBuilder() {
	jsonPath := flag.String("f", "./sample.json", "arquivo JSON")
	outputPath := flag.String("o", ".", "diretório de saida")
	template := flag.String("t", "ts", "Template (ts/js) typescript/javascript")
	extension := "ts"
	var classes []entities.ParsedClass
	var err error

	flag.Parse()

	if *template != "ts" && *template != "js" {
		extension = "ts"
	} else {
		extension = *template
	}

	if *jsonPath == "" {
		panic("-f é obrigatório, execute com -h para ajuda.")
	}

	fmt.Println("Criando builders...")
	fmt.Printf("-------------------\n\n")

	mkdirErr := os.Mkdir(*outputPath, 0755)

	if mkdirErr != nil {
		fmt.Println(mkdirErr)
	}

	if extension == "js" {
		classes, err = usecases.ConvertJsonToJs(*jsonPath)
	} else {
		classes, err = usecases.ConvertJsonToTs(*jsonPath)
	}

	if err != nil {
		panic(err)
	}

	for _, buildClass := range classes {
		file := fmt.Sprintf("%v/%v.%v", *outputPath, buildClass.Name, extension)
		err := ioutil.WriteFile(file, []byte(buildClass.BuilderClass), 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println(file)
	}

	fmt.Printf("-------------------\n\n")
	fmt.Println("Builders criados com sucesso.")

}
