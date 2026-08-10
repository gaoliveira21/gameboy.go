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
		0x08: {g._LD_n16_SP, "LD [n16], SP", 3},
		0x09: {g._TODO, "ADD HL, BC", 1},
		0x0A: {g._LD_A_BC, "LD A, [BC]", 1},
		0x0B: {g._TODO, "DEC BC", 1},
		0x0C: {g._TODO, "INC C", 1},
		0x0D: {g._TODO, "DEC C", 1},
		0x0E: {g._LD_C_n8, "LD C, n8", 2},
		0x0F: {g._TODO, "RRCA", 1},

		0x10: {g._TODO, "STOP n8", 2},
		0x11: {g._LD_DE_n16, "LD DE, n16", 3},
		0x12: {g._LD_DE_A, "LD [DE], A", 1},
		0x13: {g._TODO, "INC DE", 1},
		0x14: {g._TODO, "INC D", 1},
		0x15: {g._TODO, "DEC D", 1},
		0x16: {g._LD_D_n8, "LD D, n8", 2},
		0x17: {g._TODO, "RLA", 1},
		0x18: {g._TODO, "JR e8", 2},
		0x19: {g._TODO, "ADD HL, DE", 1},
		0x1A: {g._LD_A_DE, "LD A, [DE]", 1},
		0x1B: {g._TODO, "DEC DE", 1},
		0x1C: {g._TODO, "INC E", 1},
		0x1D: {g._TODO, "DEC E", 1},
		0x1E: {g._LD_E_n8, "LD E, n8", 2},
		0x1F: {g._TODO, "RRA", 1},

		0x20: {g._TODO, "JR NZ, e8", 2},
		0x21: {g._LD_HL_n16, "LD HL, n16 ", 3},
		0x22: {g._TODO, "LD [HL+], A", 1},
		0x23: {g._TODO, "INC HL", 1},
		0x24: {g._TODO, "INC H", 1},
		0x25: {g._TODO, "DEC H", 1},
		0x26: {g._TODO, "LD H, n8", 2},
		0x27: {g._TODO, "DAA", 1},
		0x28: {g._TODO, "JR Z, e8", 2},
		0x29: {g._TODO, "ADD HL, HL", 1},
		0x2A: {g._TODO, "LD A, [HL+]", 1},
		0x2B: {g._TODO, "DEC HL", 1},
		0x2C: {g._TODO, "INC L", 1},
		0x2D: {g._TODO, "DEC L", 1},
		0x2E: {g._TODO, "LD L, n8", 2},
		0x2F: {g._TODO, "CPL", 1},
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
	gbz.ldrpn16(&gbz.b, &gbz.c)
	return 12
}

func (gbz *GBZ80) _LD_BC_A() uint {
	gbz.ldrpr8(&gbz.b, &gbz.c, &gbz.a)
	return 8
}

func (gbz *GBZ80) _LD_B_n8() uint {
	gbz.ldr8n8(&gbz.b)
	return 8
}

func (gbz *GBZ80) _LD_n16_SP() uint {
	hb := byte(gbz.sp >> 8)
	lb := byte(gbz.sp & 0xFF)

	addr := uint16(gbz.mem.Read(gbz.pc)) | (uint16(gbz.mem.Read(gbz.pc+1)) << 8)
	gbz.pc += 2

	gbz.mem.Write(addr, lb)
	gbz.mem.Write(addr+1, hb)

	return 20
}

func (gbz *GBZ80) _LD_A_BC() uint {
	gbz.ldr8rp(&gbz.a, &gbz.b, &gbz.c)
	return 8
}

func (gbz *GBZ80) _LD_C_n8() uint {
	gbz.ldr8n8(&gbz.c)
	return 8
}

func (gbz *GBZ80) _LD_DE_n16() uint {
	gbz.ldrpn16(&gbz.d, &gbz.e)
	return 12
}

func (gbz *GBZ80) _LD_DE_A() uint {
	gbz.ldrpr8(&gbz.d, &gbz.e, &gbz.a)
	return 8
}

func (gbz *GBZ80) _LD_D_n8() uint {
	gbz.ldr8n8(&gbz.d)
	return 8
}

func (gbz *GBZ80) _LD_A_DE() uint {
	gbz.ldr8rp(&gbz.a, &gbz.d, &gbz.e)
	return 8
}

func (gbz *GBZ80) _LD_E_n8() uint {
	gbz.ldr8n8(&gbz.e)
	return 8
}

func (gbz *GBZ80) _LD_HL_n16() uint {
	gbz.ldrpn16(&gbz.h, &gbz.l)
	return 12
}
