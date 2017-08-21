package operations

import (
	"fmt"
	"chip8/system"
)

type assignRegisterParser struct {}
func(p assignRegisterParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x0
}

func(p assignRegisterParser) createOp(opcode system.OpCode) Operation {
	return AssignRegisterOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

type AssignRegisterOp struct {
	register1 byte
	register2 byte
}
func(o AssignRegisterOp) String() string {
	return fmt.Sprintf("V%X = V%X", o.register1, o.register2)
}

func(o AssignRegisterOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] = vm.Registers[o.register2]
}
