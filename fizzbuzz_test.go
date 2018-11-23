package count

import "testing"

func TestItShouldCount_1(t *testing.T) {
	r := FizzBuzz(1)
	if r != "1" {
		t.Errorf("Sum was incorrect, got: %s, want: %d.", r, 1)
	}
}
func TestItShouldCount_2(t *testing.T) {
	r := FizzBuzz(2)
	if r != "2" {
		t.Errorf("Sum was incorrect, got: %s, want: %d.", r, 2)
	}
}

func TestItShouldCount_Fizz(t *testing.T) {
	r := FizzBuzz(3)
	if r != "Fizz" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", r, "Fizz")
	}
}
func TestItShouldCount_4(t *testing.T) {
	r := FizzBuzz(4)
	if r != "4" {
		t.Errorf("Sum was incorrect, got: %s, want: %d.", r, 4)
	}
}

func TestItShouldCount_Buzz(t *testing.T) {
	r := FizzBuzz(5)
	if r != "Buzz" {
		t.Errorf("Sum was incorrect, got %s, want: %s", r, "Buzz")
	}
}

func TestItShouldCount_Fizz_when_input_is_6(t *testing.T) {
	r := FizzBuzz(6)
	if r != "Fizz" {
		t.Errorf("Sum was incorrect, got %s, want: %s", r, "Fizz")
	}
}

func TestItShouldCount_Fizz_when_input_is_9(t *testing.T) {
	r := FizzBuzz(9)
	if r != "Fizz" {
		t.Errorf("Sum was incorrect, got %s, want: %s", r, "Fizz")
	}
}

func TestItShouldCount_Buzz_when_input_is_10(t *testing.T) {
	r := FizzBuzz(10)
	if r != "Buzz" {
		t.Errorf("Sum was incorrect, got %s, want: %s", r, "Buzz")
	}
}

func TestItShouldCount_FizzBuzz(t *testing.T) {
	r := FizzBuzz(15)
	if r != "FizzBuzz" {
		t.Errorf("Sum was incorrect, got %s, want: %s", r, "FizzBuzz")
	}
}
