package memory

type BoolSlice []bool

func (p BoolSlice) Len() int           { return len(p) }
func (p BoolSlice) Less(i, j int) bool { return !p[i] && p[j] }
func (p BoolSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Memory interface {
	Read(address uint, length uint) []bool
	Write(address uint, value []bool)
	Navigator() Navigator
}

type Navigator interface {
	Jump(address uint)
	Skip()
	Back()
	Next() bool
	ReadNext(length uint) []bool
	Last() bool
	ReadLast(length uint) []bool
}
