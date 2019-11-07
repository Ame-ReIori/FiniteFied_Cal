package finite_field_cp

import (
	"fmt"
	"os"
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

func TestEveryFunction(t *testing.T) {
	var a uint8
	var b uint8
	var temp string
	f, _ := os.Create("function_test")
	defer f.Close()
	temp = fmt.Sprintln("doubleMul")
	_, _ = f.WriteString(temp)
	for a = 0; a < 0x10; a++ {
		temp = fmt.Sprintf("%04b --> %04b\n", a, doubleMulGF24(a))
		_, _ = f.WriteString(temp)
	}

	temp = fmt.Sprintln("inverse")
	_, _ = f.WriteString(temp)
	for a = 0; a < 0x10; a++ {
		temp = fmt.Sprintf("%04b --> %04b\n", a, InverseGF24(a))
		_, _ = f.WriteString(temp)
	}

	temp = fmt.Sprintln("mul")
	_, _ = f.WriteString(temp)
	for a = 0; a < 0x10; a++ {
		for b = 0; b < 0x10; b++ {
			temp = fmt.Sprintf("%04b*%04b --> %04b\n", a, b, multiplyGF24(a, b))
			_, _ = f.WriteString(temp)
		}
	}
}