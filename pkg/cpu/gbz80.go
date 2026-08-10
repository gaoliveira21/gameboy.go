package cpu

import (
	"github.com/gaoliveira21/gameboy.go/pkg/memory"
)

type instruction struct {
	operation func() uint
	Mnemonic  string
	Size      uint16
}

type GBZ80 struct {
	// Registers
	a, b, c, d, e, h, l uint8
	flags               *Flags

	sp, pc uint16

	cycles uint

	mem            memory.MemReadWriter
	instructionSet [256]*instruction
}

func NewGBZ80(m memory.MemReadWriter) *GBZ80 {
	g := &GBZ80{
		a: 0x01,
		b: 0x00,
		c: 0x13,
		d: 0x00,
		e: 0xd8,
		h: 0x01,
		l: 0x4d,
		flags: &Flags{
			value: 0xb0,
		},
		pc:  0x100,
		sp:  0xFFFE,
		mem: m,
	}

	g.instructionSet = [256]*instruction{
		0x00: {g._NOP, "NOP", 1},
		0x01: {g._LD_BC_n16, "LD BC, n16", 3},
		0x02: {g._LD_BC_A, "LD [BC], A", 1},
		0x03: {g._TODO, "INC BC", 1},
		0x04: {g._TODO, "INC B", 1},
		0x05: {g._TODO, "DEC B", 1},
		0x06: {g._LD_B_n8, "LD B, n8", 2},
		0x07: {g._TODO, "RLCA", 1},
		0x08: {g._TODO, "LD [a16], SP", 3},
		0x09: {g._TODO, "ADD HL, BC", 1},
		0x0A: {g._TODO, "LD A, [BC]", 1},
		0x0B: {g._TODO, "DEC BC", 1},
		0x0C: {g._TODO, "INC C", 1},
		0x0D: {g._TODO, "DEC C", 1},
		0x0E: {g._TODO, "LD C, n8", 2},
		0x0F: {g._TODO, "RRCA", 1},
	}

	g.boot()

	return g
}

func (gbz *GBZ80) Interrupt() {
	return
}

func (gbz *GBZ80) Run() uint {
	opcode := gbz.fetch()
	cycles := gbz.exec(opcode)
	gbz.cycles += cycles

	return cycles
}

func (gbz *GBZ80) fetch() byte {
	opcode := gbz.mem.Read(gbz.pc)
	gbz.pc++
	return opcode
}

func (gbz *GBZ80) exec(opcode byte) uint {
	instruction := gbz.instructionSet[opcode]
	return instruction.operation()
}

func (gbz *GBZ80) boot() {
	gbz.mem.Write(0xFF00, 0xCF)
	gbz.mem.Write(0xFF01, 0x00)
	gbz.mem.Write(0xFF02, 0x7E)
	gbz.mem.Write(0xFF04, 0xAB)
	gbz.mem.Write(0xFF05, 0x00)
	gbz.mem.Write(0xFF06, 0x00)
	gbz.mem.Write(0xFF07, 0xF8)
	gbz.mem.Write(0xFF0F, 0xE1)
	gbz.mem.Write(0xFF10, 0x80)
	gbz.mem.Write(0xFF11, 0xBF)
	gbz.mem.Write(0xFF12, 0xF3)
	gbz.mem.Write(0xFF13, 0xFF)
	gbz.mem.Write(0xFF14, 0xBF)
	gbz.mem.Write(0xFF16, 0x3F)
	gbz.mem.Write(0xFF17, 0x00)
	gbz.mem.Write(0xFF18, 0xFF)
	gbz.mem.Write(0xFF19, 0xBF)
	gbz.mem.Write(0xFF1A, 0x7F)
	gbz.mem.Write(0xFF1B, 0xFF)
	gbz.mem.Write(0xFF1C, 0x9F)
	gbz.mem.Write(0xFF1D, 0xFF)
	gbz.mem.Write(0xFF1E, 0xBF)
	gbz.mem.Write(0xFF20, 0xFF)
	gbz.mem.Write(0xFF21, 0x00)
	gbz.mem.Write(0xFF22, 0x00)
	gbz.mem.Write(0xFF23, 0xBF)
	gbz.mem.Write(0xFF24, 0x77)
	gbz.mem.Write(0xFF25, 0xF3)
	gbz.mem.Write(0xFF26, 0xF1)
	gbz.mem.Write(0xFF40, 0x91)
	gbz.mem.Write(0xFF41, 0x85)
	gbz.mem.Write(0xFF42, 0x00)
	gbz.mem.Write(0xFF43, 0x00)
	gbz.mem.Write(0xFF44, 0x00)
	gbz.mem.Write(0xFF45, 0x00)
	gbz.mem.Write(0xFF46, 0xFF)
	gbz.mem.Write(0xFF47, 0xFC)
	gbz.mem.Write(0xFF48, 0xFF)
	gbz.mem.Write(0xFF49, 0xFF)
	gbz.mem.Write(0xFF4A, 0x00)
	gbz.mem.Write(0xFF4B, 0x00)
	gbz.mem.Write(0xFFFF, 0x00)
}

func (gbz *GBZ80) _TODO() uint {
	panic("Instruction not implemented")
}

func (gbz *GBZ80) _NOP() uint {
	return 4
}

func (gbz *GBZ80) _LD_BC_n16() uint {
	lb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	hb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	gbz.b, gbz.c = hb, lb

	return 12
}

func (gbz *GBZ80) _LD_BC_A() uint {
	hb := uint16(gbz.b)
	lb := uint16(gbz.c)

	addr := (hb << 8) | lb
	gbz.mem.Write(addr, gbz.a)

	return 8
}

func (gbz *GBZ80) _LD_B_n8() uint {
	n := gbz.mem.Read(gbz.pc)
	gbz.pc++

	gbz.b = n

	return 8
}
