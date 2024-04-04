package utils_test

import (
	"fmt"
	"testing"
	"tracker/pkg/models"
	"tracker/pkg/utils"
)

func TestErrRes(t *testing.T) {
	mockErr := fmt.Errorf("Test Error")
	got := utils.ErrRes(mockErr)
	want := models.GenericPayload{Msg: "Test Error"}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
