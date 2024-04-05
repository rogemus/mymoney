package utils

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJson(t testing.TB, got, want string) {
	t.Helper()

	var gotI, wantI interface{}
	json.Unmarshal([]byte(got), &gotI)
	json.Unmarshal([]byte(want), &wantI)

	if !reflect.DeepEqual(gotI, wantI) {
		t.Errorf("\n want: %v \n got: %v", wantI, gotI)
	}
}

func AssertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\n want: %v \n got: %v", want, got)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("\n want: %v \n got: %v", want, got)
	}
}

func AssertStruct[T any](t testing.TB, got, want T) {
	t.Helper()

	if !cmp.Equal(got, want) {
		t.Errorf("\n want: %v \n got: %v", want, got)
	}
}

func AssertSliceOfStructs[T any](t testing.TB, got, want []T) {
	t.Helper()

	emptyArray := len(got) == 0 && len(want) == 0
	deepEqual := reflect.DeepEqual(got, want)

	if !emptyArray && !deepEqual {
		t.Errorf("\n want: %v \n got: %v", want, got)
	}
}
