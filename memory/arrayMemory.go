package memory

import "sort"

type SimpleArrayMemory struct {
	backend []bool
}

func NewSimpleArrayMemory() SimpleArrayMemory {
	return SimpleArrayMemory{backend: []bool{}}
}

func (m SimpleArrayMemory) Read(address uint, length uint64) (r []Bit) {
	for i := uint64(0); i < length; i++ {
		r = append(r, Bit(m.backend[i+uint64(address)]))
	}
	return r
}

func (m SimpleArrayMemory) Write(address uint, value []Bit) {
	valueLength := len(value)
	for len(m.backend) < valueLength {
		m.backend = append(m.backend, false)
	}
	for i := 0; i < valueLength; i++ {
		currentuint := uint(i) + address
		if currentuint == uint(len(m.backend)) {
			m.backend = append(m.backend, bool(value[i]))
		} else {
			m.backend[currentuint] = bool(value[i])
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
	current uint
}

func (mn SimpleArrayMemoryNavigator) Jump(address uint) {
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
	mn.current += uint(length)
	return r
}

func (mn SimpleArrayMemoryNavigator) Last() Bit {
	mn.current--
	return mn.memory.Read(mn.current, 1)[0]
}

func (mn SimpleArrayMemoryNavigator) ReadLast(length uint64) (r []Bit) {
	//rs := BitSlice{bits: mn.memory.Read(mn.current-uint(length), length)}
	//r = sort.Reverse(rs).(BitSlice).bits
	rs := []bool{true, false}
	t := sort.Reverse(BoolSlice(rs))
	mn.current -= uint(length)
	return r
}
