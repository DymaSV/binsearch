package binsearch

import (
	"testing"
)

// TestBinSearch calls binsearch.BinSearch with an array.
//Try to find last value
func TestBinSearch_LastValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 5
	result := BinSearch(&arr, 13, 0)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Try to find first value
func TestBinSearch_FirstValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 0
	result := BinSearch(&arr, 2, 0)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Try to find middle value
func TestBinSearch_MiddleValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 2
	result := BinSearch(&arr, 5, 0)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Did not find value
func TestBinSearch_NotFindValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := -1
	result := BinSearch(&arr, 15, 0)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Did not find value in the middle of array
func TestBinSearch_NotFindValueInMiddle(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := -1
	result := BinSearch(&arr, 1, 0)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}
