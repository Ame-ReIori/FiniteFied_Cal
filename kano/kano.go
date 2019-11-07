package kano

import (
	"finitefield/finite_field_cp"
	"fmt"
)

func printBit(s [16][16]uint8) {
	var i uint8
	var j uint8
	fmt.Printf("    ")
	for j = 0; j < 16; j++ {
		fmt.Printf("  %04b", j)
	}
	fmt.Printf("\n")
	for i = 0; i < 16; i++ {
		fmt.Printf("%04b", i)
		for j = 0; j < 16; j++ {
			fmt.Printf("     %v", s[i][j])
		}
		fmt.Printf("\n")
	}
}

func AllMultiGF24() {
	var i uint8
	var j uint8
	var s0 [16][16]uint8
	var s1 [16][16]uint8
	var s2 [16][16]uint8
	var s3 [16][16]uint8
	for i = 0; i < 16; i++ {
		for j = 0; j < 16; j++ {
			s0[i][j] = finite_field_cp.MultiplyGF24(i, j) >> 3
		}
	}
	for i = 0; i < 16; i++ {
		for j = 0; j < 16; j++ {
			s1[i][j] = finite_field_cp.MultiplyGF24(i, j) << 5 >> 7
		}
	}
	for i = 0; i < 16; i++ {
		for j = 0; j < 16; j++ {
			s2[i][j] = finite_field_cp.MultiplyGF24(i, j) << 6 >> 7
		}
	}
	for i = 0; i < 16; i++ {
		for j = 0; j < 16; j++ {
			s3[i][j] = finite_field_cp.MultiplyGF24(i, j) & 0x01
		}
	}
	printBit(s0)
	fmt.Println()
	printBit(s1)
	fmt.Println()
	printBit(s2)
	fmt.Println()
	printBit(s3)
	fmt.Println()
}