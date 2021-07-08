package ui

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nelsonlpco/createmockbuilder/domain/usecases"
)

func CliCreateMockBuilder() {
	jsonPath := flag.String("f", "./sample.json", "arquivo JSON")
	outputPath := flag.String("o", ".", "diretório de saida")
	js := flag.Bool("js", false, "Template javascript o padrão é Typescript")
	extension := "ts"

	flag.Parse()

	if *js {
		extension = "js"
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

	tsClasses, err := usecases.ConvertJsonToTs(*jsonPath)

	if err != nil {
		panic(err)
	}

	for _, tsClass := range tsClasses {
		file := fmt.Sprintf("%v/%v.%v", *outputPath, tsClass.Name, extension)
		err := ioutil.WriteFile(file, []byte(tsClass.BuilderClass), 0644)

		if err != nil {
			panic(err)
		}

		fmt.Println(file)
	}

	fmt.Printf("-------------------\n\n")
	fmt.Println("Builders criados com sucesso.")

}
