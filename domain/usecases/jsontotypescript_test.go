package usecases

import "testing"

func Test_ShouldBeReturnAvalidClassModelFromJson(t *testing.T) {
	result, err := ConvertJsonToTs("../../mocks/user.json")

	if err != nil {
		t.Errorf("Unexpected Error received: %v", err.Error())
	}

	if len(result) < 4 {
		t.Errorf("expected %v classes but receveid %v", 4, len(result))
	}
}
