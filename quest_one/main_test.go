package questone

import "testing"

func TestPartOne(t *testing.T) {
	input := "ABBAC"

	expected := 5
	actual := PartOne(input)
	if actual != expected {
		t.Errorf("PartOne(%q) = %d; want %d", input, actual, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := "AxBCDDCAxD"

	expected := 28
	actual := PartTwo(input)
	if actual != expected {
		t.Errorf("PartOne(%q) = %d; want %d", input, actual, expected)
	}
}
