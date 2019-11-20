package memory

type Bit bool

type BitSlice []Bit

func (p BitSlice) Len() int           { return len(p) }
func (p BitSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p BitSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Address uint64

type Memory interface {
	Read(address Address, length uint64) []Bit
	Write(address Address, value []Bit)
	Navigator() Navigator
}

type Navigator interface {
	Jump(address Address)
	Skip()
	Back()
	Next() Bit
	ReadNext(length Address) []Bit
	Last() Bit
	ReadLast(length Address) []Bit
}
