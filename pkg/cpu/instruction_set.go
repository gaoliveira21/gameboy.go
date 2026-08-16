package cpu

import "log"

func (g *GBZ80) initInstructionSet() {
	g.instructionSet = [256]*instruction{
		0x00: {g._NOP, "NOP", 1},
		0x01: {g._LD_BC_n16, "LD BC, n16", 3},
		0x02: {g._LD_BC_A, "LD [BC], A", 1},
		0x03: {g._INC_BC, "INC BC", 1},
		0x04: {g._INC_B, "INC B", 1},
		0x05: {g._DEC_B, "DEC B", 1},
		0x06: {g._LD_B_n8, "LD B, n8", 2},
		0x07: {g._RLCA, "RLCA", 1},
		0x08: {g._LD_a16_SP, "LD [a16], SP", 3},
		0x09: {g._ADD_HL_BC, "ADD HL, BC", 1},
		0x0A: {g._LD_A_BC, "LD A, [BC]", 1},
		0x0B: {g._DEC_BC, "DEC BC", 1},
		0x0C: {g._INC_C, "INC C", 1},
		0x0D: {g._DEC_C, "DEC C", 1},
		0x0E: {g._LD_C_n8, "LD C, n8", 2},
		0x0F: {g._RRCA, "RRCA", 1},

		0x10: {g._STOP, "STOP n8", 2},
		0x11: {g._LD_DE_n16, "LD DE, n16", 3},
		0x12: {g._LD_DE_A, "LD [DE], A", 1},
		0x13: {g._INC_DE, "INC DE", 1},
		0x14: {g._INC_D, "INC D", 1},
		0x15: {g._DEC_D, "DEC D", 1},
		0x16: {g._LD_D_n8, "LD D, n8", 2},
		0x17: {g._RLA, "RLA", 1},
		0x18: {g._JR_E8, "JR e8", 2},
		0x19: {g._ADD_HL_DE, "ADD HL, DE", 1},
		0x1A: {g._LD_A_DE, "LD A, [DE]", 1},
		0x1B: {g._DEC_DE, "DEC DE", 1},
		0x1C: {g._INC_E, "INC E", 1},
		0x1D: {g._DEC_E, "DEC E", 1},
		0x1E: {g._LD_E_n8, "LD E, n8", 2},
		0x1F: {g._RRA, "RRA", 1},

		0x20: {g._JR_NZ_E8, "JR NZ, e8", 2},
		0x21: {g._LD_HL_n16, "LD HL, n16 ", 3},
		0x22: {g._LD_HLI_A, "LD [HL+], A", 1},
		0x23: {g._INC_HL, "INC HL", 1},
		0x24: {g._INC_H, "INC H", 1},
		0x25: {g._DEC_H, "DEC H", 1},
		0x26: {g._LD_H_n8, "LD H, n8", 2},
		0x27: {g._DAA, "DAA", 1},
		0x28: {g._JR_Z_E8, "JR Z, e8", 2},
		0x29: {g._ADD_HL_HL, "ADD HL, HL", 1},
		0x2A: {g._LD_A_HLI, "LD A, [HL+]", 1},
		0x2B: {g._DEC_HL, "DEC HL", 1},
		0x2C: {g._INC_L, "INC L", 1},
		0x2D: {g._DEC_L, "DEC L", 1},
		0x2E: {g._LD_L_n8, "LD L, n8", 2},
		0x2F: {g._CPL, "CPL", 1},

		0x30: {g._JR_NC_E8, "JR NC, e8", 2},
		0x31: {g._LD_SP_n16, "LD SP, n16", 3},
		0x32: {g._LD_HLD_A, "LD [HL-], A", 1},
		0x33: {g._INC_SP, "INC SP", 1},
		0x34: {g._INC_HLX, "INC [HL]", 1},
		0x35: {g._DEC_HLX, "DEC [HL]", 1},
		0x36: {g._LD_HL_n8, "LD [HL], n8", 2},
		0x37: {g._SCF, "SCF", 1},
		0x38: {g._JR_C_E8, "JR C, e8", 2},
		0x39: {g._ADD_HL_SP, "ADD HL, SP", 1},
		0x3A: {g._LD_A_HLD, "LD A, [HL-]", 1},
		0x3B: {g._DEC_SP, "DEC SP", 1},
		0x3C: {g._INC_A, "INC A", 1},
		0x3D: {g._DEC_A, "DEC A", 1},
		0x3E: {g._LD_A_n8, "LD A, n8", 2},
		0x3F: {g._CCF, "CCF", 1},

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
		0x76: {g._HALT, "HALT", 1},
		0x77: {g._LD_HL_A, "LD HL, A", 1},
		0x78: {g._LD_A_B, "LD A, B", 1},
		0x79: {g._LD_A_C, "LD A, C", 1},
		0x7A: {g._LD_A_D, "LD A, D", 1},
		0x7B: {g._LD_A_E, "LD A, E", 1},
		0x7C: {g._LD_A_H, "LD A, H", 1},
		0x7D: {g._LD_A_L, "LD A, L", 1},
		0x7E: {g._LD_A_HL, "LD A, [HL]", 1},
		0x7F: {g._LD_A_A, "LD A, A", 1},

		0x80: {g._ADD_A_B, "ADD A, B", 1},
		0x81: {g._ADD_A_C, "ADD A, C", 1},
		0x82: {g._ADD_A_D, "ADD A, D", 1},
		0x83: {g._ADD_A_E, "ADD A, E", 1},
		0x84: {g._ADD_A_H, "ADD A, H", 1},
		0x85: {g._ADD_A_L, "ADD A, L", 1},
		0x86: {g._ADD_A_HLX, "ADD A, [HL]", 1},
		0x87: {g._ADD_A_A, "ADD A, A", 1},
		0x88: {g._ADC_A_B, "ADC A, B", 1},
		0x89: {g._ADC_A_C, "ADC A, C", 1},
		0x8A: {g._ADC_A_D, "ADC A, D", 1},
		0x8B: {g._ADC_A_E, "ADC A, E", 1},
		0x8C: {g._ADC_A_H, "ADC A, H", 1},
		0x8D: {g._ADC_A_L, "ADC A, L", 1},
		0x8E: {g._ADC_A_HLX, "ADC A, [HL]", 1},
		0x8F: {g._ADC_A_A, "ADC A, A", 1},

		0x90: {g._SUB_A_B, "SUB A, B", 1},
		0x91: {g._SUB_A_C, "SUB A, C", 1},
		0x92: {g._SUB_A_D, "SUB A, D", 1},
		0x93: {g._SUB_A_E, "SUB A, E", 1},
		0x94: {g._SUB_A_H, "SUB A, H", 1},
		0x95: {g._SUB_A_L, "SUB A, L", 1},
		0x96: {g._SUB_A_HLX, "SUB A, [HL]", 1},
		0x97: {g._SUB_A_A, "SUB A, A", 1},
		0x98: {g._SBC_A_B, "SBC A, B", 1},
		0x99: {g._SBC_A_C, "SBC A, C", 1},
		0x9A: {g._SBC_A_D, "SBC A, D", 1},
		0x9B: {g._SBC_A_E, "SBC A, E", 1},
		0x9C: {g._SBC_A_H, "SBC A, H", 1},
		0x9D: {g._SBC_A_L, "SBC A, L", 1},
		0x9E: {g._SBC_A_HLX, "SBC A, [HL]", 1},
		0x9F: {g._SBC_A_A, "SBC A, A", 1},

		0xA0: {g._AND_A_B, "AND A, B", 1},
		0xA1: {g._AND_A_C, "AND A, C", 1},
		0xA2: {g._AND_A_D, "AND A, D", 1},
		0xA3: {g._AND_A_E, "AND A, E", 1},
		0xA4: {g._AND_A_H, "AND A, H", 1},
		0xA5: {g._AND_A_L, "AND A, L", 1},
		0xA6: {g._AND_A_HLX, "AND A, [HL]", 1},
		0xA7: {g._AND_A_A, "AND A, A", 1},
		0xA8: {g._XOR_A_B, "XOR A, B", 1},
		0xA9: {g._XOR_A_C, "XOR A, C", 1},
		0xAA: {g._XOR_A_D, "XOR A, D", 1},
		0xAB: {g._XOR_A_E, "XOR A, E", 1},
		0xAC: {g._XOR_A_H, "XOR A, H", 1},
		0xAD: {g._XOR_A_L, "XOR A, L", 1},
		0xAE: {g._XOR_A_HLX, "XOR A, [HL]", 1},
		0xAF: {g._XOR_A_A, "XOR A, A", 1},

		0xB0: {g._OR_A_B, "OR A, B", 1},
		0xB1: {g._OR_A_C, "OR A, C", 1},
		0xB2: {g._OR_A_D, "OR A, D", 1},
		0xB3: {g._OR_A_E, "OR A, E", 1},
		0xB4: {g._OR_A_H, "OR A, H", 1},
		0xB5: {g._OR_A_L, "OR A, L", 1},
		0xB6: {g._OR_A_HLX, "OR A, [HL]", 1},
		0xB7: {g._OR_A_A, "OR A, A", 1},
		0xB8: {g._CP_A_B, "CP A, B", 1},
		0xB9: {g._CP_A_C, "CP A, C", 1},
		0xBA: {g._CP_A_D, "CP A, D", 1},
		0xBB: {g._CP_A_E, "CP A, E", 1},
		0xBC: {g._CP_A_H, "CP A, H", 1},
		0xBD: {g._CP_A_L, "CP A, L", 1},
		0xBE: {g._CP_A_HLX, "CP A, [HL]", 1},
		0xBF: {g._CP_A_A, "CP A, A", 1},

		0xC0: {g._RET_NZ, "RET NZ", 1},
		0xC1: {g._POP_BC, "POP BC", 1},
		0xC2: {g._JP_NZ_A16, "JP NZ, a16", 3},
		0xC3: {g._JP_A16, "JP a16", 3},
		0xC4: {g._CALL_NZ, "CALL NZ, a16", 3},
		0xC5: {g._PUSH_BC, "PUSH BC", 1},
		0xC6: {g._ADD_A_N8, "ADD A, n8", 2},
		0xC7: {g._RST_00, "RST $00", 1},
		0xC8: {g._RET_Z, "RET Z", 1},
		0xC9: {g._RET, "RET", 1},
		0xCA: {g._JP_Z_A16, "JP Z, a16", 3},
		0xCB: {g._PREFIX, "PREFIX", 1},
		0xCC: {g._CALL_Z, "CALL Z, a16", 3},
		0xCD: {g._CALL, "CALL a16", 3},
		0xCE: {g._ADC_A_N8, "ADC A, n8", 2},
		0xCF: {g._RST_08, "RST $08", 1},

		0xD0: {g._RET_NC, "RET NC", 1},
		0xD1: {g._POP_DE, "POP DE", 1},
		0xD2: {g._JP_NC_A16, "JP NC, a16", 3},
		0xD3: {g._ILLEGAL, "ILLEGAL_D3", 1},
		0xD4: {g._CALL_NC, "CALL NC, a16", 3},
		0xD5: {g._PUSH_DE, "PUSH DE", 1},
		0xD6: {g._SUB_A_N8, "SUB A, n8", 2},
		0xD7: {g._RST_10, "RST $10", 1},
		0xD8: {g._RET_C, "RET C", 1},
		0xD9: {g._RETI, "RETI", 1},
		0xDA: {g._JP_C_A16, "JP C, a16", 3},
		0xDB: {g._ILLEGAL, "ILLEGAL_DB", 1},
		0xDC: {g._CALL_C, "CALL C, a16", 3},
		0xDD: {g._ILLEGAL, "ILLEGAL_DD", 1},
		0xDE: {g._SBC_A_N8, "SBC A, n8", 2},
		0xDF: {g._RST_18, "RST $18", 1},

		0xE0: {g._LDH_A8_A, "LDH [a8], A", 2},
		0xE1: {g._POP_HL, "POP HL", 1},
		0xE2: {g._LDH_C_A, "LDH [C], A", 1},
		0xE3: {g._ILLEGAL, "ILLEGAL_E3", 1},
		0xE4: {g._ILLEGAL, "ILLEGAL_E4", 1},
		0xE5: {g._PUSH_HL, "PUSH HL", 1},
		0xE6: {g._AND_A_N8, "AND A, n8", 2},
		0xE7: {g._RST_20, "RST $20", 1},
		0xE8: {g._ADD_SP_E8, "ADD SP, e8", 2},
		0xE9: {g._JP_HL, "JP HL", 1},
		0xEA: {g._LD_A16_A, "LD [a16], A", 3},
		0xEB: {g._ILLEGAL, "ILLEGAL_EB", 1},
		0xEC: {g._ILLEGAL, "ILLEGAL_EC", 1},
		0xED: {g._ILLEGAL, "ILLEGAL_ED", 1},
		0xEE: {g._XOR_A_N8, "XOR A, n8", 2},
		0xEF: {g._RST_28, "RST $28", 1},

		0xF0: {g._LDH_A_A8, "LDH A, [a8]", 2},
		0xF1: {g._POP_AF, "POP AF", 1},
		0xF2: {g._LDH_A_C, "LDH A, [C]", 1},
		0xF3: {g._DI, "DI", 1},
		0xF4: {g._NOP, "NOP", 1},
		0xF5: {g._PUSH_AF, "PUSH AF", 1},
		0xF6: {g._OR_A_N8, "OR A, n8", 2},
		0xF7: {g._RST_30, "RST $30", 1},
		0xF8: {g._LD_HL_SP_E8, "LD HL, SP + e8", 2},
		0xF9: {g._LD_SP_HL, "LD SP, HL", 1},
		0xFA: {g._LD_A_A16, "LD A, [a16]", 3},
		0xFB: {g._EI, "EI", 1},
		0xFC: {g._NOP, "NOP", 1},
		0xFD: {g._ILLEGAL, "ILLEGAL_FD", 1},
		0xFE: {g._CP_A_N8, "CP A, n8", 2},
		0xFF: {g._RST_38, "RST $38", 1},
	}
}

func (gbz *GBZ80) _NOP() uint {
	return 4
}

func (gbz *GBZ80) _ILLEGAL() uint {
	panic("Illegal instruction executed")
}

func (gbz *GBZ80) _DI() uint {
	gbz.InterruptEnabled = false
	return 4
}

func (gbz *GBZ80) _EI() uint {
	gbz.InterruptEnabled = true
	return 4
}

func (gbz *GBZ80) _HALT() uint {
	gbz.Halt = true
	return 4
}

func (gbz *GBZ80) _STOP() uint {
	gbz.pc++ // Skips the immediate byte following the STOP opcode
	gbz.IsStopped = true
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

func (gbz *GBZ80) _LDH_A8_A() uint {
	a8 := gbz.mem.Read(gbz.pc)
	gbz.pc++

	gbz.mem.Write(0xFF00+uint16(a8), gbz.a)
	return 12
}

func (gbz *GBZ80) _LDH_A_A8() uint {
	a8 := gbz.mem.Read(gbz.pc)
	gbz.pc++

	addr := 0xFF00 + uint16(a8)
	gbz.a = gbz.mem.Read(addr)

	return 12
}

func (gbz *GBZ80) _LDH_C_A() uint {
	gbz.mem.Write(0xFF00+uint16(gbz.c), gbz.a)
	return 8
}

func (gbz *GBZ80) _LDH_A_C() uint {
	addr := 0xFF00 + uint16(gbz.c)
	gbz.a = gbz.mem.Read(addr)
	return 8
}

func (gbz *GBZ80) _LD_A16_A() uint {
	lb := gbz.mem.Read(gbz.pc)
	hb := gbz.mem.Read(gbz.pc + 1)
	gbz.pc += 2

	addr := (uint16(hb) << 8) | uint16(lb)
	gbz.mem.Write(addr, gbz.a)

	return 16
}

func (gbz *GBZ80) _LD_A_A16() uint {
	lb := gbz.mem.Read(gbz.pc)
	hb := gbz.mem.Read(gbz.pc + 1)
	gbz.pc += 2

	addr := (uint16(hb) << 8) | uint16(lb)
	gbz.a = gbz.mem.Read(addr)

	return 16
}

func (gbz *GBZ80) _LD_SP_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.sp = hl
	return 8
}

func (gbz *GBZ80) _LD_HL_SP_E8() uint {
	origSp := int32(gbz.sp)
	e8 := int32(int8(gbz.mem.Read(gbz.pc)))
	gbz.pc++

	r := origSp + e8
	tmpV := origSp ^ e8 ^ r

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, (tmpV&0x10) == 0x10)
	gbz.flags.Set(Carry, (tmpV&0x100) == 0x100)

	rU16 := uint16(r)
	gbz.h = uint8(rU16 >> 8)
	gbz.l = uint8(rU16 & 0xFF)

	return 12
}

func (gbz *GBZ80) _INC_BC() uint {
	gbz.incr16(&gbz.b, &gbz.c)
	return 8
}

func (gbz *GBZ80) _INC_DE() uint {
	gbz.incr16(&gbz.d, &gbz.e)
	return 8
}

func (gbz *GBZ80) _INC_HL() uint {
	gbz.incr16(&gbz.h, &gbz.l)
	return 8
}

func (gbz *GBZ80) _INC_SP() uint {
	gbz.sp++
	return 8
}

func (gbz *GBZ80) _INC_B() uint {
	gbz.incr8(&gbz.b)
	return 4
}

func (gbz *GBZ80) _INC_C() uint {
	gbz.incr8(&gbz.c)
	return 4
}

func (gbz *GBZ80) _INC_D() uint {
	gbz.incr8(&gbz.d)
	return 4
}

func (gbz *GBZ80) _INC_E() uint {
	gbz.incr8(&gbz.e)
	return 4
}

func (gbz *GBZ80) _INC_H() uint {
	gbz.incr8(&gbz.h)
	return 4
}

func (gbz *GBZ80) _INC_L() uint {
	gbz.incr8(&gbz.l)
	return 4
}

func (gbz *GBZ80) _INC_A() uint {
	gbz.incr8(&gbz.a)
	return 4
}

func (gbz *GBZ80) _INC_HLX() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	orig := b
	b++
	gbz.mem.Write(hl, b)

	gbz.flags.Set(Zero, b == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, orig&0x0F == 0x0F)

	return 12
}

func (gbz *GBZ80) _DEC_BC() uint {
	gbz.decr16(&gbz.b, &gbz.c)
	return 8
}

func (gbz *GBZ80) _DEC_DE() uint {
	gbz.decr16(&gbz.d, &gbz.e)
	return 8
}

func (gbz *GBZ80) _DEC_HL() uint {
	gbz.decr16(&gbz.h, &gbz.l)
	return 8
}

func (gbz *GBZ80) _DEC_B() uint {
	gbz.decr8(&gbz.b)
	return 4
}

func (gbz *GBZ80) _DEC_C() uint {
	gbz.decr8(&gbz.c)
	return 4
}

func (gbz *GBZ80) _DEC_D() uint {
	gbz.decr8(&gbz.d)
	return 4
}

func (gbz *GBZ80) _DEC_E() uint {
	gbz.decr8(&gbz.e)
	return 4
}

func (gbz *GBZ80) _DEC_H() uint {
	gbz.decr8(&gbz.h)
	return 4
}

func (gbz *GBZ80) _DEC_L() uint {
	gbz.decr8(&gbz.l)
	return 4
}

func (gbz *GBZ80) _DEC_A() uint {
	gbz.decr8(&gbz.a)
	return 4
}

func (gbz *GBZ80) _DEC_HLX() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	orig := b
	b--
	gbz.mem.Write(hl, b)

	gbz.flags.Set(Zero, b == 0)
	gbz.flags.Set(Sub, true)
	gbz.flags.Set(HalfCarry, orig&0x0F == 0)

	return 12
}

func (gbz *GBZ80) _DEC_SP() uint {
	gbz.sp--
	return 8
}

func (gbz *GBZ80) _CPL() uint {
	gbz.a = ^gbz.a
	gbz.flags.Set(Sub, true)
	gbz.flags.Set(HalfCarry, true)
	return 4
}

func (gbz *GBZ80) _SCF() uint {
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, true)
	return 4
}

func (gbz *GBZ80) _CCF() uint {
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, !gbz.flags.Get(Carry))
	return 4
}

func (gbz *GBZ80) _DAA() uint {
	adjust := uint8(0)

	if gbz.flags.Get(Sub) {
		if gbz.flags.Get(HalfCarry) {
			adjust += 0x6
		}

		if gbz.flags.Get(Carry) {
			adjust += 0x60
		}

		gbz.a -= adjust
	} else {
		if gbz.flags.Get(HalfCarry) || gbz.a&0xF > 0x9 {
			adjust += 0x6
		}

		if gbz.flags.Get(Carry) || gbz.a > 0x99 {
			adjust += 0x60
			gbz.flags.Set(Carry, true)
		}

		gbz.a += adjust
	}

	gbz.flags.Set(Zero, gbz.a == 0)
	gbz.flags.Set(HalfCarry, false)

	return 4
}

func (gbz *GBZ80) _ADD_A_N8() uint {
	n8 := gbz.mem.Read(gbz.pc)
	gbz.pc++
	gbz.add(n8, 0)
	return 8
}

func (gbz *GBZ80) _ADD_SP_E8() uint {
	origSp := gbz.sp
	e8 := int8(gbz.mem.Read(gbz.pc))
	gbz.pc++

	r := uint16(int32(gbz.sp) + int32(e8))
	gbz.sp = r

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, (origSp^r)&0x10 == 0x10)
	gbz.flags.Set(Carry, (origSp^r)&0x100 == 0x100)

	return 16
}

func (gbz *GBZ80) _ADD_A_B() uint {
	gbz.add(gbz.b, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_C() uint {
	gbz.add(gbz.c, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_D() uint {
	gbz.add(gbz.d, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_E() uint {
	gbz.add(gbz.e, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_H() uint {
	gbz.add(gbz.h, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_L() uint {
	gbz.add(gbz.l, 0)
	return 4
}

func (gbz *GBZ80) _ADD_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.add(gbz.mem.Read(addr), 0)
	return 8
}

func (gbz *GBZ80) _ADD_A_A() uint {
	gbz.add(gbz.a, 0)
	return 4
}

func (gbz *GBZ80) _ADD_HL_BC() uint {
	bc := (uint16(gbz.b) << 8) | uint16(gbz.c)
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)

	r := uint32(hl) + uint32(bc)

	gbz.h = uint8(r >> 8)
	gbz.l = uint8(r & 0xFF)

	gbz.flags.Set(Carry, r > 0xFFFF)
	gbz.flags.Set(HalfCarry, (uint32(hl)^uint32(bc)^r)&0x1000 > 0)
	gbz.flags.Set(Sub, false)

	return 8
}

func (gbz *GBZ80) _ADD_HL_DE() uint {
	de := (uint16(gbz.d) << 8) | uint16(gbz.e)
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)

	r := uint32(hl) + uint32(de)

	gbz.h = uint8(r >> 8)
	gbz.l = uint8(r & 0xFF)

	gbz.flags.Set(Carry, r > 0xFFFF)
	gbz.flags.Set(HalfCarry, (uint32(hl)^uint32(de)^r)&0x1000 > 0)
	gbz.flags.Set(Sub, false)

	return 8
}

func (gbz *GBZ80) _ADD_HL_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)

	r := uint32(hl) * 2

	gbz.h = uint8(r >> 8)
	gbz.l = uint8(r & 0xFF)

	gbz.flags.Set(Carry, r > 0xFFFF)
	gbz.flags.Set(HalfCarry, r&0x1000 > 0)
	gbz.flags.Set(Sub, false)

	return 8
}

func (gbz *GBZ80) _ADD_HL_SP() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)

	r := uint32(hl) + uint32(gbz.sp)

	gbz.h = uint8(r >> 8)
	gbz.l = uint8(r & 0xFF)

	gbz.flags.Set(Carry, r > 0xFFFF)
	gbz.flags.Set(HalfCarry, (uint32(hl)^uint32(gbz.sp)^r)&0x1000 > 0)
	gbz.flags.Set(Sub, false)

	return 8
}

func (gbz *GBZ80) _ADC_A_B() uint {
	gbz.adc(gbz.b)
	return 4
}

func (gbz *GBZ80) _ADC_A_C() uint {
	gbz.adc(gbz.c)
	return 4
}

func (gbz *GBZ80) _ADC_A_D() uint {
	gbz.adc(gbz.d)
	return 4
}

func (gbz *GBZ80) _ADC_A_E() uint {
	gbz.adc(gbz.e)
	return 4
}

func (gbz *GBZ80) _ADC_A_H() uint {
	gbz.adc(gbz.h)
	return 4
}

func (gbz *GBZ80) _ADC_A_L() uint {
	gbz.adc(gbz.l)
	return 4
}

func (gbz *GBZ80) _ADC_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.adc(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _ADC_A_A() uint {
	gbz.adc(gbz.a)
	return 4
}

func (gbz *GBZ80) _ADC_A_N8() uint {
	gbz.adc(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _SUB_A_B() uint {
	gbz.sub(gbz.b, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_C() uint {
	gbz.sub(gbz.c, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_D() uint {
	gbz.sub(gbz.d, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_E() uint {
	gbz.sub(gbz.e, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_H() uint {
	gbz.sub(gbz.h, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_L() uint {
	gbz.sub(gbz.l, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.sub(gbz.mem.Read(addr), 0)
	return 8
}

func (gbz *GBZ80) _SUB_A_A() uint {
	gbz.sub(gbz.a, 0)
	return 4
}

func (gbz *GBZ80) _SUB_A_N8() uint {
	gbz.sub(gbz.mem.Read(gbz.pc), 0)
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _SBC_A_B() uint {
	gbz.sbc(gbz.b)
	return 4
}

func (gbz *GBZ80) _SBC_A_C() uint {
	gbz.sbc(gbz.c)
	return 4
}

func (gbz *GBZ80) _SBC_A_D() uint {
	gbz.sbc(gbz.d)
	return 4
}

func (gbz *GBZ80) _SBC_A_E() uint {
	gbz.sbc(gbz.e)
	return 4
}

func (gbz *GBZ80) _SBC_A_H() uint {
	gbz.sbc(gbz.h)
	return 4
}

func (gbz *GBZ80) _SBC_A_L() uint {
	gbz.sbc(gbz.l)
	return 4
}

func (gbz *GBZ80) _SBC_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.sbc(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _SBC_A_A() uint {
	gbz.sbc(gbz.a)
	return 4
}

func (gbz *GBZ80) _SBC_A_N8() uint {
	gbz.sbc(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _AND_A_B() uint {
	gbz.and(gbz.b)
	return 4
}

func (gbz *GBZ80) _XOR_A_B() uint {
	gbz.xor(gbz.b)
	return 4
}

func (gbz *GBZ80) _OR_A_B() uint {
	gbz.or(gbz.b)
	return 4
}

func (gbz *GBZ80) _AND_A_C() uint {
	gbz.and(gbz.c)
	return 4
}

func (gbz *GBZ80) _AND_A_D() uint {
	gbz.and(gbz.d)
	return 4
}

func (gbz *GBZ80) _AND_A_E() uint {
	gbz.and(gbz.e)
	return 4
}

func (gbz *GBZ80) _AND_A_H() uint {
	gbz.and(gbz.h)
	return 4
}

func (gbz *GBZ80) _AND_A_L() uint {
	gbz.and(gbz.l)
	return 4
}

func (gbz *GBZ80) _AND_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.and(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _AND_A_A() uint {
	gbz.and(gbz.a)
	return 4
}

func (gbz *GBZ80) _AND_A_N8() uint {
	gbz.and(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _XOR_A_C() uint {
	gbz.xor(gbz.c)
	return 4
}

func (gbz *GBZ80) _XOR_A_D() uint {
	gbz.xor(gbz.d)
	return 4
}

func (gbz *GBZ80) _XOR_A_E() uint {
	gbz.xor(gbz.e)
	return 4
}

func (gbz *GBZ80) _XOR_A_H() uint {
	gbz.xor(gbz.h)
	return 4
}

func (gbz *GBZ80) _XOR_A_L() uint {
	gbz.xor(gbz.l)
	return 4
}

func (gbz *GBZ80) _XOR_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.xor(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _XOR_A_A() uint {
	gbz.xor(gbz.a)
	return 4
}

func (gbz *GBZ80) _XOR_A_N8() uint {
	gbz.xor(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _OR_A_C() uint {
	gbz.or(gbz.c)
	return 4
}

func (gbz *GBZ80) _OR_A_D() uint {
	gbz.or(gbz.d)
	return 4
}

func (gbz *GBZ80) _OR_A_E() uint {
	gbz.or(gbz.e)
	return 4
}

func (gbz *GBZ80) _OR_A_H() uint {
	gbz.or(gbz.h)
	return 4
}

func (gbz *GBZ80) _OR_A_L() uint {
	gbz.or(gbz.l)
	return 4
}

func (gbz *GBZ80) _OR_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.or(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _OR_A_A() uint {
	gbz.or(gbz.a)
	return 4
}

func (gbz *GBZ80) _OR_A_N8() uint {
	gbz.or(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _CP_A_B() uint {
	gbz.cp(gbz.b)
	return 4
}

func (gbz *GBZ80) _CP_A_C() uint {
	gbz.cp(gbz.c)
	return 4
}

func (gbz *GBZ80) _CP_A_D() uint {
	gbz.cp(gbz.d)
	return 4
}

func (gbz *GBZ80) _CP_A_E() uint {
	gbz.cp(gbz.e)
	return 4
}

func (gbz *GBZ80) _CP_A_H() uint {
	gbz.cp(gbz.h)
	return 4
}

func (gbz *GBZ80) _CP_A_L() uint {
	gbz.cp(gbz.l)
	return 4
}

func (gbz *GBZ80) _CP_A_HLX() uint {
	addr := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.cp(gbz.mem.Read(addr))
	return 8
}

func (gbz *GBZ80) _CP_A_A() uint {
	gbz.cp(gbz.a)
	return 4
}

func (gbz *GBZ80) _CP_A_N8() uint {
	gbz.cp(gbz.mem.Read(gbz.pc))
	gbz.pc++
	return 8
}

func (gbz *GBZ80) _JR_E8() uint {
	gbz.jumpE8()
	return 12
}

func (gbz *GBZ80) _JR_NZ_E8() uint {
	if !gbz.flags.Get(Zero) {
		gbz.jumpE8()
		return 12
	}

	gbz.pc++
	return 8
}

func (gbz *GBZ80) _JR_NC_E8() uint {
	if !gbz.flags.Get(Carry) {
		gbz.jumpE8()
		return 12
	}

	gbz.pc++
	return 8
}

func (gbz *GBZ80) _JR_Z_E8() uint {
	if gbz.flags.Get(Zero) {
		gbz.jumpE8()
		return 12
	}

	gbz.pc++
	return 8
}

func (gbz *GBZ80) _JR_C_E8() uint {
	if gbz.flags.Get(Carry) {
		gbz.jumpE8()
		return 12
	}

	gbz.pc++
	return 8
}

func (gbz *GBZ80) _JP_A16() uint {
	gbz.jumpA16()
	return 16
}

func (gbz *GBZ80) _JP_NZ_A16() uint {
	if !gbz.flags.Get(Zero) {
		gbz.jumpA16()
		return 16
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _JP_NC_A16() uint {
	if !gbz.flags.Get(Carry) {
		gbz.jumpA16()
		return 16
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _JP_Z_A16() uint {
	if gbz.flags.Get(Zero) {
		gbz.jumpA16()
		return 16
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _JP_C_A16() uint {
	if gbz.flags.Get(Carry) {
		gbz.jumpA16()
		return 16
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _JP_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	gbz.pc = hl
	return 4
}

func (gbz *GBZ80) _RET() uint {
	gbz.ret()
	return 16
}

func (gbz *GBZ80) _RET_NZ() uint {
	if !gbz.flags.Get(Zero) {
		gbz.ret()
		return 20
	}

	return 8
}

func (gbz *GBZ80) _RET_NC() uint {
	if !gbz.flags.Get(Carry) {
		gbz.ret()
		return 20
	}

	return 8
}

func (gbz *GBZ80) _RET_Z() uint {
	if gbz.flags.Get(Zero) {
		gbz.ret()
		return 20
	}

	return 8
}

func (gbz *GBZ80) _RET_C() uint {
	if gbz.flags.Get(Carry) {
		gbz.ret()
		return 20
	}

	return 8
}

func (gbz *GBZ80) _RETI() uint {
	gbz.ret()
	gbz.InterruptEnabled = true
	return 16
}

func (gbz *GBZ80) _CALL() uint {
	gbz.call()
	return 24
}

func (gbz *GBZ80) _CALL_NZ() uint {
	if !gbz.flags.Get(Zero) {
		gbz.call()
		return 24
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _CALL_NC() uint {
	if !gbz.flags.Get(Carry) {
		gbz.call()
		return 24
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _CALL_Z() uint {
	if gbz.flags.Get(Zero) {
		gbz.call()
		return 24
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _CALL_C() uint {
	if gbz.flags.Get(Carry) {
		gbz.call()
		return 24
	}

	gbz.pc += 2
	return 12
}

func (gbz *GBZ80) _RST_00() uint {
	gbz.rst(0x0000)
	return 16
}

func (gbz *GBZ80) _RST_08() uint {
	gbz.rst(0x0008)
	return 16
}

func (gbz *GBZ80) _RST_10() uint {
	gbz.rst(0x0010)
	return 16
}

func (gbz *GBZ80) _RST_18() uint {
	gbz.rst(0x0018)
	return 16
}

func (gbz *GBZ80) _RST_20() uint {
	gbz.rst(0x0020)
	return 16
}

func (gbz *GBZ80) _RST_28() uint {
	gbz.rst(0x0028)
	return 16
}

func (gbz *GBZ80) _RST_30() uint {
	gbz.rst(0x0030)
	return 16
}

func (gbz *GBZ80) _RST_38() uint {
	gbz.rst(0x0038)
	return 16
}

func (gbz *GBZ80) _POP_BC() uint {
	hb, lb := gbz.pop()
	gbz.b = hb
	gbz.c = lb
	return 12
}

func (gbz *GBZ80) _POP_DE() uint {
	hb, lb := gbz.pop()
	gbz.d = hb
	gbz.e = lb
	return 12
}

func (gbz *GBZ80) _POP_HL() uint {
	hb, lb := gbz.pop()
	gbz.h = hb
	gbz.l = lb
	return 12
}

func (gbz *GBZ80) _POP_AF() uint {
	hb, lb := gbz.pop()

	v := ((uint16(hb) << 8) | uint16(lb)) & 0xFFF0

	gbz.a = byte((v & 0xFF00) >> 8)
	gbz.flags.SetValue(byte(v & 0xFF))

	return 12
}

func (gbz *GBZ80) _PUSH_BC() uint {
	gbz.push(gbz.b, gbz.c)
	return 16
}

func (gbz *GBZ80) _PUSH_DE() uint {
	gbz.push(gbz.d, gbz.e)
	return 16
}

func (gbz *GBZ80) _PUSH_HL() uint {
	gbz.push(gbz.h, gbz.l)
	return 16
}

func (gbz *GBZ80) _PUSH_AF() uint {
	gbz.push(gbz.a, gbz.flags.Value())
	return 16
}

func (gbz *GBZ80) _RLCA() uint {
	a := gbz.a

	gbz.a = gbz.a<<1 | gbz.a>>7

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, a > 0x7F)

	return 4
}

func (gbz *GBZ80) _RRCA() uint {
	a := gbz.a

	gbz.a = a>>1 | (a&0x1)<<7

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, gbz.a > 0x7F)

	return 4
}

func (gbz *GBZ80) _RLA() uint {
	c := uint8(0x0)
	if gbz.flags.Get(Carry) {
		c = 0x1
	}

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, gbz.a&0x80 > 0x0)

	gbz.a = gbz.a<<1 | c

	return 4
}

func (gbz *GBZ80) _RRA() uint {
	c := uint8(0x0)
	if gbz.flags.Get(Carry) {
		c = 0x80
	}

	gbz.flags.Set(Zero, false)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, gbz.a&0x1 > 0x0)

	gbz.a = c | gbz.a>>1

	return 4
}

func (gbz *GBZ80) _PREFIX() uint {
	cb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	nextInstr := gbz.prefixedSet[cb]
	if nextInstr != nil {
		cycles := nextInstr.operation()
		return cycles
	}

	log.Fatalf("Unknown CB opcode: %X \n", cb)

	return 4
}
