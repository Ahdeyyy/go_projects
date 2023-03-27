package exercises

import (
	"reflect"
	"testing"
)

func TestFilter_1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []int{2, 4, 6, 8, 10}
	got := Filter_nums(input, IsEven)

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []int{1, 3, 5, 7, 9}
	got := Filter_nums(input, IsOdd)

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []int{2, 3, 5, 7}
	got := Filter_nums(input, IsPrime)

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_4(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []int{3, 5, 7}
	got := Filter_nums(input, IsOddPrime)

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_5(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := []int{10, 20}
	got := Filter_nums(input, IsEvenMultipleOf(5))

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_6(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output := []int{15}
	got := Filter_nums(input, IsOddGreaterAndMultipleOf(10, 3))

	if reflect.DeepEqual(got, output) {
		t.Errorf("got %v wanted %v ", got, output)
	}

}
func TestFilter_7(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output1 := []int{9, 15}
	//output2 := []int{6,12}
	greater_than_five := IsGreater(5)
	multiple_of_three := IsMultipleOf(3)
	got1 := FilterByConds(input1, All, IsOdd, greater_than_five, multiple_of_three)
	if reflect.DeepEqual(got1, output1) {
		t.Errorf("got %v wanted %v ", got1, output1)
	}

}
func TestFilter_8(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	output1 := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	//output2 := []int{6,12}
	greater_than_fifteen := IsGreater(15)
	multiple_of_five := IsMultipleOf(5)
	got1 := FilterByConds(input1, Any, IsPrime, greater_than_fifteen, multiple_of_five)
	if reflect.DeepEqual(got1, output1) {
		t.Errorf("got %v wanted %v ", got1, output1)
	}

}
