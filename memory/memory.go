package memory

type Bit bool

type Address uint64

type Memory interface {
	Read(address Address, length uint64) []Bit
	Write(address Address, value []Bit)
	Navigator() Navigator
}

type Navigator interface {
	Jump(address Address)
	Skip()
	Next() Bit
	Last() Bit
}
