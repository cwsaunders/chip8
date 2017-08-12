package operations

import (
	"fmt"
	"chip8/system"
)

type IfRegisterParser struct {}
func(p IfRegisterParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x5 && opcode & 0x000F == 0x0
}

func(p IfRegisterParser) CreateOp(opcode system.OpCode) Operation {
	return IfRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type IfRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o IfRegisterOp) String() string {
	return fmt.Sprintf("If V%X == V%X", o.register1, o.register2)
}

func(o IfRegisterOp) Execute(machine *system.VirtualMachine) {
	if (machine.Registers[o.register1] == machine.Registers[o.register2]) {
		// TODO:  Refactor this
		machine.ProgramCounter += 2
	}
}
