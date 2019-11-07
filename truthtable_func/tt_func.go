package truthtable_func

import "finitefield/finite_field_cp"

func split(a0, a1, a2, a3 uint8) (uint8, uint8, uint8, uint8) {
	a := a0 << 3 | a1 << 2 | a2 << 1 | a3
	t := finite_field_cp.InverseGF24(a)
	return t & 0x08 >> 3, t & 0x04 >> 2, t & 0x02 >> 1, t & 0x01
}

func GetCofficient() ([]uint8, []uint8, []uint8, []uint8) {

}