package main

import (
	"fmt"
	"strings"
	"testing"
)

func compareExpectAndGot(expect, got [][]string, test *testing.T) {
	for i, v := range got {
		gotRow := strings.Join(v, " ")
		expectedRow := strings.Join(expect[i], " ")
		if gotRow != expectedRow {
			test.Fatalf("Not equal.\nExpected:%s\nGot:%s", expectedRow, gotRow)
			// test.Errorf("Not equal.\nExpected:%s\nGot:%s", expectedRow, gotRow)
		}
	}
}
func TestDefaultSort(t *testing.T) {
	fmt.Println(t.Name())
	expectSlice := [][]string{
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"cfdf", "1"},
		{"dsad", "11"},
		{"zfds", "310"},
	}

	sortingObj := SortingObj{rowsSlices: [][]string{
		{"zfds", "310"},
		{"cfdf", "1"},
		{"dsad", "11"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
	}}
	sortingObj.Sort()

	compareExpectAndGot(expectSlice, sortingObj.rowsSlices, t)

}

func TestColumnSort(t *testing.T) {
	fmt.Println(t.Name())
	expectSlice := [][]string{
		{"cfdf", "1"},
		{"dsad", "11"},
		{"zfds", "310"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
	}

	sortingObj := SortingObj{
		rowsSlices: [][]string{
			{"zfds", "310"},
			{"cfdf", "1"},
			{"dsad", "11"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
		}, Sortoptions: Sortoptions{numOfSortColumn: 1},
	}
	sortingObj.Sort()

	compareExpectAndGot(expectSlice, sortingObj.rowsSlices, t)

}

func TestColumnAndNumericSort(t *testing.T) {
	fmt.Println(t.Name())
	expectSlice := [][]string{
		{"cfdf", "1"},
		{"dsad", "11"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"zfds", "310"},
	}

	sortingObj := SortingObj{
		rowsSlices: [][]string{
			{"zfds", "310"},
			{"cfdf", "1"},
			{"dsad", "11"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
		}, Sortoptions: Sortoptions{numOfSortColumn: 1, isNumericSort: true},
	}
	sortingObj.Sort()

	compareExpectAndGot(expectSlice, sortingObj.rowsSlices, t)

}

func TestReverseSort(t *testing.T) {
	fmt.Println(t.Name())
	expectSlice := [][]string{
		{"zfds", "310"},
		{"dsad", "11"},
		{"cfdf", "1"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
		{"aada", "65"},
	}

	sortingObj := SortingObj{
		rowsSlices: [][]string{
			{"zfds", "310"},
			{"cfdf", "1"},
			{"dsad", "11"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
			{"aada", "65"},
		}, Sortoptions: Sortoptions{isReverseSort: true},
	}
	sortingObj.Sort()

	compareExpectAndGot(expectSlice, sortingObj.rowsSlices, t)

}
