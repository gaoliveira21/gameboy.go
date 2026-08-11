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

	interruptEnabled bool

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
		0x08: {g._LD_a16_SP, "LD [a16], SP", 3},
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
		0x22: {g._LD_HLI_A, "LD [HL+], A", 1},
		0x23: {g._TODO, "INC HL", 1},
		0x24: {g._TODO, "INC H", 1},
		0x25: {g._TODO, "DEC H", 1},
		0x26: {g._LD_H_n8, "LD H, n8", 2},
		0x27: {g._TODO, "DAA", 1},
		0x28: {g._TODO, "JR Z, e8", 2},
		0x29: {g._TODO, "ADD HL, HL", 1},
		0x2A: {g._LD_A_HLI, "LD A, [HL+]", 1},
		0x2B: {g._TODO, "DEC HL", 1},
		0x2C: {g._TODO, "INC L", 1},
		0x2D: {g._TODO, "DEC L", 1},
		0x2E: {g._LD_L_n8, "LD L, n8", 2},
		0x2F: {g._TODO, "CPL", 1},

		0x30: {g._TODO, "JR NC, e8", 2},
		0x31: {g._LD_SP_n16, "LD SP, n16", 3},
		0x32: {g._LD_HLD_A, "LD [HL-], A", 1},
		0x33: {g._TODO, "INC SP", 1},
		0x34: {g._TODO, "INC [HL]", 1},
		0x35: {g._TODO, "DEC [HL]", 1},
		0x36: {g._LD_HL_n8, "LD [HL], n8", 2},
		0x37: {g._TODO, "SCF", 1},
		0x38: {g._TODO, "JR C, e8", 2},
		0x39: {g._TODO, "ADD HL, SP", 1},
		0x3A: {g._LD_A_HLD, "LD A, [HL-]", 1},
		0x3B: {g._TODO, "DEC SP", 1},
		0x3C: {g._TODO, "INC A", 1},
		0x3D: {g._TODO, "DEC A", 1},
		0x3E: {g._LD_A_n8, "LD A, n8", 2},
		0x3F: {g._TODO, "CCF", 1},

		0x40: {g._LD_B_B, "LD B, B", 1},
		0x41: {g._LD_B_C, "LD B, C", 1},
		0x42: {g._LD_B_D, "LD B, D", 1},
		0x43: {g._LD_B_E, "LD B, E", 1},
		0x44: {g._LD_B_H, "LD B, H", 1},
		0x45: {g._LD_B_L, "LD B, L", 1},
		0x46: {g._LD_B_HL, "LD B, [HL]", 1},
		0x47: {g._LD_B_A, "LD B, A", 1},
		0x48: {g._LD_C_B, "LD C, B", 1},
		0x49: {g._LD_C_C, "LD C, C", 1},
		0x4A: {g._LD_C_D, "LD C, D", 1},
		0x4B: {g._LD_C_E, "LD C, E", 1},
		0x4C: {g._LD_C_H, "LD C, H", 1},
		0x4D: {g._LD_C_L, "LD C, L", 1},
		0x4E: {g._LD_C_HL, "LD C, [HL]", 1},
		0x4F: {g._LD_C_A, "LD C, A", 1},

		0x50: {g._LD_D_B, "LD D, B", 1},
		0x51: {g._LD_D_C, "LD D, C", 1},
		0x52: {g._LD_D_D, "LD D, D", 1},
		0x53: {g._LD_D_E, "LD D, E", 1},
		0x54: {g._LD_D_H, "LD D, H", 1},
		0x55: {g._LD_D_L, "LD D, L", 1},
		0x56: {g._LD_D_HL, "LD D, [HL]", 1},
		0x57: {g._LD_D_A, "LD D, A", 1},
		0x58: {g._LD_E_B, "LD E, B", 1},
		0x59: {g._LD_E_C, "LD E, C", 1},
		0x5A: {g._LD_E_D, "LD E, D", 1},
		0x5B: {g._LD_E_E, "LD E, E", 1},
		0x5C: {g._LD_E_H, "LD E, H", 1},
		0x5D: {g._LD_E_L, "LD E, L", 1},
		0x5E: {g._LD_E_HL, "LD E, [HL]", 1},
		0x5F: {g._LD_E_A, "LD E, A", 1},

		0x60: {g._LD_H_B, "LD H, B", 1},
		0x61: {g._LD_H_C, "LD H, C", 1},
		0x62: {g._LD_H_D, "LD H, D", 1},
		0x63: {g._LD_H_E, "LD H, E", 1},
		0x64: {g._LD_H_H, "LD H, H", 1},
		0x65: {g._LD_H_L, "LD H, L", 1},
		0x66: {g._LD_H_HL, "LD H, [HL]", 1},
		0x67: {g._LD_H_A, "LD H, A", 1},
		0x68: {g._LD_L_B, "LD L, B", 1},
		0x69: {g._LD_L_C, "LD L, C", 1},
		0x6A: {g._LD_L_D, "LD L, D", 1},
		0x6B: {g._LD_L_E, "LD L, E", 1},
		0x6C: {g._LD_L_H, "LD L, H", 1},
		0x6D: {g._LD_L_L, "LD L, L", 1},
		0x6E: {g._LD_L_HL, "LD L, [HL]", 1},
		0x6F: {g._LD_L_A, "LD L, A", 1},

		0x70: {g._LD_HL_B, "LD HL, B", 1},
		0x71: {g._LD_HL_C, "LD HL, C", 1},
		0x72: {g._LD_HL_D, "LD HL, D", 1},
		0x73: {g._LD_HL_E, "LD HL, E", 1},
		0x74: {g._LD_HL_H, "LD HL, H", 1},
		0x75: {g._LD_HL_L, "LD HL, L", 1},
		0x76: {g._TODO, "HALT", 1},
		0x77: {g._LD_HL_A, "LD HL, A", 1},
		0x78: {g._LD_A_B, "LD A, B", 1},
		0x79: {g._LD_A_C, "LD A, C", 1},
		0x7A: {g._LD_A_D, "LD A, D", 1},
		0x7B: {g._LD_A_E, "LD A, E", 1},
		0x7C: {g._LD_A_H, "LD A, H", 1},
		0x7D: {g._LD_A_L, "LD A, L", 1},
		0x7E: {g._LD_A_HL, "LD A, [HL]", 1},
		0x7F: {g._LD_A_A, "LD A, A", 1},

		0x80: {g._TODO, "ADD A, B", 1},
		0x81: {g._TODO, "ADD A, C", 1},
		0x82: {g._TODO, "ADD A, D", 1},
		0x83: {g._TODO, "ADD A, E", 1},
		0x84: {g._TODO, "ADD A, H", 1},
		0x85: {g._TODO, "ADD A, L", 1},
		0x86: {g._TODO, "ADD A, [HL]", 1},
		0x87: {g._TODO, "ADD A, A", 1},
		0x88: {g._TODO, "ADC A, B", 1},
		0x89: {g._TODO, "ADC A, C", 1},
		0x8A: {g._TODO, "ADC A, D", 1},
		0x8B: {g._TODO, "ADC A, E", 1},
		0x8C: {g._TODO, "ADC A, H", 1},
		0x8D: {g._TODO, "ADC A, L", 1},
		0x8E: {g._TODO, "ADC A, [HL]", 1},
		0x8F: {g._TODO, "ADC A, A", 1},

		0x90: {g._TODO, "SUB A, B", 1},
		0x91: {g._TODO, "SUB A, C", 1},
		0x92: {g._TODO, "SUB A, D", 1},
		0x93: {g._TODO, "SUB A, E", 1},
		0x94: {g._TODO, "SUB A, H", 1},
		0x95: {g._TODO, "SUB A, L", 1},
		0x96: {g._TODO, "SUB A, [HL]", 1},
		0x97: {g._TODO, "SUB A, A", 1},
		0x98: {g._TODO, "SBC A, B", 1},
		0x99: {g._TODO, "SBC A, C", 1},
		0x9A: {g._TODO, "SBC A, D", 1},
		0x9B: {g._TODO, "SBC A, E", 1},
		0x9C: {g._TODO, "SBC A, H", 1},
		0x9D: {g._TODO, "SBC A, L", 1},
		0x9E: {g._TODO, "SBC A, [HL]", 1},
		0x9F: {g._TODO, "SBC A, A", 1},

		0xA0: {g._TODO, "AND A, B", 1},
		0xA1: {g._TODO, "AND A, C", 1},
		0xA2: {g._TODO, "AND A, D", 1},
		0xA3: {g._TODO, "AND A, E", 1},
		0xA4: {g._TODO, "AND A, H", 1},
		0xA5: {g._TODO, "AND A, L", 1},
		0xA6: {g._TODO, "AND A, [HL]", 1},
		0xA7: {g._TODO, "AND A, A", 1},
		0xA8: {g._TODO, "XOR A, B", 1},
		0xA9: {g._TODO, "XOR A, C", 1},
		0xAA: {g._TODO, "XOR A, D", 1},
		0xAB: {g._TODO, "XOR A, E", 1},
		0xAC: {g._TODO, "XOR A, H", 1},
		0xAD: {g._TODO, "XOR A, L", 1},
		0xAE: {g._TODO, "XOR A, [HL]", 1},
		0xAF: {g._TODO, "XOR A, A", 1},

		0xB0: {g._TODO, "OR A, B", 1},
		0xB1: {g._TODO, "OR A, C", 1},
		0xB2: {g._TODO, "OR A, D", 1},
		0xB3: {g._TODO, "OR A, E", 1},
		0xB4: {g._TODO, "OR A, H", 1},
		0xB5: {g._TODO, "OR A, L", 1},
		0xB6: {g._TODO, "OR A, [HL]", 1},
		0xB7: {g._TODO, "OR A, A", 1},
		0xB8: {g._TODO, "CP A, B", 1},
		0xB9: {g._TODO, "CP A, C", 1},
		0xBA: {g._TODO, "CP A, D", 1},
		0xBB: {g._TODO, "CP A, E", 1},
		0xBC: {g._TODO, "CP A, H", 1},
		0xBD: {g._TODO, "CP A, L", 1},
		0xBE: {g._TODO, "CP A, [HL]", 1},
		0xBF: {g._TODO, "CP A, A", 1},

		0xC0: {g._TODO, "RET NZ", 1},
		0xC1: {g._TODO, "POP BC", 1},
		0xC2: {g._TODO, "JP NZ, a16", 3},
		0xC3: {g._TODO, "JP a16", 3},
		0xC4: {g._TODO, "CALL NZ, a16", 3},
		0xC5: {g._TODO, "PUSH BC", 1},
		0xC6: {g._TODO, "ADD A, n8", 2},
		0xC7: {g._TODO, "RST $00", 1},
		0xC8: {g._TODO, "RET Z", 1},
		0xC9: {g._TODO, "RET", 1},
		0xCA: {g._TODO, "JP Z, a16", 3},
		0xCB: {g._TODO, "PREFIX", 1},
		0xCC: {g._TODO, "CALL Z, a16", 3},
		0xCD: {g._TODO, "CALL a16", 3},
		0xCE: {g._TODO, "ADC A, n8", 2},
		0xCF: {g._TODO, "RST $08", 1},

		0xD0: {g._TODO, "RET NC", 1},
		0xD1: {g._TODO, "POP DE", 1},
		0xD2: {g._TODO, "JP NC, a16", 3},
		0xD3: {g._ILLEGAL, "ILLEGAL_D3", 1},
		0xD4: {g._TODO, "CALL NC, a16", 3},
		0xD5: {g._TODO, "PUSH DE", 1},
		0xD6: {g._TODO, "SUB A, n8", 2},
		0xD7: {g._TODO, "RST $10", 1},
		0xD8: {g._TODO, "RET C", 1},
		0xD9: {g._TODO, "RETI", 1},
		0xDA: {g._TODO, "JP C, a16", 3},
		0xDB: {g._ILLEGAL, "ILLEGAL_DB", 1},
		0xDC: {g._TODO, "CALL C, a16", 3},
		0xDD: {g._ILLEGAL, "ILLEGAL_DD", 1},
		0xDE: {g._TODO, "SBC A, n8", 2},
		0xDF: {g._TODO, "RST $18", 1},

		0xE0: {g._TODO, "LDH [a8], A", 2},
		0xE1: {g._TODO, "POP HL", 1},
		0xE2: {g._TODO, "LDH [C], A", 1},
		0xE3: {g._ILLEGAL, "ILLEGAL_E3", 1},
		0xE4: {g._ILLEGAL, "ILLEGAL_E4", 1},
		0xE5: {g._TODO, "PUSH HL", 1},
		0xE6: {g._TODO, "AND A, n8", 2},
		0xE7: {g._TODO, "RST $20", 1},
		0xE8: {g._TODO, "ADD SP, e8", 2},
		0xE9: {g._TODO, "JP HL", 1},
		0xEA: {g._TODO, "LD [a16], A", 3},
		0xEB: {g._ILLEGAL, "ILLEGAL_EB", 1},
		0xEC: {g._ILLEGAL, "ILLEGAL_EC", 1},
		0xED: {g._ILLEGAL, "ILLEGAL_ED", 1},
		0xEE: {g._TODO, "XOR A, n8", 2},
		0xEF: {g._TODO, "RST $28", 1},

		0xF0: {g._TODO, "LDH A, [a8]", 2},
		0xF1: {g._TODO, "POP AF", 1},
		0xF2: {g._TODO, "LDH A, [C]", 1},
		0xF3: {g._DI, "DI", 1},
		0xF4: {g._NOP, "NOP", 1},
		0xF5: {g._TODO, "PUSH AF", 1},
		0xF6: {g._TODO, "OR A, n8", 2},
		0xF7: {g._TODO, "RST $30", 1},
		0xF8: {g._TODO, "LD HL, SP + e8", 2},
		0xF9: {g._TODO, "LD SP, HL", 1},
		0xFA: {g._TODO, "LD A, [a16]", 3},
		0xFB: {g._EI, "EI", 1},
		0xFC: {g._NOP, "NOP", 1},
		0xFD: {g._ILLEGAL, "ILLEGAL_FD", 1},
		0xFE: {g._TODO, "CP A, n8", 2},
		0xFF: {g._TODO, "RST $38", 1},
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

func (gbz *GBZ80) _ILLEGAL() uint {
	panic("Illegal instruction executed")
}

func (gbz *GBZ80) _DI() uint {
	gbz.interruptEnabled = false
	return 4
}

func (gbz *GBZ80) _EI() uint {
	gbz.interruptEnabled = true
	return 4
}

func (gbz *GBZ80) _LD_BC_n16() uint {
	gbz.ldrpn16(&gbz.b, &gbz.c)
	return 12
}

func (gbz *GBZ80) _LD_BC_A() uint {
	gbz.ldrpr8(&gbz.b, &gbz.c, &gbz.a, 0)
	return 8
}

func (gbz *GBZ80) _LD_B_n8() uint {
	gbz.ldr8n8(&gbz.b)
	return 8
}

func (gbz *GBZ80) _LD_a16_SP() uint {
	hb := byte(gbz.sp >> 8)
	lb := byte(gbz.sp & 0xFF)

	addr := uint16(gbz.mem.Read(gbz.pc)) | (uint16(gbz.mem.Read(gbz.pc+1)) << 8)
	gbz.pc += 2

	gbz.mem.Write(addr, lb)
	gbz.mem.Write(addr+1, hb)

	return 20
}

func (gbz *GBZ80) _LD_A_BC() uint {
	gbz.ldr8rp(&gbz.a, &gbz.b, &gbz.c, 0)
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
	gbz.ldrpr8(&gbz.d, &gbz.e, &gbz.a, 0)
	return 8
}

func (gbz *GBZ80) _LD_D_n8() uint {
	gbz.ldr8n8(&gbz.d)
	return 8
}

func (gbz *GBZ80) _LD_A_DE() uint {
	gbz.ldr8rp(&gbz.a, &gbz.d, &gbz.e, 0)
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

func (gbz *GBZ80) _LD_HLI_A() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.a, 1)
	return 8
}

func (gbz *GBZ80) _LD_H_n8() uint {
	gbz.ldr8n8(&gbz.h)
	return 8
}

func (gbz *GBZ80) _LD_A_HLI() uint {
	gbz.ldr8rp(&gbz.a, &gbz.h, &gbz.l, 1)
	return 8
}

func (gbz *GBZ80) _LD_L_n8() uint {
	gbz.ldr8n8(&gbz.l)
	return 8
}

func (gbz *GBZ80) _LD_SP_n16() uint {
	lb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	hb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	gbz.sp = (uint16(hb) << 8) | uint16(lb)
	return 12
}

func (gbz *GBZ80) _LD_HLD_A() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.a, -1)
	return 8
}

func (gbz *GBZ80) _LD_HL_n8() uint {
	n := gbz.mem.Read(gbz.pc)
	gbz.pc++

	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.mem.Write(addr, n)

	return 12
}

func (gbz *GBZ80) _LD_A_HLD() uint {
	gbz.ldr8rp(&gbz.a, &gbz.h, &gbz.l, -1)
	return 8
}

func (gbz *GBZ80) _LD_A_n8() uint {
	gbz.ldr8n8(&gbz.a)
	return 8
}

func (gbz *GBZ80) _LD_B_B() uint {
	gbz.ldr8r8(&gbz.b, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_B_C() uint {
	gbz.ldr8r8(&gbz.b, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_B_D() uint {
	gbz.ldr8r8(&gbz.b, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_B_E() uint {
	gbz.ldr8r8(&gbz.b, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_B_H() uint {
	gbz.ldr8r8(&gbz.b, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_B_L() uint {
	gbz.ldr8r8(&gbz.b, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_B_HL() uint {
	gbz.ldr8rp(&gbz.b, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_B_A() uint {
	gbz.ldr8r8(&gbz.b, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_C_B() uint {
	gbz.ldr8r8(&gbz.c, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_C_C() uint {
	gbz.ldr8r8(&gbz.c, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_C_D() uint {
	gbz.ldr8r8(&gbz.c, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_C_E() uint {
	gbz.ldr8r8(&gbz.c, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_C_H() uint {
	gbz.ldr8r8(&gbz.c, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_C_L() uint {
	gbz.ldr8r8(&gbz.c, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_C_HL() uint {
	gbz.ldr8rp(&gbz.c, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_C_A() uint {
	gbz.ldr8r8(&gbz.c, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_D_B() uint {
	gbz.ldr8r8(&gbz.d, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_D_C() uint {
	gbz.ldr8r8(&gbz.d, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_D_D() uint {
	gbz.ldr8r8(&gbz.d, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_D_E() uint {
	gbz.ldr8r8(&gbz.d, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_D_H() uint {
	gbz.ldr8r8(&gbz.d, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_D_L() uint {
	gbz.ldr8r8(&gbz.d, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_D_HL() uint {
	gbz.ldr8rp(&gbz.d, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_D_A() uint {
	gbz.ldr8r8(&gbz.d, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_E_B() uint {
	gbz.ldr8r8(&gbz.e, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_E_C() uint {
	gbz.ldr8r8(&gbz.e, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_E_D() uint {
	gbz.ldr8r8(&gbz.e, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_E_E() uint {
	gbz.ldr8r8(&gbz.e, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_E_H() uint {
	gbz.ldr8r8(&gbz.e, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_E_L() uint {
	gbz.ldr8r8(&gbz.e, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_E_HL() uint {
	gbz.ldr8rp(&gbz.e, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_E_A() uint {
	gbz.ldr8r8(&gbz.e, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_H_B() uint {
	gbz.ldr8r8(&gbz.h, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_H_C() uint {
	gbz.ldr8r8(&gbz.h, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_H_D() uint {
	gbz.ldr8r8(&gbz.h, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_H_E() uint {
	gbz.ldr8r8(&gbz.h, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_H_H() uint {
	gbz.ldr8r8(&gbz.h, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_H_L() uint {
	gbz.ldr8r8(&gbz.h, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_H_HL() uint {
	gbz.ldr8rp(&gbz.h, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_H_A() uint {
	gbz.ldr8r8(&gbz.h, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_L_B() uint {
	gbz.ldr8r8(&gbz.l, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_L_C() uint {
	gbz.ldr8r8(&gbz.l, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_L_D() uint {
	gbz.ldr8r8(&gbz.l, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_L_E() uint {
	gbz.ldr8r8(&gbz.l, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_L_H() uint {
	gbz.ldr8r8(&gbz.l, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_L_L() uint {
	gbz.ldr8r8(&gbz.l, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_L_HL() uint {
	gbz.ldr8rp(&gbz.l, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_L_A() uint {
	gbz.ldr8r8(&gbz.l, &gbz.a)
	return 4
}

func (gbz *GBZ80) _LD_HL_B() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.b, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_C() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.c, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_D() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.d, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_E() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.e, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_H() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.h, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_L() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_HL_A() uint {
	gbz.ldrpr8(&gbz.h, &gbz.l, &gbz.a, 0)
	return 8
}

func (gbz *GBZ80) _LD_A_B() uint {
	gbz.ldr8r8(&gbz.a, &gbz.b)
	return 4
}

func (gbz *GBZ80) _LD_A_C() uint {
	gbz.ldr8r8(&gbz.a, &gbz.c)
	return 4
}

func (gbz *GBZ80) _LD_A_D() uint {
	gbz.ldr8r8(&gbz.a, &gbz.d)
	return 4
}

func (gbz *GBZ80) _LD_A_E() uint {
	gbz.ldr8r8(&gbz.a, &gbz.e)
	return 4
}

func (gbz *GBZ80) _LD_A_H() uint {
	gbz.ldr8r8(&gbz.a, &gbz.h)
	return 4
}

func (gbz *GBZ80) _LD_A_L() uint {
	gbz.ldr8r8(&gbz.a, &gbz.l)
	return 4
}

func (gbz *GBZ80) _LD_A_HL() uint {
	gbz.ldr8rp(&gbz.a, &gbz.h, &gbz.l, 0)
	return 8
}

func (gbz *GBZ80) _LD_A_A() uint {
	gbz.ldr8r8(&gbz.a, &gbz.a)
	return 4
}
