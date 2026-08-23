package cpu

func (g *GBZ80) initPrefixedSet() {
	g.prefixedSet = [256]*instruction{
		0x00: {g._RLC_B, "RLC B", 2},
		0x01: {g._RLC_C, "RLC C", 2},
		0x02: {g._RLC_D, "RLC D", 2},
		0x03: {g._RLC_E, "RLC E", 2},
		0x04: {g._RLC_H, "RLC H", 2},
		0x05: {g._RLC_L, "RLC L", 2},
		0x06: {g._RLC_HL, "RLC HL", 2},
		0x07: {g._RLC_A, "RLC A", 2},
		0x08: {g._RRC_B, "RRC B", 2},
		0x09: {g._RRC_C, "RRC C", 2},
		0x0A: {g._RRC_D, "RRC D", 2},
		0x0B: {g._RRC_E, "RRC E", 2},
		0x0C: {g._RRC_H, "RRC H", 2},
		0x0D: {g._RRC_L, "RRC L", 2},
		0x0E: {g._RRC_HL, "RRC HL", 2},
		0x0F: {g._RRC_A, "RRC A", 2},
		0x10: {g._RL_B, "RL B", 2},
		0x11: {g._RL_C, "RL C", 2},
		0x12: {g._RL_D, "RL D", 2},
		0x13: {g._RL_E, "RL E", 2},
		0x14: {g._RL_H, "RL H", 2},
		0x15: {g._RL_L, "RL L", 2},
		0x16: {g._RL_HL, "RL HL", 2},
		0x17: {g._RL_A, "RL A", 2},
		0x18: {g._RR_B, "RR B", 2},
		0x19: {g._RR_C, "RR C", 2},
		0x1A: {g._RR_D, "RR D", 2},
		0x1B: {g._RR_E, "RR E", 2},
		0x1C: {g._RR_H, "RR H", 2},
		0x1D: {g._RR_L, "RR L", 2},
		0x1E: {g._RR_HL, "RR HL", 2},
		0x1F: {g._RR_A, "RR A", 2},
		0x20: {g._SLA_B, "SLA B", 2},
		0x21: {g._SLA_C, "SLA C", 2},
		0x22: {g._SLA_D, "SLA D", 2},
		0x23: {g._SLA_E, "SLA E", 2},
		0x24: {g._SLA_H, "SLA H", 2},
		0x25: {g._SLA_L, "SLA L", 2},
		0x26: {g._SLA_HL, "SLA HL", 2},
		0x27: {g._SLA_A, "SLA A", 2},
		0x28: {g._SRA_B, "SRA B", 2},
		0x29: {g._SRA_C, "SRA C", 2},
		0x2A: {g._SRA_D, "SRA D", 2},
		0x2B: {g._SRA_E, "SRA E", 2},
		0x2C: {g._SRA_H, "SRA H", 2},
		0x2D: {g._SRA_L, "SRA L", 2},
		0x2E: {g._SRA_HL, "SRA HL", 2},
		0x2F: {g._SRA_A, "SRA A", 2},
		0x30: {g._SWAP_B, "SWAP B", 2},
		0x31: {g._SWAP_C, "SWAP C", 2},
		0x32: {g._SWAP_D, "SWAP D", 2},
		0x33: {g._SWAP_E, "SWAP E", 2},
		0x34: {g._SWAP_H, "SWAP H", 2},
		0x35: {g._SWAP_L, "SWAP L", 2},
		0x36: {g._SWAP_HL, "SWAP HL", 2},
		0x37: {g._SWAP_A, "SWAP A", 2},
		0x38: {g._SRL_B, "SRL B", 2},
		0x39: {g._SRL_C, "SRL C", 2},
		0x3A: {g._SRL_D, "SRL D", 2},
		0x3B: {g._SRL_E, "SRL E", 2},
		0x3C: {g._SRL_H, "SRL H", 2},
		0x3D: {g._SRL_L, "SRL L", 2},
		0x3E: {g._SRL_HL, "SRL HL", 2},
		0x3F: {g._SRL_A, "SRL A", 2},
		0x40: {g._BIT_0_B, "BIT 0 B", 2},
		0x41: {g._BIT_0_C, "BIT 0 C", 2},
		0x42: {g._BIT_0_D, "BIT 0 D", 2},
		0x43: {g._BIT_0_E, "BIT 0 E", 2},
		0x44: {g._BIT_0_H, "BIT 0 H", 2},
		0x45: {g._BIT_0_L, "BIT 0 L", 2},
		0x46: {g._BIT_0_HL, "BIT 0 HL", 2},
		0x47: {g._BIT_0_A, "BIT 0 A", 2},
		0x48: {g._BIT_1_B, "BIT 1 B", 2},
		0x49: {g._BIT_1_C, "BIT 1 C", 2},
		0x4A: {g._BIT_1_D, "BIT 1 D", 2},
		0x4B: {g._BIT_1_E, "BIT 1 E", 2},
		0x4C: {g._BIT_1_H, "BIT 1 H", 2},
		0x4D: {g._BIT_1_L, "BIT 1 L", 2},
		0x4E: {g._BIT_1_HL, "BIT 1 HL", 2},
		0x4F: {g._BIT_1_A, "BIT 1 A", 2},
		0x50: {g._BIT_2_B, "BIT 2 B", 2},
		0x51: {g._BIT_2_C, "BIT 2 C", 2},
		0x52: {g._BIT_2_D, "BIT 2 D", 2},
		0x53: {g._BIT_2_E, "BIT 2 E", 2},
		0x54: {g._BIT_2_H, "BIT 2 H", 2},
		0x55: {g._BIT_2_L, "BIT 2 L", 2},
		0x56: {g._BIT_2_HL, "BIT 2 HL", 2},
		0x57: {g._BIT_2_A, "BIT 2 A", 2},
		0x58: {g._BIT_3_B, "BIT 3 B", 2},
		0x59: {g._BIT_3_C, "BIT 3 C", 2},
		0x5A: {g._BIT_3_D, "BIT 3 D", 2},
		0x5B: {g._BIT_3_E, "BIT 3 E", 2},
		0x5C: {g._BIT_3_H, "BIT 3 H", 2},
		0x5D: {g._BIT_3_L, "BIT 3 L", 2},
		0x5E: {g._BIT_3_HL, "BIT 3 HL", 2},
		0x5F: {g._BIT_3_A, "BIT 3 A", 2},
		0x60: {g._BIT_4_B, "BIT 4 B", 2},
		0x61: {g._BIT_4_C, "BIT 4 C", 2},
		0x62: {g._BIT_4_D, "BIT 4 D", 2},
		0x63: {g._BIT_4_E, "BIT 4 E", 2},
		0x64: {g._BIT_4_H, "BIT 4 H", 2},
		0x65: {g._BIT_4_L, "BIT 4 L", 2},
		0x66: {g._BIT_4_HL, "BIT 4 HL", 2},
		0x67: {g._BIT_4_A, "BIT 4 A", 2},
		0x68: {g._BIT_5_B, "BIT 5 B", 2},
		0x69: {g._BIT_5_C, "BIT 5 C", 2},
		0x6A: {g._BIT_5_D, "BIT 5 D", 2},
		0x6B: {g._BIT_5_E, "BIT 5 E", 2},
		0x6C: {g._BIT_5_H, "BIT 5 H", 2},
		0x6D: {g._BIT_5_L, "BIT 5 L", 2},
		0x6E: {g._BIT_5_HL, "BIT 5 HL", 2},
		0x6F: {g._BIT_5_A, "BIT 5 A", 2},
		0x70: {g._BIT_6_B, "BIT 6 B", 2},
		0x71: {g._BIT_6_C, "BIT 6 C", 2},
		0x72: {g._BIT_6_D, "BIT 6 D", 2},
		0x73: {g._BIT_6_E, "BIT 6 E", 2},
		0x74: {g._BIT_6_H, "BIT 6 H", 2},
		0x75: {g._BIT_6_L, "BIT 6 L", 2},
		0x76: {g._BIT_6_HL, "BIT 6 HL", 2},
		0x77: {g._BIT_6_A, "BIT 6 A", 2},
		0x78: {g._BIT_7_B, "BIT 7 B", 2},
		0x79: {g._BIT_7_C, "BIT 7 C", 2},
		0x7A: {g._BIT_7_D, "BIT 7 D", 2},
		0x7B: {g._BIT_7_E, "BIT 7 E", 2},
		0x7C: {g._BIT_7_H, "BIT 7 H", 2},
		0x7D: {g._BIT_7_L, "BIT 7 L", 2},
		0x7E: {g._BIT_7_HL, "BIT 7 HL", 2},
		0x7F: {g._BIT_7_A, "BIT 7 A", 2},
		0x80: {g._RES_0_B, "RES 0 B", 2},
		0x81: {g._RES_0_C, "RES 0 C", 2},
		0x82: {g._RES_0_D, "RES 0 D", 2},
		0x83: {g._RES_0_E, "RES 0 E", 2},
		0x84: {g._RES_0_H, "RES 0 H", 2},
		0x85: {g._RES_0_L, "RES 0 L", 2},
		0x86: {g._RES_0_HL, "RES 0 HL", 2},
		0x87: {g._RES_0_A, "RES 0 A", 2},
		0x88: {g._RES_1_B, "RES 1 B", 2},
		0x89: {g._RES_1_C, "RES 1 C", 2},
		0x8A: {g._RES_1_D, "RES 1 D", 2},
		0x8B: {g._RES_1_E, "RES 1 E", 2},
		0x8C: {g._RES_1_H, "RES 1 H", 2},
		0x8D: {g._RES_1_L, "RES 1 L", 2},
		0x8E: {g._RES_1_HL, "RES 1 HL", 2},
		0x8F: {g._RES_1_A, "RES 1 A", 2},
		0x90: {g._RES_2_B, "RES 2 B", 2},
		0x91: {g._RES_2_C, "RES 2 C", 2},
		0x92: {g._RES_2_D, "RES 2 D", 2},
		0x93: {g._RES_2_E, "RES 2 E", 2},
		0x94: {g._RES_2_H, "RES 2 H", 2},
		0x95: {g._RES_2_L, "RES 2 L", 2},
		0x96: {g._RES_2_HL, "RES 2 HL", 2},
		0x97: {g._RES_2_A, "RES 2 A", 2},
		0x98: {g._RES_3_B, "RES 3 B", 2},
		0x99: {g._RES_3_C, "RES 3 C", 2},
		0x9A: {g._RES_3_D, "RES 3 D", 2},
		0x9B: {g._RES_3_E, "RES 3 E", 2},
		0x9C: {g._RES_3_H, "RES 3 H", 2},
		0x9D: {g._RES_3_L, "RES 3 L", 2},
		0x9E: {g._RES_3_HL, "RES 3 HL", 2},
		0x9F: {g._RES_3_A, "RES 3 A", 2},
		0xA0: {g._RES_4_B, "RES 4 B", 2},
		0xA1: {g._RES_4_C, "RES 4 C", 2},
		0xA2: {g._RES_4_D, "RES 4 D", 2},
		0xA3: {g._RES_4_E, "RES 4 E", 2},
		0xA4: {g._RES_4_H, "RES 4 H", 2},
		0xA5: {g._RES_4_L, "RES 4 L", 2},
		0xA6: {g._RES_4_HL, "RES 4 HL", 2},
		0xA7: {g._RES_4_A, "RES 4 A", 2},
		0xA8: {g._RES_5_B, "RES 5 B", 2},
		0xA9: {g._RES_5_C, "RES 5 C", 2},
		0xAA: {g._RES_5_D, "RES 5 D", 2},
		0xAB: {g._RES_5_E, "RES 5 E", 2},
		0xAC: {g._RES_5_H, "RES 5 H", 2},
		0xAD: {g._RES_5_L, "RES 5 L", 2},
		0xAE: {g._RES_5_HL, "RES 5 HL", 2},
		0xAF: {g._RES_5_A, "RES 5 A", 2},
		0xB0: {g._RES_6_B, "RES 6 B", 2},
		0xB1: {g._RES_6_C, "RES 6 C", 2},
		0xB2: {g._RES_6_D, "RES 6 D", 2},
		0xB3: {g._RES_6_E, "RES 6 E", 2},
		0xB4: {g._RES_6_H, "RES 6 H", 2},
		0xB5: {g._RES_6_L, "RES 6 L", 2},
		0xB6: {g._RES_6_HL, "RES 6 HL", 2},
		0xB7: {g._RES_6_A, "RES 6 A", 2},
		0xB8: {g._RES_7_B, "RES 7 B", 2},
		0xB9: {g._RES_7_C, "RES 7 C", 2},
		0xBA: {g._RES_7_D, "RES 7 D", 2},
		0xBB: {g._RES_7_E, "RES 7 E", 2},
		0xBC: {g._RES_7_H, "RES 7 H", 2},
		0xBD: {g._RES_7_L, "RES 7 L", 2},
		0xBE: {g._RES_7_HL, "RES 7 HL", 2},
		0xBF: {g._RES_7_A, "RES 7 A", 2},
		0xC0: {g._SET_0_B, "SET 0 B", 2},
		0xC1: {g._SET_0_C, "SET 0 C", 2},
		0xC2: {g._SET_0_D, "SET 0 D", 2},
		0xC3: {g._SET_0_E, "SET 0 E", 2},
		0xC4: {g._SET_0_H, "SET 0 H", 2},
		0xC5: {g._SET_0_L, "SET 0 L", 2},
		0xC6: {g._SET_0_HL, "SET 0 HL", 2},
		0xC7: {g._SET_0_A, "SET 0 A", 2},
		0xC8: {g._SET_1_B, "SET 1 B", 2},
		0xC9: {g._SET_1_C, "SET 1 C", 2},
		0xCA: {g._SET_1_D, "SET 1 D", 2},
		0xCB: {g._SET_1_E, "SET 1 E", 2},
		0xCC: {g._SET_1_H, "SET 1 H", 2},
		0xCD: {g._SET_1_L, "SET 1 L", 2},
		0xCE: {g._SET_1_HL, "SET 1 HL", 2},
		0xCF: {g._SET_1_A, "SET 1 A", 2},
		0xD0: {g._SET_2_B, "SET 2 B", 2},
		0xD1: {g._SET_2_C, "SET 2 C", 2},
		0xD2: {g._SET_2_D, "SET 2 D", 2},
		0xD3: {g._SET_2_E, "SET 2 E", 2},
		0xD4: {g._SET_2_H, "SET 2 H", 2},
		0xD5: {g._SET_2_L, "SET 2 L", 2},
		0xD6: {g._SET_2_HL, "SET 2 HL", 2},
		0xD7: {g._SET_2_A, "SET 2 A", 2},
		0xD8: {g._SET_3_B, "SET 3 B", 2},
		0xD9: {g._SET_3_C, "SET 3 C", 2},
		0xDA: {g._SET_3_D, "SET 3 D", 2},
		0xDB: {g._SET_3_E, "SET 3 E", 2},
		0xDC: {g._SET_3_H, "SET 3 H", 2},
		0xDD: {g._SET_3_L, "SET 3 L", 2},
		0xDE: {g._SET_3_HL, "SET 3 HL", 2},
		0xDF: {g._SET_3_A, "SET 3 A", 2},
		0xE0: {g._SET_4_B, "SET 4 B", 2},
		0xE1: {g._SET_4_C, "SET 4 C", 2},
		0xE2: {g._SET_4_D, "SET 4 D", 2},
		0xE3: {g._SET_4_E, "SET 4 E", 2},
		0xE4: {g._SET_4_H, "SET 4 H", 2},
		0xE5: {g._SET_4_L, "SET 4 L", 2},
		0xE6: {g._SET_4_HL, "SET 4 HL", 2},
		0xE7: {g._SET_4_A, "SET 4 A", 2},
		0xE8: {g._SET_5_B, "SET 5 B", 2},
		0xE9: {g._SET_5_C, "SET 5 C", 2},
		0xEA: {g._SET_5_D, "SET 5 D", 2},
		0xEB: {g._SET_5_E, "SET 5 E", 2},
		0xEC: {g._SET_5_H, "SET 5 H", 2},
		0xED: {g._SET_5_L, "SET 5 L", 2},
		0xEE: {g._SET_5_HL, "SET 5 HL", 2},
		0xEF: {g._SET_5_A, "SET 5 A", 2},
		0xF0: {g._SET_6_B, "SET 6 B", 2},
		0xF1: {g._SET_6_C, "SET 6 C", 2},
		0xF2: {g._SET_6_D, "SET 6 D", 2},
		0xF3: {g._SET_6_E, "SET 6 E", 2},
		0xF4: {g._SET_6_H, "SET 6 H", 2},
		0xF5: {g._SET_6_L, "SET 6 L", 2},
		0xF6: {g._SET_6_HL, "SET 6 HL", 2},
		0xF7: {g._SET_6_A, "SET 6 A", 2},
		0xF8: {g._SET_7_B, "SET 7 B", 2},
		0xF9: {g._SET_7_C, "SET 7 C", 2},
		0xFA: {g._SET_7_D, "SET 7 D", 2},
		0xFB: {g._SET_7_E, "SET 7 E", 2},
		0xFC: {g._SET_7_H, "SET 7 H", 2},
		0xFD: {g._SET_7_L, "SET 7 L", 2},
		0xFE: {g._SET_7_HL, "SET 7 HL", 2},
		0xFF: {g._SET_7_A, "SET 7 A", 2},
	}
}

func (gbz *GBZ80) _RLC_B() uint {
	gbz.rlc(&gbz.b)
	return 8
}

func (gbz *GBZ80) _RLC_C() uint {
	gbz.rlc(&gbz.c)
	return 8
}

func (gbz *GBZ80) _RLC_D() uint {
	gbz.rlc(&gbz.d)
	return 8
}

func (gbz *GBZ80) _RLC_E() uint {
	gbz.rlc(&gbz.e)
	return 8
}

func (gbz *GBZ80) _RLC_H() uint {
	gbz.rlc(&gbz.h)
	return 8
}

func (gbz *GBZ80) _RLC_L() uint {
	gbz.rlc(&gbz.l)
	return 8
}

func (gbz *GBZ80) _RLC_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.rlc(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RLC_A() uint {
	gbz.rlc(&gbz.a)
	return 8
}

func (gbz *GBZ80) _RRC_B() uint {
	gbz.rrc(&gbz.b)
	return 8
}

func (gbz *GBZ80) _RRC_C() uint {
	gbz.rrc(&gbz.c)
	return 8
}

func (gbz *GBZ80) _RRC_D() uint {
	gbz.rrc(&gbz.d)
	return 8
}

func (gbz *GBZ80) _RRC_E() uint {
	gbz.rrc(&gbz.e)
	return 8
}

func (gbz *GBZ80) _RRC_H() uint {
	gbz.rrc(&gbz.h)
	return 8
}

func (gbz *GBZ80) _RRC_L() uint {
	gbz.rrc(&gbz.l)
	return 8
}

func (gbz *GBZ80) _RRC_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.rrc(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RRC_A() uint {
	gbz.rrc(&gbz.a)
	return 8
}

func (gbz *GBZ80) _RL_B() uint {
	gbz.rl(&gbz.b)
	return 8
}

func (gbz *GBZ80) _RL_C() uint {
	gbz.rl(&gbz.c)
	return 8
}

func (gbz *GBZ80) _RL_D() uint {
	gbz.rl(&gbz.d)
	return 8
}

func (gbz *GBZ80) _RL_E() uint {
	gbz.rl(&gbz.e)
	return 8
}

func (gbz *GBZ80) _RL_H() uint {
	gbz.rl(&gbz.h)
	return 8
}

func (gbz *GBZ80) _RL_L() uint {
	gbz.rl(&gbz.l)
	return 8
}

func (gbz *GBZ80) _RL_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.rl(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RL_A() uint {
	gbz.rl(&gbz.a)
	return 8
}

func (gbz *GBZ80) _RR_B() uint {
	gbz.rr(&gbz.b)
	return 8
}

func (gbz *GBZ80) _RR_C() uint {
	gbz.rr(&gbz.c)
	return 8
}

func (gbz *GBZ80) _RR_D() uint {
	gbz.rr(&gbz.d)
	return 8
}

func (gbz *GBZ80) _RR_E() uint {
	gbz.rr(&gbz.e)
	return 8
}

func (gbz *GBZ80) _RR_H() uint {
	gbz.rr(&gbz.h)
	return 8
}

func (gbz *GBZ80) _RR_L() uint {
	gbz.rr(&gbz.l)
	return 8
}

func (gbz *GBZ80) _RR_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.rr(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RR_A() uint {
	gbz.rr(&gbz.a)
	return 8
}

func (gbz *GBZ80) _SLA_B() uint {
	gbz.sla(&gbz.b)
	return 8
}

func (gbz *GBZ80) _SLA_C() uint {
	gbz.sla(&gbz.c)
	return 8
}

func (gbz *GBZ80) _SLA_D() uint {
	gbz.sla(&gbz.d)
	return 8
}

func (gbz *GBZ80) _SLA_E() uint {
	gbz.sla(&gbz.e)
	return 8
}

func (gbz *GBZ80) _SLA_H() uint {
	gbz.sla(&gbz.h)
	return 8
}

func (gbz *GBZ80) _SLA_L() uint {
	gbz.sla(&gbz.l)
	return 8
}

func (gbz *GBZ80) _SLA_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.sla(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SLA_A() uint {
	gbz.sla(&gbz.a)
	return 8
}

func (gbz *GBZ80) _SRA_B() uint {
	gbz.sra(&gbz.b)
	return 8
}

func (gbz *GBZ80) _SRA_C() uint {
	gbz.sra(&gbz.c)
	return 8
}

func (gbz *GBZ80) _SRA_D() uint {
	gbz.sra(&gbz.d)
	return 8
}

func (gbz *GBZ80) _SRA_E() uint {
	gbz.sra(&gbz.e)
	return 8
}

func (gbz *GBZ80) _SRA_H() uint {
	gbz.sra(&gbz.h)
	return 8
}

func (gbz *GBZ80) _SRA_L() uint {
	gbz.sra(&gbz.l)
	return 8
}

func (gbz *GBZ80) _SRA_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.sra(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SRA_A() uint {
	gbz.sra(&gbz.a)
	return 8
}

func (gbz *GBZ80) _SWAP_B() uint {
	gbz.swap(&gbz.b)
	return 8
}

func (gbz *GBZ80) _SWAP_C() uint {
	gbz.swap(&gbz.c)
	return 8
}

func (gbz *GBZ80) _SWAP_D() uint {
	gbz.swap(&gbz.d)
	return 8
}

func (gbz *GBZ80) _SWAP_E() uint {
	gbz.swap(&gbz.e)
	return 8
}

func (gbz *GBZ80) _SWAP_H() uint {
	gbz.swap(&gbz.h)
	return 8
}

func (gbz *GBZ80) _SWAP_L() uint {
	gbz.swap(&gbz.l)
	return 8
}

func (gbz *GBZ80) _SWAP_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.swap(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SWAP_A() uint {
	gbz.swap(&gbz.a)
	return 8
}

func (gbz *GBZ80) _SRL_B() uint {
	gbz.srl(&gbz.b)
	return 8
}

func (gbz *GBZ80) _SRL_C() uint {
	gbz.srl(&gbz.c)
	return 8
}

func (gbz *GBZ80) _SRL_D() uint {
	gbz.srl(&gbz.d)
	return 8
}

func (gbz *GBZ80) _SRL_E() uint {
	gbz.srl(&gbz.e)
	return 8
}

func (gbz *GBZ80) _SRL_H() uint {
	gbz.srl(&gbz.h)
	return 8
}

func (gbz *GBZ80) _SRL_L() uint {
	gbz.srl(&gbz.l)
	return 8
}

func (gbz *GBZ80) _SRL_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.srl(&b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SRL_A() uint {
	gbz.srl(&gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_0_B() uint {
	gbz.bit(0, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_0_C() uint {
	gbz.bit(0, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_0_D() uint {
	gbz.bit(0, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_0_E() uint {
	gbz.bit(0, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_0_H() uint {
	gbz.bit(0, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_0_L() uint {
	gbz.bit(0, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_0_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(0, &b)
	return 12
}

func (gbz *GBZ80) _BIT_0_A() uint {
	gbz.bit(0, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_1_B() uint {
	gbz.bit(1, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_1_C() uint {
	gbz.bit(1, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_1_D() uint {
	gbz.bit(1, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_1_E() uint {
	gbz.bit(1, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_1_H() uint {
	gbz.bit(1, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_1_L() uint {
	gbz.bit(1, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_1_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(1, &b)
	return 12
}

func (gbz *GBZ80) _BIT_1_A() uint {
	gbz.bit(1, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_2_B() uint {
	gbz.bit(2, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_2_C() uint {
	gbz.bit(2, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_2_D() uint {
	gbz.bit(2, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_2_E() uint {
	gbz.bit(2, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_2_H() uint {
	gbz.bit(2, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_2_L() uint {
	gbz.bit(2, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_2_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(2, &b)
	return 12
}

func (gbz *GBZ80) _BIT_2_A() uint {
	gbz.bit(2, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_3_B() uint {
	gbz.bit(3, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_3_C() uint {
	gbz.bit(3, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_3_D() uint {
	gbz.bit(3, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_3_E() uint {
	gbz.bit(3, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_3_H() uint {
	gbz.bit(3, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_3_L() uint {
	gbz.bit(3, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_3_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(3, &b)
	return 12
}

func (gbz *GBZ80) _BIT_3_A() uint {
	gbz.bit(3, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_4_B() uint {
	gbz.bit(4, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_4_C() uint {
	gbz.bit(4, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_4_D() uint {
	gbz.bit(4, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_4_E() uint {
	gbz.bit(4, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_4_H() uint {
	gbz.bit(4, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_4_L() uint {
	gbz.bit(4, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_4_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(4, &b)
	return 12
}

func (gbz *GBZ80) _BIT_4_A() uint {
	gbz.bit(4, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_5_B() uint {
	gbz.bit(5, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_5_C() uint {
	gbz.bit(5, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_5_D() uint {
	gbz.bit(5, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_5_E() uint {
	gbz.bit(5, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_5_H() uint {
	gbz.bit(5, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_5_L() uint {
	gbz.bit(5, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_5_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(5, &b)
	return 12
}

func (gbz *GBZ80) _BIT_5_A() uint {
	gbz.bit(5, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_6_B() uint {
	gbz.bit(6, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_6_C() uint {
	gbz.bit(6, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_6_D() uint {
	gbz.bit(6, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_6_E() uint {
	gbz.bit(6, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_6_H() uint {
	gbz.bit(6, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_6_L() uint {
	gbz.bit(6, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_6_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(6, &b)
	return 12
}

func (gbz *GBZ80) _BIT_6_A() uint {
	gbz.bit(6, &gbz.a)
	return 8
}

func (gbz *GBZ80) _BIT_7_B() uint {
	gbz.bit(7, &gbz.b)
	return 8
}

func (gbz *GBZ80) _BIT_7_C() uint {
	gbz.bit(7, &gbz.c)
	return 8
}

func (gbz *GBZ80) _BIT_7_D() uint {
	gbz.bit(7, &gbz.d)
	return 8
}

func (gbz *GBZ80) _BIT_7_E() uint {
	gbz.bit(7, &gbz.e)
	return 8
}

func (gbz *GBZ80) _BIT_7_H() uint {
	gbz.bit(7, &gbz.h)
	return 8
}

func (gbz *GBZ80) _BIT_7_L() uint {
	gbz.bit(7, &gbz.l)
	return 8
}

func (gbz *GBZ80) _BIT_7_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.bit(7, &b)
	return 12
}

func (gbz *GBZ80) _BIT_7_A() uint {
	gbz.bit(7, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_0_B() uint {
	gbz.res(0, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_0_C() uint {
	gbz.res(0, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_0_D() uint {
	gbz.res(0, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_0_E() uint {
	gbz.res(0, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_0_H() uint {
	gbz.res(0, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_0_L() uint {
	gbz.res(0, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_0_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(0, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_0_A() uint {
	gbz.res(0, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_1_B() uint {
	gbz.res(1, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_1_C() uint {
	gbz.res(1, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_1_D() uint {
	gbz.res(1, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_1_E() uint {
	gbz.res(1, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_1_H() uint {
	gbz.res(1, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_1_L() uint {
	gbz.res(1, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_1_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(1, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_1_A() uint {
	gbz.res(1, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_2_B() uint {
	gbz.res(2, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_2_C() uint {
	gbz.res(2, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_2_D() uint {
	gbz.res(2, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_2_E() uint {
	gbz.res(2, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_2_H() uint {
	gbz.res(2, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_2_L() uint {
	gbz.res(2, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_2_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(2, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_2_A() uint {
	gbz.res(2, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_3_B() uint {
	gbz.res(3, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_3_C() uint {
	gbz.res(3, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_3_D() uint {
	gbz.res(3, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_3_E() uint {
	gbz.res(3, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_3_H() uint {
	gbz.res(3, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_3_L() uint {
	gbz.res(3, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_3_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(3, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_3_A() uint {
	gbz.res(3, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_4_B() uint {
	gbz.res(4, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_4_C() uint {
	gbz.res(4, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_4_D() uint {
	gbz.res(4, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_4_E() uint {
	gbz.res(4, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_4_H() uint {
	gbz.res(4, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_4_L() uint {
	gbz.res(4, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_4_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(4, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_4_A() uint {
	gbz.res(4, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_5_B() uint {
	gbz.res(5, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_5_C() uint {
	gbz.res(5, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_5_D() uint {
	gbz.res(5, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_5_E() uint {
	gbz.res(5, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_5_H() uint {
	gbz.res(5, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_5_L() uint {
	gbz.res(5, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_5_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(5, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_5_A() uint {
	gbz.res(5, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_6_B() uint {
	gbz.res(6, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_6_C() uint {
	gbz.res(6, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_6_D() uint {
	gbz.res(6, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_6_E() uint {
	gbz.res(6, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_6_H() uint {
	gbz.res(6, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_6_L() uint {
	gbz.res(6, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_6_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(6, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_6_A() uint {
	gbz.res(6, &gbz.a)
	return 8
}

func (gbz *GBZ80) _RES_7_B() uint {
	gbz.res(7, &gbz.b)
	return 8
}

func (gbz *GBZ80) _RES_7_C() uint {
	gbz.res(7, &gbz.c)
	return 8
}

func (gbz *GBZ80) _RES_7_D() uint {
	gbz.res(7, &gbz.d)
	return 8
}

func (gbz *GBZ80) _RES_7_E() uint {
	gbz.res(7, &gbz.e)
	return 8
}

func (gbz *GBZ80) _RES_7_H() uint {
	gbz.res(7, &gbz.h)
	return 8
}

func (gbz *GBZ80) _RES_7_L() uint {
	gbz.res(7, &gbz.l)
	return 8
}

func (gbz *GBZ80) _RES_7_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.res(7, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _RES_7_A() uint {
	gbz.res(7, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_0_B() uint {
	gbz.set(0, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_0_C() uint {
	gbz.set(0, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_0_D() uint {
	gbz.set(0, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_0_E() uint {
	gbz.set(0, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_0_H() uint {
	gbz.set(0, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_0_L() uint {
	gbz.set(0, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_0_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(0, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_0_A() uint {
	gbz.set(0, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_1_B() uint {
	gbz.set(1, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_1_C() uint {
	gbz.set(1, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_1_D() uint {
	gbz.set(1, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_1_E() uint {
	gbz.set(1, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_1_H() uint {
	gbz.set(1, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_1_L() uint {
	gbz.set(1, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_1_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(1, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_1_A() uint {
	gbz.set(1, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_2_B() uint {
	gbz.set(2, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_2_C() uint {
	gbz.set(2, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_2_D() uint {
	gbz.set(2, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_2_E() uint {
	gbz.set(2, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_2_H() uint {
	gbz.set(2, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_2_L() uint {
	gbz.set(2, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_2_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(2, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_2_A() uint {
	gbz.set(2, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_3_B() uint {
	gbz.set(3, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_3_C() uint {
	gbz.set(3, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_3_D() uint {
	gbz.set(3, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_3_E() uint {
	gbz.set(3, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_3_H() uint {
	gbz.set(3, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_3_L() uint {
	gbz.set(3, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_3_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(3, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_3_A() uint {
	gbz.set(3, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_4_B() uint {
	gbz.set(4, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_4_C() uint {
	gbz.set(4, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_4_D() uint {
	gbz.set(4, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_4_E() uint {
	gbz.set(4, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_4_H() uint {
	gbz.set(4, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_4_L() uint {
	gbz.set(4, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_4_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(4, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_4_A() uint {
	gbz.set(4, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_5_B() uint {
	gbz.set(5, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_5_C() uint {
	gbz.set(5, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_5_D() uint {
	gbz.set(5, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_5_E() uint {
	gbz.set(5, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_5_H() uint {
	gbz.set(5, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_5_L() uint {
	gbz.set(5, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_5_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(5, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_5_A() uint {
	gbz.set(5, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_6_B() uint {
	gbz.set(6, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_6_C() uint {
	gbz.set(6, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_6_D() uint {
	gbz.set(6, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_6_E() uint {
	gbz.set(6, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_6_H() uint {
	gbz.set(6, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_6_L() uint {
	gbz.set(6, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_6_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(6, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_6_A() uint {
	gbz.set(6, &gbz.a)
	return 8
}

func (gbz *GBZ80) _SET_7_B() uint {
	gbz.set(7, &gbz.b)
	return 8
}

func (gbz *GBZ80) _SET_7_C() uint {
	gbz.set(7, &gbz.c)
	return 8
}

func (gbz *GBZ80) _SET_7_D() uint {
	gbz.set(7, &gbz.d)
	return 8
}

func (gbz *GBZ80) _SET_7_E() uint {
	gbz.set(7, &gbz.e)
	return 8
}

func (gbz *GBZ80) _SET_7_H() uint {
	gbz.set(7, &gbz.h)
	return 8
}

func (gbz *GBZ80) _SET_7_L() uint {
	gbz.set(7, &gbz.l)
	return 8
}

func (gbz *GBZ80) _SET_7_HL() uint {
	hl := (uint16(gbz.h) << 8) | uint16(gbz.l)
	b := gbz.mem.Read(hl)
	gbz.set(7, &b)
	gbz.mem.Write(hl, b)
	return 16
}

func (gbz *GBZ80) _SET_7_A() uint {
	gbz.set(7, &gbz.a)
	return 8
}
