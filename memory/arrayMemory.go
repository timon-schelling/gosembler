package memory

type SimpleArrayMemory struct {
	Backend []bool
}

func NewSimpleArrayMemory() SimpleArrayMemory {
	return SimpleArrayMemory{Backend: []bool{}}
}

func (m SimpleArrayMemory) Read(address Address, length uint64) (r []Bit) {
	for i := uint64(0); i < length; i++ {
		r = append(r, Bit(m.Backend[i+uint64(address)]))
	}
	return r
}

func (SimpleArrayMemory) Write(address Address, value []Bit) {
	panic("implement me")
}

func (SimpleArrayMemory) Navigator() Navigator {
	panic("implement me")
}
