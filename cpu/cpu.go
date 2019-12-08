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
		c.doStep(InstructionCode(optCode))
	}
}

type InstructionCode int

const (
	HLT InstructionCode = iota

	MOV_REG_TO_REG
	MOV_ADDRESS_TO_REG
	MOV_REGADDRESS_TO_REG
	MOV_REG_TO_ADDRESS
	MOV_REG_TO_REGADDRESS
	MOV_NUMBER_TO_REG
	MOV_NUMBER_TO_ADDRESS
	MOV_NUMBER_TO_REGADDRESS

	ADD_REG_TO_REG
	ADD_REGADDRESS_TO_REG
	ADD_ADDRESS_TO_REG
	ADD_NUMBER_TO_REG

	SUB_REG_FROM_REG
	SUB_REGADDRESS_FROM_REG
	SUB_ADDRESS_FROM_REG
	SUB_NUMBER_FROM_REG

	INC_REG

	DEC_REG

	CMP_REG_WITH_REG
	CMP_REGADDRESS_WITH_REG
	CMP_ADDRESS_WITH_REG
	CMP_NUMBER_WITH_REG

	JMP_REGADDRESS
	JMP_ADDRESS

	JC_REGADDRESS
	JC_ADDRESS

	JNC_REGADDRESS
	JNC_ADDRESS

	JZ_REGADDRESS
	JZ_ADDRESS

	JNZ_REGADDRESS
	JNZ_ADDRESS

	JA_REGADDRESS
	JA_ADDRESS

	JNA_REGADDRESS
	JNA_ADDRESS

	PUSH_REG
	PUSH_REGADDRESS
	PUSH_ADDRESS
	PUSH_NUMBER

	POP_REG

	CALL_REGADDRESS
	CALL_ADDRESS

	RET

	MUL_REG
	MUL_REGADDRESS
	MUL_ADDRESS
	MUL_NUMBER

	DIV_REG
	DIV_REGADDRESS
	DIV_ADDRESS
	DIV_NUMBER

	AND_REG_WITH_REG
	AND_REGADDRESS_WITH_REG
	AND_ADDRESS_WITH_REG
	AND_NUMBER_WITH_REG

	OR_REG_WITH_REG
	OR_REGADDRESS_WITH_REG
	OR_ADDRESS_WITH_REG
	OR_NUMBER_WITH_REG

	XOR_REG_WITH_REG
	XOR_REGADDRESS_WITH_REG
	XOR_ADDRESS_WITH_REG
	XOR_NUMBER_WITH_REG

	NOT_REG

	SHL_REG_WITH_REG
	SHL_REGADDRESS_WITH_REG
	SHL_ADDRESS_WITH_REG
	SHL_NUMBER_WITH_REG

	SHR_REG_WITH_REG
	SHR_REGADDRESS_WITH_REG
	SHR_ADDRESS_WITH_REG
	SHR_NUMBER_WITH_REG
)

func (c Cpu) doStep(optCode InstructionCode) {
	switch optCode {
	case HLT:
		c.execHlt()

	case MOV_REG_TO_REG:
		c.execMovRegToReg()
	case MOV_ADDRESS_TO_REG:
		c.execMovAddressToReg()
	case MOV_REGADDRESS_TO_REG:
		c.execMovRegAddressToReg()
	case MOV_REG_TO_ADDRESS:
		c.execMovRegToAddress()
	case MOV_REG_TO_REGADDRESS:
		c.execMovRegToRegAddress()
	case MOV_NUMBER_TO_REG:
		c.execMovNumberToReg()
	case MOV_NUMBER_TO_ADDRESS:
		c.execMovNumberToAddress()
	case MOV_NUMBER_TO_REGADDRESS:
		c.execMovNumberToRegAddress()

	case ADD_REG_TO_REG:
		c.execAddRegToReg()
	case ADD_REGADDRESS_TO_REG:
		c.execAddRegAddressToReg()
	case ADD_ADDRESS_TO_REG:
		c.execAddAddressToReg()
	case ADD_NUMBER_TO_REG:
		c.execAddNumberToReg()

	case SUB_REG_FROM_REG:
		c.execSubRegFromReg()
	case SUB_REGADDRESS_FROM_REG:
		c.execSubRegAddressFromReg()
	case SUB_ADDRESS_FROM_REG:
		c.execSubAddressFromReg()
	case SUB_NUMBER_FROM_REG:
		c.execSubNumberFromReg()

	case INC_REG:
		c.execIncReg()

	case DEC_REG:
		c.execDecReg()

	case CMP_REG_WITH_REG:
		c.execCmpRegWithReg()
	case CMP_REGADDRESS_WITH_REG:
		c.execCmpRegAddressWithReg()
	case CMP_ADDRESS_WITH_REG:
		c.execCmpAddressWithReg()
	case CMP_NUMBER_WITH_REG:
		c.execCmpNumberWithReg()

	case JMP_REGADDRESS:
		c.execJmpToRegAddress()
	case JMP_ADDRESS:
		c.execJmpToAddress()

	case JC_REGADDRESS:
		c.execJcToRegAddress()
	case JC_ADDRESS:
		c.execJcToAddress()

	case JNC_REGADDRESS:
		c.execJncToRegAddress()
	case JNC_ADDRESS:
		c.execJncToAddress()

	case JZ_REGADDRESS:
		c.execJzToRegAddress()
	case JZ_ADDRESS:
		c.execJzToAddress()

	case JNZ_REGADDRESS:
		c.execJnzToRegAddress()
	case JNZ_ADDRESS:
		c.execJnzToAddress()

	case JA_REGADDRESS:
		c.execJaToRegAddress()
	case JA_ADDRESS:
		c.execJaToAddress()

	case JNA_REGADDRESS:
		c.execJnaToRegAddress()
	case JNA_ADDRESS:
		c.execJnaToAddress()

	case PUSH_REG:
		c.execPushReg()
	case PUSH_REGADDRESS:
		c.execPushRegAddress()
	case PUSH_ADDRESS:
		c.execPushAddress()
	case PUSH_NUMBER:
		c.execPushNumber()

	case POP_REG:
		c.execPopReg()

	case CALL_REGADDRESS:
		c.execCallRegAddress()
	case CALL_ADDRESS:
		c.execCallAddress()

	case RET:
		c.execRet()

	case MUL_REG:
		c.execMulReg()
	case MUL_REGADDRESS:
		c.execMulRegAddress()
	case MUL_ADDRESS:
		c.execMulAddress()
	case MUL_NUMBER:
		c.execMulNumber()

	case DIV_REG:
		c.execDivReg()
	case DIV_REGADDRESS:
		c.execDivRegAddress()
	case DIV_ADDRESS:
		c.execDivAddress()
	case DIV_NUMBER:
		c.execDivNumber()

	case AND_REG_WITH_REG:
		c.execAndRegWithReg()
	case AND_REGADDRESS_WITH_REG:
		c.execAndRegAddressWithReg()
	case AND_ADDRESS_WITH_REG:
		c.execAndAddressWithReg()
	case AND_NUMBER_WITH_REG:
		c.execAndNumberWithReg()

	case OR_REG_WITH_REG:
		c.execOrRegWithReg()
	case OR_REGADDRESS_WITH_REG:
		c.execOrRegAddressWithReg()
	case OR_ADDRESS_WITH_REG:
		c.execAndAddressWithReg()
	case OR_NUMBER_WITH_REG:
		c.execAndNumberWithReg()

	case XOR_REG_WITH_REG:
		c.execXorRegWithReg()
	case XOR_REGADDRESS_WITH_REG:
		c.execXorRegAddressWithReg()
	case XOR_ADDRESS_WITH_REG:
		c.execXorAddressWithReg()
	case XOR_NUMBER_WITH_REG:
		c.execXorNumberWithReg()

	case NOT_REG:
		c.execNotReg()

	case SHL_REG_WITH_REG:
		c.execShlRegWithReg()
	case SHL_REGADDRESS_WITH_REG:
		c.execShlRegAddressWithReg()
	case SHL_ADDRESS_WITH_REG:
		c.execShlAddressWithReg()
	case SHL_NUMBER_WITH_REG:
		c.execShlNumberWithReg()

	case SHR_REG_WITH_REG:
		c.execShrRegWithReg()
	case SHR_REGADDRESS_WITH_REG:
		c.execShrRegAddressWithReg()
	case SHR_ADDRESS_WITH_REG:
		c.execShrAddressWithReg()
	case SHR_NUMBER_WITH_REG:
		c.execShrNumberWithReg()

	default:
		c.error()
	}
}

func (c Cpu) error() {

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
