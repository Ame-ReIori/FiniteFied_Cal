package finite_field_origin

const IRREDUCIBLEPOLY uint16 = 0b111110101

func xtime(x uint8) uint8 {
	var FX uint16 = IRREDUCIBLEPOLY // for golang cannot convert the type of const
	var result uint8 = x << 1
	if x & 0x80 == 0x80 {
		result ^= uint8(FX)
	} else {
		result ^= 0x00
	}
	return result
}

func Multiply(a, b uint8) uint8 {
	var result uint8 = 0x00
	var tempList []uint8
	tempList = append(tempList, a)
	for i := 1; i < 8; i++ {
		tempList = append(tempList, xtime(tempList[i-1]))
	}
	for i := 0; i < 8; i++ {
		if b & 0x01 == 0x01 {
			result ^= tempList[i]
		}
		b >>= 1
	}
	return result
}