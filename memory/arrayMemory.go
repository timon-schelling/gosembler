package memory

import "sort"

type SimpleArrayMemory struct {
	backend []bool
}

func NewSimpleArrayMemory() SimpleArrayMemory {
	return SimpleArrayMemory{backend: []bool{}}
}

func (m SimpleArrayMemory) Read(address Address, length uint64) (r []Bit) {
	for i := uint64(0); i < length; i++ {
		r = append(r, Bit(m.backend[i+uint64(address)]))
	}
	return r
}

func (m SimpleArrayMemory) Write(address Address, value []Bit) {
	valueLength := len(value)
	for len(m.backend) < valueLength {
		m.backend = append(m.backend, false)
	}
	for i := 0; i < valueLength; i++ {
		currentAddress := Address(i) + address
		if currentAddress == Address(len(m.backend)) {
			m.backend = append(m.backend, bool(value[i]))
		} else {
			m.backend[currentAddress] = bool(value[i])
		}
	}
}

func (m SimpleArrayMemory) Navigator() Navigator {
	return SimpleArrayMemoryNavigator{
		memory:  m,
		current: 0,
	}
}

type SimpleArrayMemoryNavigator struct {
	memory  SimpleArrayMemory
	current Address
}

func (mn SimpleArrayMemoryNavigator) Jump(address Address) {
	mn.current = address
}

func (mn SimpleArrayMemoryNavigator) Skip() {
	mn.current++
}

func (mn SimpleArrayMemoryNavigator) Back() {
	mn.current--
}

func (mn SimpleArrayMemoryNavigator) Next() (r Bit) {
	r = mn.memory.Read(mn.current, 1)[0]
	mn.current++
	return r
}

func (mn SimpleArrayMemoryNavigator) ReadNext(length uint64) (r []Bit) {
	r = mn.memory.Read(mn.current, length)
	mn.current += Address(length)
	return r
}

func (mn SimpleArrayMemoryNavigator) Last() Bit {
	mn.current--
	return mn.memory.Read(mn.current, 1)[0]
}

func (mn SimpleArrayMemoryNavigator) ReadLast(length uint64) (r []Bit) {
	r = sort.Reverse(mn.memory.Read(mn.current-Address(length), length))
	mn.current -= Address(length)
	return r
}
