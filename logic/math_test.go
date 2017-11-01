package logic

import "testing"

func TestAdd(t *testing.T) {
	a := 200.0
	b := 100.0
	expected := 300.0
	result := Math{}.Addition(a, b)
	if result != expected {
		t.Errorf("Expected %s=%f but found =%f", "Add return=", expected, result)
	}
}

func TestSub(t *testing.T) {
	a := 200.0
	b := 100.0
	expected := 100.0
	result := Math{}.Subtraction(a, b)
	if result != expected {
		t.Errorf("Expected %s=%f but found =%f", "Sub return=", expected, result)
	}
}

func TestMul(t *testing.T) {
	a := 200.0
	b := 100.0
	expected := 20000.0
	result := Math{}.Multiplication(a, b)
	if result != expected {
		t.Errorf("Expected %s=%f but found =%f", "Mul return=", expected, result)
	}
}

func TestDiv(t *testing.T) {
	a := 200.0
	b := 100.0
	expected := 2.0
	result := Math{}.Division(a, b)
	if result != expected {
		t.Errorf("Expected %s=%f but found =%f", "Sub return=", expected, result)
	}
}
