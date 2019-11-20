package cpu

import "github.com/timon-schelling/gosembler/memory"

type Cpu struct {
	memory memory.Memory
	exp    memory.Address
	sp     memory.Address
	run    bool
	err    bool
	carry  bool
	zero   bool
}

func (c Cpu) NewCpu(memory memory.Memory) Cpu {
	return Cpu{
		memory: memory,
		run:    true,
	}
}

func (c Cpu) Load(source []memory.Bit) {
	c.memory.Write(0, source)
}

func (c Cpu) running() {
	memoryNavigator := c.memory.Navigator()
	for c.run {

	}
}

func (c Cpu) Start() {

}

func (c Cpu) Stop() {

}
