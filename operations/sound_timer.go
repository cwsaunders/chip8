package operations

import (
	"fmt"
	"chip8/system"
)

type soundTimerParser struct {}
func(p soundTimerParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x18
}

func(p soundTimerParser) createOp(opcode system.OpCode) Operation {
	return SoundTimerOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type SoundTimerOp struct {
	register byte
}
func(o SoundTimerOp) String() string {
	return fmt.Sprintf("sound_timer = V%X", o.register)
}

func(o SoundTimerOp) Execute(vm *system.VirtualMachine) {
	vm.SoundTimer = vm.Registers[o.register]
}
