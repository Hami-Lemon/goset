package goset

import (
	"reflect"
	"testing"
)

func Equal[T comparable](v1, v2 T, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("assert: %v == %v ? false!", v1, v2)
	}
}

func NotEqual[T comparable](v1, v2 T, t *testing.T) {
	t.Helper()
	if reflect.DeepEqual(v1, v2) {
		t.Errorf("assert: %v != %v ? false!", v1, v2)
	}
}

func Nil[T any](v T, t *testing.T) {
	t.Helper()
	vKind := reflect.ValueOf(v).Kind()
	if vKind != reflect.Invalid {
		t.Errorf("assert: %v == nil ? false!", v)
	}
}

func NotNil[T any](v T, t *testing.T) {
	t.Helper()
	vKind := reflect.ValueOf(v).Kind()
	if vKind == reflect.Invalid {
		t.Errorf("assert: %v != nil ? false!", v)
	}
}

func MapContains[K comparable, V any](m map[K]V, k K, t *testing.T) {
	t.Helper()
	_, ok := m[k]
	if !ok {
		t.Errorf("assert: map:%v contains %v ? false!", m, k)
	}
}

func MapNotContains[K comparable, V any](m map[K]V, k K, t *testing.T) {
	t.Helper()
	_, ok := m[k]
	if ok {
		t.Errorf("assert: map:%v not-contains %v ? false!", m, k)
	}
}

func InSlice[T any](value T, slice []T, t *testing.T) {
	t.Helper()
	for _, v := range slice {
		if reflect.DeepEqual(value, v) {
			return
		}
	}
	t.Errorf("assert: %v in slice:%v ? false !", value, slice)
}
