package operations

import (
	"fmt"
	"chip8/system"
)

type ifNotRegisterParser struct {}
func(p ifNotRegisterParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x9 && opcode & 0x000F == 0x0
}

func(p ifNotRegisterParser) createOp(opcode system.OpCode) Operation {
	return IfNotRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type IfNotRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o IfNotRegisterOp) String() string {
	return fmt.Sprintf("If V%X != V%X", o.register1, o.register2)
}

func(o IfNotRegisterOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register1] != vm.Registers[o.register2] {
		vm.IncrementPC()
	}
}
