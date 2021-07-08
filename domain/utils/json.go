package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func OpenJson(path string) (map[string]interface{}, string, error) {
	if !strings.Contains(path, ".json") {
		return nil, "", fmt.Errorf("the file %v is not a json", path)
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, "", fmt.Errorf("error on open file %v", path)
	}

	defer file.Close()
	byteValue, readError := ioutil.ReadAll(file)

	if readError != nil {
		return nil, "", fmt.Errorf("error on read file %v", path)
	}

	var parsedJson map[string]interface{}
	errorOnParseJson := json.Unmarshal(byteValue, &parsedJson)

	if errorOnParseJson != nil {
		return nil, "", fmt.Errorf("invalid json file %v", path)
	}

	splitedPath := strings.Split(path, "/")
	fileName := splitedPath[len(splitedPath)-1]

	return parsedJson, strings.Replace(fileName, ".json", "", -1), nil
}
