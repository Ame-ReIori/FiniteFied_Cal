/*
	For easy implementation on FPGA, do something special
 */
package finite_field_cp

func gf282gf24(a uint8) uint8 {
	/*
		input: x[0...7]
		output: y[0...7]
	 */
	// split 1 byte to 8 bits
	/*
	matrix := [[0, 0, 1, 0, 0, 1, 0, 0],
	           [1, 1, 1, 1, 0, 1, 0, 0],
			   [0, 1, 0, 0, 1, 0, 0, 1],
			   [1, 1, 1, 1, 0, 0, 0, 0],
			   [1, 1, 0, 0, 1, 1, 1, 1],
			   [1, 1, 0, 0, 1, 1, 0, 1],
			   [1, 0, 0, 0, 1, 1, 1, 0],
			   [0, 0, 0, 0, 0, 0, 0, 1]]
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
	resultList[1] = aBit[1] ^ aBit[2] ^ aBit[3] ^ aBit[4] ^ aBit[5]
	resultList[2] = aBit[0] ^ aBit[1] ^ aBit[3]
	resultList[3] = aBit[1] ^ aBit[3]
	resultList[4] = aBit[2] ^ aBit[4] ^ aBit[5] ^ aBit[6]
	resultList[5] = aBit[0] ^ aBit[1] ^ aBit[4] ^ aBit[5] ^ aBit[6]
	resultList[6] = aBit[4] ^ aBit[6]
	resultList[7] = aBit[2] ^ aBit[4] ^ aBit[5] ^ aBit[7]
	for i := 0; i < 8; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func gf242gf28(a uint8) uint8 {
	/*
	input: x[0...7]
	output: y[0...7]
	*/
	/*
	matrix := [[0, 1, 1, 1, 0, 1, 0, 0],
	           [0, 0, 0, 0, 1, 0, 1, 1],
	           [1, 1, 0, 1, 0, 0, 0, 0],
	           [1, 0, 1, 1, 1, 1, 1, 1],
	           [0, 0, 1, 0, 1, 0, 1, 0],
	           [0, 1, 0, 1, 0, 0, 0, 0],
	           [0, 0, 0, 0, 1, 1, 0, 0],
	           [0, 0, 0, 0, 0, 0, 0, 1]]
	*/
	var aBit [8]uint8
	var resultList [8]uint8
	var result uint8 = 0x00
	// means the input is x = (x0, ..., x7)
	for i := 0; i < 8; i++ {
		aBit[7 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[2] ^ aBit[3]
	resultList[1] = aBit[0] ^ aBit[2] ^ aBit[5]
	resultList[2] = aBit[0] ^ aBit[3] ^ aBit[4]
	resultList[3] = aBit[0] ^ aBit[2] ^ aBit[3] ^ aBit[5]
	resultList[4] = aBit[1] ^ aBit[3] ^ aBit[4] ^ aBit[6]
	resultList[5] = aBit[0] ^ aBit[3] ^ aBit[6]
	resultList[6] = aBit[1] ^ aBit[3] ^ aBit[4]
	resultList[7] = aBit[1] ^ aBit[3] ^ aBit[7]
	for i := 0; i < 8; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func gf242gf22(a uint8) uint8 {
	/*
	matrix = [[1, 1, 1, 0],
			  [0, 1, 1, 0],
			  [0, 1, 0, 0],
			  [0, 0, 0, 1]]
	 */
	var aBit [4]uint8
	var resultList [4]uint8
	var result uint8 = 0x00
	// means the input is x = (x0, ..., x7)
	for i := 0; i < 4; i++ {
		aBit[3 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[0]
	resultList[1] = aBit[0] ^ aBit[1] ^ aBit[2]
	resultList[2] = aBit[0] ^ aBit[1]
	resultList[3] = aBit[3]
	for i := 0; i < 4; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func gf222gf24(a uint8) uint8 {
	/*
		matrix = [[1, 1, 0, 0],
				  [0, 0, 1, 0],
				  [0, 1, 1, 0],
				  [0, 0, 0, 1]]
	*/
	var aBit [4]uint8
	var resultList [4]uint8
	var result uint8 = 0x00
	// means the input is x = (x0, ..., x7)
	for i := 0; i < 4; i++ {
		aBit[3 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[0]
	resultList[1] = aBit[0] ^ aBit[2]
	resultList[2] = aBit[1] ^ aBit[2]
	resultList[3] = aBit[3]
	for i := 0; i < 4; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func multiplyGF22(a, b uint8) uint8 {
	var result uint8
	a1 := (a & 0x02) >> 1
	a0 := a & 0x01
	b1 := (b & 0x02) >> 1
	b0 := b & 0x01
	result1 := ((a1 ^ a0) & (b1 ^ b0)) ^ (a0 & b0)
	result0 := a1 & b1 ^ a0 & b0
	result = result1 << 1 | result0
	return result
}

func inverseGF22(a uint8) uint8 {
	var result uint8
	a1 := (a & 0x02) >> 1
	a0 := a & 0x01
	result = (a1 << 1) | (a0 ^ a1)
	return result
}

func multiplyGF24ByGF22(a, b uint8) uint8 {
	a = gf242gf22(a)
	b = gf242gf22(b)
	var result uint8
	a1 := (a & 0x0c) >> 2
	a0 := a & 0x03
	b1 := (b & 0x0c) >> 2
	b0 := b & 0x03
	temp1 := multiplyGF22(a1 ^ a0, b1 ^ b0) ^ multiplyGF22(a0, b0)
	temp0 := multiplyGF22(a0, b0) ^ multiplyGF22(multiplyGF22(a1, b1), 0x02)
	result = temp1 << 2 | temp0
	result = gf222gf24(result)
	return result
}

func inverseGF24ByGF22(a uint8) uint8 {
	a = gf242gf22(a)
	var result uint8
	a1 := (a & 0x0c) >> 2
	a0 := a & 0x03
	temp1 := multiplyGF22(inverseGF22(multiplyGF22(inverseGF22(a1), 0x02) ^ multiplyGF22(a1, a0) ^ inverseGF22(a0)), a1)
	temp0 := multiplyGF22(inverseGF22(multiplyGF22(inverseGF22(a1), 0x02) ^ multiplyGF22(a1, a0) ^ inverseGF22(a0)), (a0 ^ a1))
	result = temp1 << 2 | temp0
	result = gf222gf24(result)
	return result
}

func multiplyGF24(a, b uint8) uint8 {
	/*
		a = a[0...3]
		b = b[0...3]
		result = result[0...3]
	 */
	var result uint8
	var resultList [4]uint8
	var aBit [4]uint8
	var bBit [4]uint8
	for i := 0; i < 4; i++ {
		aBit[3 - i] = a & 0x01
		bBit[3 - i] = b & 0x01
		a >>= 1
		b >>= 1
	}
	resultList[0] = (bBit[0] & aBit[3]) ^ (bBit[1] & aBit[2]) ^ (bBit[2] & aBit[1]) ^ ((bBit[3] ^ bBit[0]) & aBit[0])
	resultList[1] = (bBit[1] & aBit[3]) ^ (bBit[2] & aBit[2]) ^((bBit[3] ^ bBit[0]) & aBit[1]) ^ ((bBit[0] ^ bBit[1]) & aBit[0])
	resultList[2] = (bBit[2] & aBit[3]) ^ ((bBit[3] ^ bBit[0]) & aBit[2]) ^ ((bBit[0] ^ bBit[1]) & aBit[1]) ^ ((bBit[1] ^ bBit[2]) & aBit[0])
	resultList[3] = (bBit[3] & aBit[3]) ^ (bBit[0] & aBit[2]) ^ (bBit[1] & aBit[1]) ^ (bBit[2] & aBit[0])
	for i := 0; i < 4; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func doubleMulGF24(a uint8) uint8 {
	var result uint8
	var resultList [4]uint8
	var aBit [4]uint8
	for i := 0; i < 4; i++ {
		aBit[3 - i] = a & 0x01
		a >>= 1
	}
	resultList[0] = aBit[0]
	resultList[1] = aBit[0] ^ aBit[2]
	resultList[2] = aBit[1]
	resultList[3] = aBit[1] ^ aBit[3]
	for i := 0; i < 4; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func InverseGF24(a uint8) uint8 {

	var result uint8 = 0x00
	var resultList [4]uint8
	var aBit [4]uint8
	for i := 0; i < 4; i++ {
		aBit[3 - i] = a & 0x01
		a >>= 1
	}
	// Important!!!!
	// (^a & 0x01) means ~a in general
	// because in golang ^a means inverse by bit
	// so for a = 0x01 ==> ^a = 0xfe
	// but in this situation, we only need the last bit
	// so calculate a & 0x01
	// in verilog, we can inverse directly.(I guess -v- )

	/*
		the program is complicated.
		maybe it is clear
	in = (in0, in1, in2, in3)
	out = (out0, out1, out2, out3)

	out0 =
	~in0 ~in1 in2 + ~in0 in1 ~in2 + in0 ~in1 ~in3 +
	in0 ~in2 ~in3 + in0 in1 in2 in3

	out1 =
	~in0 in2 in3 + ~in0 in1 ~in3 + in0 ~in1 ~in3 +
	in0 ~in1 in2 + in0 in1 ~in2 in3

	out2 =
	~in0 in2 in3 + ~in0 in1 in3 + ~in0 in1 in2 +
	in0 ~in1 ~in2 + in0 in1 ~in3

	out3 =
	~in0 ~in2 in3 + ~in0 in2 ~in3 + ~in0 in1 ~in2 +
	in1 in2 ~in3 + in0 ~in1 ~in2 ~in3 + in0 ~in1 in2 in3

	*/
	resultList[0] = (^aBit[0] & 0x01) & (aBit[1] ^ aBit[2]) | aBit[0] & ((^(aBit[3] | aBit[1] & aBit[2]) & 0x01)| aBit[1] & aBit[2] & aBit[3])
	resultList[1] = (^aBit[0] & 0x01) & aBit[2] & aBit[3] | (aBit[0] ^ aBit[1]) & (^aBit[3] & 0x01) | aBit[0] & (^aBit[1] & 0x01) & aBit[2] | aBit[0] & aBit[1] & (^aBit[2] & 0x01) & aBit[3]
	resultList[2] = (^aBit[0] & 0x01) & aBit[2] & aBit[3] | aBit[1] & (aBit[0] ^ aBit[3]) | (^aBit[0] & 0x01) & aBit[1] & aBit[2] | aBit[0] & (^(aBit[1] | aBit[2]) & 0x01)
	resultList[3] = (^aBit[0] & 0x01) & (aBit[2] ^ aBit[3]) | (^aBit[0] & 0x01) & aBit[1] & (^aBit[2] & 0x01) | aBit[1] & aBit[2] & (^aBit[3] & 0x01) | aBit[0] & (^aBit[1] & 0x01) & ((^(aBit[2] | aBit[3]) & 0x01) | aBit[2] & aBit[3])
	for i := 0; i < 4; i++ {
		result <<= 1
		result ^= resultList[i]
	}
	return result
}

func Multiply(a, b uint8) uint8 {
	a = gf282gf24(a)
	b = gf282gf24(b)
	var result uint8
	a1 := a >> 4
	a0 := a & 0x0f
	b1 := b >> 4
	b0 := b & 0x0f
	temp1 := multiplyGF24ByGF22(a1, b0) ^ multiplyGF24ByGF22(a0, b1) ^ multiplyGF24ByGF22(a1, b1)
	temp0 := multiplyGF24ByGF22(a0, b0) ^ multiplyGF24ByGF22(multiplyGF24ByGF22(a1, b1), 0x09)
	result = temp1 << 4 | temp0
	result = gf242gf28(result)
	return result
}

func Inverse(a uint8) uint8 {
	a = gf282gf24(a)
	var result uint8
	a1 := a >> 4
	a0 := a & 0x0f
	temp1 := multiplyGF24ByGF22(inverseGF24ByGF22(multiplyGF24ByGF22(doubleMulGF24(a1), 0x09) ^ multiplyGF24ByGF22(a1, a0) ^ doubleMulGF24(a0)), a1)
	temp0 := multiplyGF24ByGF22(inverseGF24ByGF22(multiplyGF24ByGF22(doubleMulGF24(a1), 0x09) ^ multiplyGF24ByGF22(a1, a0) ^ doubleMulGF24(a0)), (a0 ^ a1))
	result = temp1 << 4 | temp0
	result = gf242gf28(result)
	return result
}
