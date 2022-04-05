package vm

type Instruction uint32

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

func (instruction Instruction) Opcode() int {
	return int(instruction & 0x3F)
}

func (instruction Instruction) ABC() (a, b, c int) {
	a = int(instruction >> 6 & 0xFF)
	c = int(instruction >> 14 & 0xFF)
	b = int(instruction >> 23 & 0xFF)
	return
}

func (instruction Instruction) ABx() (a, bx int) {
	a = int(instruction >> 6 & 0xFF)
	bx = int(instruction >> 14)
	return
}

func (instruction Instruction) AsBx() (a, sbx int) {
	a, bx := instruction.ABx()
	return a, bx - MAXARG_sBx
}

func (instruction Instruction) Ax() int {
	return int(instruction >> 6)
}

func (instruction Instruction) OpName() string {
	return opcodes[instruction.Opcode()].name
}

func (instruction Instruction) OpMode() byte {
	return opcodes[instruction.Opcode()].opMode
}

func (instruction Instruction) BMode() byte {
	return opcodes[instruction.OpMode()].argBMode
}

func (instruction Instruction) CMode() byte {
	return opcodes[instruction.OpMode()].argCMode
}
