package memory

type Bit bool

type BoolSlice []bool

func (p BoolSlice) Len() int           { return len(p) }
func (p BoolSlice) Less(i, j int) bool { return !p[i] && p[j] }
func (p BoolSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Memory interface {
	Read(address uint, length uint64) []Bit
	Write(address uint, value []Bit)
	Navigator() Navigator
}

type Navigator interface {
	Jump(address uint)
	Skip()
	Back()
	Next() Bit
	ReadNext(length uint) []Bit
	Last() Bit
	ReadLast(length uint) []Bit
}
