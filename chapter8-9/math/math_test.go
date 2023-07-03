package math

import "testing"

type avgTestPair struct {
	values []float64
	avg    float64
}

var avgTestCases = []avgTestPair{
	{[]float64{}, 0},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range avgTestCases {
		result := Average(pair.values)
		if result != pair.avg {
			t.Error(
				"For", pair.values,
				"Expected", pair.avg,
				"Got", result,
			)
		}
	}
}

type mathTestPair struct {
	values   []int
	max, min int
}

var mathTestCases = []mathTestPair{
	{[]int{5, 9, 4, 3}, 9, 3},
	{[]int{1, 36, 88, 38}, 88, 1},
	{[]int{20, 94, 51, 76}, 94, 20},
	{[]int{11, 9, 30, 85}, 85, 9},
}

func TestMinAndMax(t *testing.T) {
	for _, pair := range mathTestCases {
		min := FindMin(pair.values...)
		if min != pair.min {
			t.Error(
				"For", pair.values,
				"Expected", pair.min,
				"Got", min,
			)
		}

		max := FindMax(pair.values...)
		if max != pair.max {
			t.Error(
				"For", pair.values,
				"Expected", pair.max,
				"Got", max,
			)
		}
	}
}
