package utils

import (
	"fmt"
	"testing"
	"tracker/pkg/models"
)

func TestErrRes(t *testing.T) {
	mockErr := fmt.Errorf("Test Error")
	got := ErrRes(mockErr)
	want := models.GenericPayload{Msg: "Test Error"}

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
