package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestIfConstantParser_Matches(t *testing.T) {
	parser := IfConstantParser{}

	assert.True(t, parser.Matches(0x37B6))
}

func TestIfConstantParser_DoesNotMatch(t *testing.T) {
	parser := IfConstantParser{}

	assert.False(t, parser.Matches(0x57B6))
}

func TestIfConstantParser_CreateOp(t *testing.T) {
	parser := IfConstantParser{}
	expected := IfConstantOp{register: 0x7, value: 0xB6}

	assert.Equal(t, expected, parser.CreateOp(0x37B6))
}

func TestIfConstantOp_String(t *testing.T) {
	op := IfConstantOp{register: 0x5, value: 0xAC}

	assert.Equal(t, "If V5 == AC", op.String())
}

func TestIfConstantOp_ExecuteMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x8
	vm.Registers[0x4] = 0xA1

	op := IfConstantOp{register: 0x4, value: 0xA1}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0xA), vm.ProgramCounter)
}

func TestIfConstantOp_ExecuteNoMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x8
	vm.Registers[0x4] = 0xAC

	op := IfConstantOp{register: 0x4, value: 0xA1}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x8), vm.ProgramCounter)
}
