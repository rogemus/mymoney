package testing

import (
	"reflect"
	"testing"
)

func AssertEqualInt(t testing.TB, got, want int) {
	t.Helper()

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func AssertSliceOfStructs[T any](t testing.TB, got, want []T) {
	t.Helper()

	emptyArray := len(got) == 0 && len(want) == 0
	deepEqual := reflect.DeepEqual(got, want)

	if !emptyArray && !deepEqual {
		t.Fatalf("want %v, got %v", want, got)
	}
}
