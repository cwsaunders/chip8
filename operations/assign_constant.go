package operations

import (
	"fmt"
	"chip8/system"
)

type assignConstantParser struct {}
func(p assignConstantParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x6
}

func(p assignConstantParser) createOp(opcode system.OpCode) Operation {
	return AssignConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode & 0x00FF),
	}
}

type AssignConstantOp struct {
	register byte
	value byte
}
func(o AssignConstantOp) String() string {
	return fmt.Sprintf("V%X = %X", o.register, o.value)
}

func(o AssignConstantOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] = o.value
}
