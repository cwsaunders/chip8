package system


type VirtualMachine struct {
	Memory [4096]byte
	Registers [16]byte
	Stack []uint16

	ProgramCounter uint16
	IndexRegister uint16

	DelayTimer byte
	SoundTimer byte

	// Represents the state of key presses
	Keyboard [16]bool
	Pixels [32]int64
}

func NewVirtualMachine() VirtualMachine {
	// TODO:  Explain memory layout here...
	vm := VirtualMachine{}

	// Load the memory from the font set into the
	// lower "program" memory space
	for i := 0; i < len(fontSet); i++ {
		vm.Memory[i] = fontSet[i]
	}

	return vm
}

func (vm *VirtualMachine) Load(data []byte) {

	// Load the memory starting in application space
	for i := 0; i < len(data); i++ {
		vm.Memory[512 + i] = data[i]
	}

	vm.ProgramCounter = 512
}

func (vm *VirtualMachine) Run() {
	// TODO:  Very naive implementation to get going

}

// CHIP-8 Font Set.
var fontSet = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}