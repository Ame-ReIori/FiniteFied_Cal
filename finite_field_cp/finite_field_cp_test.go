package finite_field_cp

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	var a uint8 = 0b10110111
	var b uint8 = 0b10100101
	fmt.Printf("0b%b\n", Multiply(a, b))
}

func TestInverse(t *testing.T) {
	var a uint8 = 0b10100111
	fmt.Printf("0b%b\n", Inverse(a))
}