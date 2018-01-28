package svg

import (
	"testing"
	"math"
	"reflect"
	"fmt"
)

func TestParamLogic(t *testing.T) {
	testParams := map[string][]string {
		"width": {"1"},
		"height": {"2"},
	}
	intResult, floatResult := getParams(testParams)
	expectedIntResult := map[string]int{
		"width": 1,
		"height": 2,
		"cells": 100,
	}
	xyrange := 30.0

	expectedFloatResult := map[string]float64{
		"xyrange": xyrange,
		"xyscale": float64(expectedIntResult["width"]) / 2 / xyrange,
		"zscale": float64(expectedIntResult["height"]) * 0.4,
		"angle": math.Pi / 6,
	}

	if ! reflect.DeepEqual(intResult, expectedIntResult) ||
		! reflect.DeepEqual(floatResult, expectedFloatResult) {
		t.Fail()
		fmt.Println("Not equal: ")
		fmt.Println(intResult)
		fmt.Println(expectedIntResult)
		fmt.Println(floatResult)
		fmt.Println(expectedFloatResult)
	}

}
