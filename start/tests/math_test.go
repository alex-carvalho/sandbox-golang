package tests

import (
	"runtime"
	"testing"
)

func TestAverage(t *testing.T) {
	expected := 7.28
	average := Average(7.2, 9.9, 6.1, 5.9)

	if average != expected {
		t.Errorf("Expected %v result %v", expected, average)
	}
}

func TestArch(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Skipping on amd64")
	}

	t.Errorf("Fail if not amd64")
}
