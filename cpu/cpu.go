package cpu

import (
	"github.com/timon-schelling/gosembler/converter"
	"github.com/timon-schelling/gosembler/memory"
)

type Cpu struct {
	memory memory.Memory
	exp    uint
	sp     uint
	jmp    bool
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

func (c Cpu) Start() {
	go c.running()
}

func (c Cpu) Stop() {

}

func (c Cpu) Reset() {

}

func (c Cpu) running() {
	memoryNavigator := c.memory.Navigator()
	c.run = true
	for c.run {
		c.exp = memoryNavigator.Address()
		optCode := converter.BoolsToBytes(memoryNavigator.Read(8))[0]
		c.doStep(optCode)
	}
}

func (c Cpu) doStep(optCode byte) {
	switch optCode {
	case 0:
		c.execHlt()

	case 1:
		c.execMovRegToReg()
	case 2:
		c.execMovAddressToReg()
	case 3:
		c.execMovRegAddressToReg()
	case 4:
		c.execMovRegToAddress()
	case 5:
		c.execMovRegToRegAddress()
	case 6:
		c.execMovNumberToReg()
	case 7:
		c.execMovNumberToAddress()
	case 8:
		c.execMovNumberToRegAddress()

	case 10:
		c.execAddRegToReg()
	case 11:
		c.execAddRegAddressToReg()
	case 12:
		c.execAddAddressToReg()
	case 13:
		c.execAddNumberToReg()

	case 14:
		c.execSubRegFromReg()
	case 15:
		c.execSubRegAddressFromReg()
	case 16:
	case 17:

	default:
		c.execHlt()
	}
}

//HLT
func (c Cpu) execHlt() {

}

//MOVE
func (c Cpu) execMovRegToReg() {

}

func (c Cpu) execMovAddressToReg() {

}

func (c Cpu) execMovRegAddressToReg() {

}

func (c Cpu) execMovRegToAddress() {

}

func (c Cpu) execMovRegToRegAddress() {

}

func (c Cpu) execMovNumberToReg() {

}

func (c Cpu) execMovNumberToAddress() {

}

func (c Cpu) execMovNumberToRegAddress() {

}

//ADD
func (c Cpu) execAddRegToReg() {

}

func (c Cpu) execAddRegAddressToReg() {

}

func (c Cpu) execAddAddressToReg() {

}

func (c Cpu) execAddNumberToReg() {

}

//SUB
func (c Cpu) execSubRegFromReg() {

}

func (c Cpu) execSubRegAddressFromReg() {

}

func (c Cpu) execSubAddressFromReg() {

}

func (c Cpu) execSubNumberFromReg() {

}

//INC
func (c Cpu) execIncReg() {

}

//DEC
func (c Cpu) execDecReg() {

}

//CMP
func (c Cpu) execCmpRegWithReg() {

}

func (c Cpu) execCmpRegAddressWithReg() {

}

func (c Cpu) execCmpAddressWithReg() {

}

func (c Cpu) execCmpNumberWithReg() {

}

//JMP
func (c Cpu) execJmpToRegAddress() {

}

func (c Cpu) execJmpToAddress() {

}

//JC
func (c Cpu) execJcToRegAddress() {

}

func (c Cpu) execJcToAddress() {

}

//JNC
func (c Cpu) execJncToRegAddress() {

}

func (c Cpu) execJncToAddress() {

}

//JZ
func (c Cpu) execJzToRegAddress() {

}

func (c Cpu) execJzToAddress() {

}

//JNZ
func (c Cpu) execJnzToRegAddress() {

}

func (c Cpu) execJnzToAddress() {

}

//JA
func (c Cpu) execJaToRegAddress() {

}

func (c Cpu) execJaToAddress() {

}

//JNA
func (c Cpu) execJnaToRegAddress() {

}

func (c Cpu) execJnaToAddress() {

}

//PUSH
func (c Cpu) execPushReg() {

}

func (c Cpu) execPushRegAddress() {

}

func (c Cpu) execPushAddress() {

}

func (c Cpu) execPushNumber() {

}

//POP
func (c Cpu) execPopReg() {

}

//CALL
func (c Cpu) execCallRegAddress() {

}

func (c Cpu) execCallAddress() {

}

//RET
func (c Cpu) execRet() {

}

//MUL
func (c Cpu) execMulReg() {

}

func (c Cpu) execMulRegAddress() {

}

func (c Cpu) execMulAddress() {

}

func (c Cpu) execMulNumber() {

}

//DIV
func (c Cpu) execDivReg() {

}

func (c Cpu) execDivRegAddress() {

}

func (c Cpu) execDivAddress() {

}

func (c Cpu) execDivNumber() {

}

//AND
func (c Cpu) execAndRegWithReg() {

}

func (c Cpu) execAndRegAddressWithReg() {

}

func (c Cpu) execAndAddressWithReg() {

}

func (c Cpu) execAndNumberWithReg() {

}

//OR
func (c Cpu) execOrRegWithReg() {

}

func (c Cpu) execOrRegAddressWithReg() {

}

func (c Cpu) execOrAddressWithReg() {

}

func (c Cpu) execOrNumberWithReg() {

}

//XOR
func (c Cpu) execXorRegWithReg() {

}

func (c Cpu) execXorRegAddressWithReg() {

}

func (c Cpu) execXorAddressWithReg() {

}

func (c Cpu) execXorNumberWithReg() {

}

//NOT
func (c Cpu) execNotReg() {

}

//SHL
func (c Cpu) execShlRegWithReg() {

}

func (c Cpu) execShlRegAddressWithReg() {

}

func (c Cpu) execShlAddressWithReg() {

}

func (c Cpu) execShlNumberWithReg() {

}

//SHR
func (c Cpu) execShrRegWithReg() {

}

func (c Cpu) execShrRegAddressWithReg() {

}

func (c Cpu) execShrAddressWithReg() {

}

func (c Cpu) execShrNumberWithReg() {

}
