package memory

type SimpleArrayMemory struct {
	backend []bool
}

func NewSimpleArrayMemory() SimpleArrayMemory {
	return SimpleArrayMemory{backend: []bool{}}
}

func (m SimpleArrayMemory) Read(address uint, length uint) (r []bool) {
	for i := uint(0); i < length; i++ {
		r = append(r, m.backend[i+address])
	}
	return r
}

func (m SimpleArrayMemory) Write(address uint, values []bool) {
	valueLength := uint(len(values))
	for uint(len(m.backend)) < valueLength {
		m.backend = append(m.backend, false)
	}
	for i := uint(0); i < valueLength; i++ {
		currentBit := uint(i) + address
		if currentBit == uint(len(m.backend)) {
			m.backend = append(m.backend, bool(values[i]))
		} else {
			m.backend[currentBit] = bool(values[i])
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

func (mn SimpleArrayMemoryNavigator) Address() uint {
	return mn.current
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

func (mn SimpleArrayMemoryNavigator) Next() (r bool) {
	r = mn.memory.Read(mn.current, 1)[0]
	mn.current++
	return r
}

func (mn SimpleArrayMemoryNavigator) Read(length uint) (r []bool) {
	r = mn.memory.Read(mn.current, length)
	mn.current += length
	return r
}
