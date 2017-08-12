package operations

import (
	"fmt"
	"chip8/system"
)

type SetIndexParser struct {}
func(p SetIndexParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xA
}

func(p SetIndexParser) CreateOp(opcode system.OpCode) Operation {
	return SetIndexOp{
		value: uint16(opcode) & 0x0FFF,
	}
}

type SetIndexOp struct {
	value uint16
}
func(o SetIndexOp) String() string {
	return fmt.Sprintf("I = %X", o.value)
}

func(o SetIndexOp) Execute(machine *system.VirtualMachine) {
	machine.IndexRegister = o.value
}
