package test

import (
	"reflect"
	"testing"
)

func AssertNil(t *testing.T, got interface{}) {
	if got != nil {
		t.Errorf("want nil, got %s", got)
	}
}

func AssertEquals(t *testing.T, got interface{}, want interface{}) {
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func AssertDeepEquals(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	if got != true {
		t.Errorf("got %t, want %t", got, true)
	}
}
