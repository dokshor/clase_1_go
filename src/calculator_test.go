package calculator_test

import (
	calculator "calculator/src"
	"testing"
)

func TestSumar(t *testing.T) {
	got := calculator.Sumar(10, 5)
	want := 15

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRestar(t *testing.T) {
	got := calculator.Restar(10, 5)
	want := 5

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDividir(t *testing.T) {
	got := calculator.Division(10, 5)
	want := 2

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
