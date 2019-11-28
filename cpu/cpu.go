package cpu

import "github.com/timon-schelling/gosembler/memory"

type Cpu struct {
	memory memory.Memory
	exp    uint
	sp     uint
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

func (c Cpu) Load(source []bool) {
	c.memory.Write(0, source)
}

func (c Cpu) running() {
	memoryNavigator := c.memory.Navigator()
	for c.run {
		optCode := memoryNavigator.ReadNext(8)
	}
}

func (c Cpu) Start() {

}

func (c Cpu) Stop() {

}
