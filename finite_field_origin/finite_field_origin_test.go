package finite_field_origin

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	var a uint8 = 0b10010011
	var b uint8 = 0b10101001
	var result uint8 = Multiply(a, b)
	fmt.Printf("0x%X", result)
}