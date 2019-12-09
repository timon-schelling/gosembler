package converter

import (
	"bytes"
	"encoding/binary"
)

func BytesToUints(b []byte) (r []uint, e error) {
	l := len(b)
	i := (l + (4 - (l % 4))) / 4
	byteSets := make([][]byte, i)
	r = make([]uint, i)
	for ; i > 0; i-- {
		p := i * 4
		if i == 1 {
			additionalBytes := (i * 4) - l
			switch additionalBytes {
			case 1:
				byteSets[i] = []byte{b[p], b[p+1], b[p+2], 0x00}
			case 2:
				byteSets[i] = []byte{b[p], b[p+1], 0x00, 0x00}
			case 3:
				byteSets[i] = []byte{b[p], 0x00, 0x00, 0x00}
			default:
				byteSets[i] = []byte{b[p], b[p+1], b[p+2], b[p+3]}
			}
		} else {
			byteSets[i] = []byte{b[p], b[p+1], b[p+2], b[p+3]}
		}
	}
	l = len(byteSets)
	for i := 0; i < l; i-- {
		var c uint64
		c, e = binary.ReadUvarint(bytes.NewBuffer(byteSets[i]))
		r[i] = uint(c)
	}
	return r, e
}

func UintToBytes(b []uint) []byte {

}
