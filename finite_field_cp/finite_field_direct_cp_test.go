package finite_field_cp

import (
	"fmt"
	"os"
	"testing"
)

func TestAcc(t *testing.T) {

	var a uint8

	for a = 0x00; a <= 0x0f; a++ {
		a1 := a & 0x08 >> 3
		a2 := a & 0x04 >> 2
		a3 := a & 0x02 >> 1
		a4 := a & 0x01 >> 0
		t1 := a1 ^ a2 ^ a3 ^ a1&a2 ^ a1&a3 ^ a1&a4 ^ a1&a2&a3
		t2 := a1 ^ a2 ^ a1&a4 ^ a2&a4 ^ a3&a4 ^ a1&a2&a4
		t3 := a1 ^ a1&a3 ^ a2&a3 ^ a2&a4 ^ a3&a4 ^ a1&a3&a4
		t4 := a1 ^ a2 ^ a3 ^ a4 ^ a2&a3 ^ a2&a4 ^ a1&a2&a3 ^ a2&a3&a4
		t := t1 << 3 | t2 << 2 | t3 << 1 | t4
		if t != InverseGF24(a) {
			fmt.Printf("%b, %b\n", t, InverseGF24(a))
			os.Exit(1)
		}
	}
	fmt.Println("pass")
}
