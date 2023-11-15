package calculator_test

import (
	"reflect"
	"testing"

	"github.com/LeoCBS/bazel-workshop/calculator"
)

func TestSumInt(t *testing.T) {
	result := calculator.SumInt(2, 2)
	expectedResult
	assertEqual(t, expectedResult, result)
}

func assertEqual(t *testing.T, obj1 interface{}, obj2 interface{}) {
	t.Helper()
	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("obj {%v} not equal to obj {%v}", obj1, obj2)
	}
}
