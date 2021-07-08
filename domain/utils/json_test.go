package utils

import (
	"fmt"
	"testing"
)

func Test_ShouldBeReturnAnErrorWhenInputAinvalidJsonFile(t *testing.T) {
	invaliFilepath := "./file.txt"
	expectedError := fmt.Sprintf("the file %v is not a json", invaliFilepath)

	_, _, err := OpenJson(invaliFilepath)

	if err == nil {
		t.Errorf("expected a invalid file error.")
	}

	if err.Error() != expectedError {
		t.Errorf("%v but received %v", err.Error(), expectedError)
	}
}

func Test_ShouldBeReturnAnErrorWhenInputAinvalidFilePath(t *testing.T) {
	invaliFilepath := "./file.json"
	expectedError := fmt.Errorf("error on open file %v", invaliFilepath)

	_, _, err := OpenJson(invaliFilepath)

	if err == nil {
		t.Errorf("Expected a invalid file path")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("%v but received %v", err.Error(), expectedError.Error())
	}
}

func Test_ShouldBeReturnAnErrorWhenInputAinvalidJson(t *testing.T) {
	path := "../../mocks/invalid.json"
	expectedError := fmt.Errorf("invalid json file %v", path)

	_, _, err := OpenJson(path)

	if err == nil {
		t.Errorf("Expected a invalid file path")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("%v but received %v", err.Error(), expectedError.Error())
	}
}

func Test_ShouldBeReturnAvalidJsonData(t *testing.T) {
	path := "../../mocks/people.json"
	expectedName := "people"

	json, name, err := OpenJson(path)

	if err != nil {
		t.Errorf("not expected error but receiveid %v", err.Error())
	}

	if len(json) != 4 {
		t.Errorf("expected 4 elements but received %v", len(json))
	}

	if name != expectedName {
		t.Errorf("expected %v but received %v", expectedName, name)
	}
}
