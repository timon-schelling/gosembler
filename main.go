package main

import (
	"github.com/timon-schelling/gosembler/memory"
)

func toBoolArray(s string) (r []bool) {
	for _, v := range s {
		if v == '0' {
			r = append(r, false)
		} else if v == '1' {
			r = append(r, true)
		}
	}
	return r
}

func main() {
	mv := toBoolArray("1011000010111010001011010000010111101011101101011111001000101000111110100010001001010010101010011010101010100100100110011010010011011010101010101101000101010100010100100101011011010000000000011111101001010000111110110100100001011111000101000000010101010100110")
	v := memory.NewSimpleArrayMemory()
	v.Write(0, mv)
	r := v.Read(1, 40)
	println(r)
}
