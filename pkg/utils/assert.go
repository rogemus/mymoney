package utils

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func AssertEqualInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func AssertStruct[T any](t testing.TB, got, want T) {
	t.Helper()

	if !cmp.Equal(got, want) {
		t.Fatalf("want %v, got %v", want, got)
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
