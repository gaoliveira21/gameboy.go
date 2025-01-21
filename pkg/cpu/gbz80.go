package cpu

type GBZ80 struct {
	// Registers
	a, b, c, d, e, h, l uint8

	sp, pc uint16

	mem *memory
}
