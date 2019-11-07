package finite_field_cp

func gf282gf222(a uint8) uint8 {
	/*
	[[0 1 0 0 0 1 1 0]
	 [1 1 0 1 0 1 1 0]
	 [0 1 1 0 1 1 1 1]
	 [1 1 0 1 0 0 0 0]
	 [1 0 0 0 1 1 0 1]
	 [1 0 0 0 1 0 0 1]
	 [1 1 1 0 1 1 0 0]
	 [0 0 0 0 0 0 0 1]]
	 */
	var aBit [8]uint8
	var resultList [8]uint8
	var result uint8 = 0x00
	// means the input is x = (x0, ..., x7)
	for i := 0; i < 8; i++ {
		aBit[7 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[1] ^ aBit[3] ^ aBit[4] ^ aBit[5] ^ aBit[6]
	resultList[1] = aBit[0] ^ aBit[1] ^ aBit[2] ^ aBit[3] ^ aBit[6]
	resultList[2] = aBit[2] ^ aBit[6]
	resultList[3] = aBit[1] ^ aBit[3]
	resultList[4] = aBit[2] ^ aBit[4] ^ aBit[5] ^ aBit[6]
	resultList[5] = aBit[0] ^ aBit[1] ^ aBit[2] ^ aBit[4] ^ aBit[6]
	resultList[6] = aBit[0] ^ aBit[1] ^ aBit[2]
	resultList[7] = aBit[2] ^ aBit[4] ^ aBit[5] ^ aBit[7]
	for i := 0; i < 8; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func gf2222gf28(a uint8) uint8 {
	/*
	[[0 1 1 1 1 1 1 1]
	 [1 1 0 1 0 0 0 0]
	 [1 1 0 1 1 0 1 1]
	 [1 0 1 1 1 1 1 1]
	 [0 1 1 1 1 0 1 0]
	 [0 0 0 0 1 1 0 0]
	 [0 1 0 1 1 1 0 0]
	 [0 0 0 0 0 0 0 1]]
	 */
	var aBit [8]uint8
	var resultList [8]uint8
	var result uint8 = 0x00
	// means the input is x = (x0, ..., x7)
	for i := 0; i < 8; i++ {
		aBit[7 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[1] ^ aBit[2] ^ aBit[3]
	resultList[1] = aBit[0] ^ aBit[1] ^ aBit[2] ^ aBit[4] ^ aBit[6]
	resultList[2] = aBit[0] ^ aBit[3] ^ aBit[4]
	resultList[3] = aBit[0] ^ aBit[1] ^ aBit[2] ^ aBit[3] ^ aBit[4] ^ aBit[6]
	resultList[4] = aBit[0] ^ aBit[2] ^ aBit[3] ^ aBit[4] ^ aBit[5] ^ aBit[6]
	resultList[5] = aBit[0] ^ aBit[3] ^ aBit[5] ^ aBit[6]
	resultList[6] = aBit[0] ^ aBit[2] ^ aBit[3] ^ aBit[4]
	resultList[7] = aBit[0] ^ aBit[2] ^ aBit[3] ^ aBit[7]
	for i := 0; i < 8; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func mulGF22(a, b uint8) uint8 {
	var result uint8
	a0 := a & 0x02 >> 1
	a1 := a & 0x01 >> 0
	b0 := b & 0x02 >> 1
	b1 := b & 0x01 >> 0
	c0 := (a0 ^ a1) & (b0 ^ b1) ^ (a1 & b1)
	c1 := a0 & b0 ^ a1 & b1
	result = c0 << 1 | c1
	return result
}
/*
func mulGF24(a, b uint8) uint8 {
	var result uint8
	a = gf242gf22(a)
	b = gf242gf22(b)
	a0 := a & 0x08 >> 3
	a1 := a & 0x04 >> 2
	a2 := a & 0x02 >> 1
	a3 := a & 0x01 >> 0
	b0 := b & 0x08 >> 3
	b1 := b & 0x04 >> 2
	b2 := b & 0x02 >> 1
	b3 := b & 0x01 >> 0

	return result
}

 */


func InverseDirect(a uint8) uint8 {
	var result uint8
	a = gf282gf222(a)
	return result
}



