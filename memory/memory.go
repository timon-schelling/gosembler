package memory

type Memory interface {
	Read(address uint, length uint) []bool
	Write(address uint, value []bool)
	Navigator(address *uint) Navigator
}

type Navigator interface {
	Address() uint
	Jump(address uint)
	Skip()
	Back()
	Next() bool
	Read(length uint) []bool
}
