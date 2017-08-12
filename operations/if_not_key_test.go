package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestIfNotKeyParser_Matches(t *testing.T) {
	parser := IfNotKeyParser{}

	assert.True(t, parser.Matches(0xE9A1))
}

func TestIfNotKeyParser_DoesNotMatchFirst(t *testing.T) {
	parser := IfNotKeyParser{}

	assert.False(t, parser.Matches(0xF9A1))
}

func TestIfNotKeyParser_DoesNotMatchLast(t *testing.T) {
	parser := IfNotKeyParser{}

	assert.False(t, parser.Matches(0xE901))
}

func TestIfNotKeyParser_CreateOp(t *testing.T) {
	parser := IfNotKeyParser{}
	expected := IfNotKeyOp{register: 0xA}

	assert.Equal(t, expected, parser.CreateOp(0xEAA1))
}

func TestIfNotKeyOp_String(t *testing.T) {
	op := IfNotKeyOp{register: 0x3}

	assert.Equal(t, "If key != V3", op.String())
}

func TestIfNotKeyOp_ExecuteFalse(t *testing.T) {
	vm := system.VirtualMachine{}
	vm.Keyboard[0x3] = false
	vm.Registers[0x0] = 0x3
	vm.ProgramCounter = 0x07

	op := IfNotKeyOp{register: 0x0}

	op.Execute(&vm)

	assert.Equal(t, uint16(0x9), vm.ProgramCounter)
}

func TestIfNotKeyOp_ExecuteTrue(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Keyboard[0x3] = true
	vm.Registers[0x0] = 0x3
	vm.ProgramCounter = 0x07

	op := IfNotKeyOp{register: 0x0}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x07), vm.ProgramCounter)
}
