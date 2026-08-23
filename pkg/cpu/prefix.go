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
		0x40: {g._TODO, "BIT 0 B", 2},
		0x41: {g._TODO, "BIT 0 C", 2},
		0x42: {g._TODO, "BIT 0 D", 2},
		0x43: {g._TODO, "BIT 0 E", 2},
		0x44: {g._TODO, "BIT 0 H", 2},
		0x45: {g._TODO, "BIT 0 L", 2},
		0x46: {g._TODO, "BIT 0 HL", 2},
		0x47: {g._TODO, "BIT 0 A", 2},
		0x48: {g._TODO, "BIT 1 B", 2},
		0x49: {g._TODO, "BIT 1 C", 2},
		0x4A: {g._TODO, "BIT 1 D", 2},
		0x4B: {g._TODO, "BIT 1 E", 2},
		0x4C: {g._TODO, "BIT 1 H", 2},
		0x4D: {g._TODO, "BIT 1 L", 2},
		0x4E: {g._TODO, "BIT 1 HL", 2},
		0x4F: {g._TODO, "BIT 1 A", 2},
		0x50: {g._TODO, "BIT 2 B", 2},
		0x51: {g._TODO, "BIT 2 C", 2},
		0x52: {g._TODO, "BIT 2 D", 2},
		0x53: {g._TODO, "BIT 2 E", 2},
		0x54: {g._TODO, "BIT 2 H", 2},
		0x55: {g._TODO, "BIT 2 L", 2},
		0x56: {g._TODO, "BIT 2 HL", 2},
		0x57: {g._TODO, "BIT 2 A", 2},
		0x58: {g._TODO, "BIT 3 B", 2},
		0x59: {g._TODO, "BIT 3 C", 2},
		0x5A: {g._TODO, "BIT 3 D", 2},
		0x5B: {g._TODO, "BIT 3 E", 2},
		0x5C: {g._TODO, "BIT 3 H", 2},
		0x5D: {g._TODO, "BIT 3 L", 2},
		0x5E: {g._TODO, "BIT 3 HL", 2},
		0x5F: {g._TODO, "BIT 3 A", 2},
		0x60: {g._TODO, "BIT 4 B", 2},
		0x61: {g._TODO, "BIT 4 C", 2},
		0x62: {g._TODO, "BIT 4 D", 2},
		0x63: {g._TODO, "BIT 4 E", 2},
		0x64: {g._TODO, "BIT 4 H", 2},
		0x65: {g._TODO, "BIT 4 L", 2},
		0x66: {g._TODO, "BIT 4 HL", 2},
		0x67: {g._TODO, "BIT 4 A", 2},
		0x68: {g._TODO, "BIT 5 B", 2},
		0x69: {g._TODO, "BIT 5 C", 2},
		0x6A: {g._TODO, "BIT 5 D", 2},
		0x6B: {g._TODO, "BIT 5 E", 2},
		0x6C: {g._TODO, "BIT 5 H", 2},
		0x6D: {g._TODO, "BIT 5 L", 2},
		0x6E: {g._TODO, "BIT 5 HL", 2},
		0x6F: {g._TODO, "BIT 5 A", 2},
		0x70: {g._TODO, "BIT 6 B", 2},
		0x71: {g._TODO, "BIT 6 C", 2},
		0x72: {g._TODO, "BIT 6 D", 2},
		0x73: {g._TODO, "BIT 6 E", 2},
		0x74: {g._TODO, "BIT 6 H", 2},
		0x75: {g._TODO, "BIT 6 L", 2},
		0x76: {g._TODO, "BIT 6 HL", 2},
		0x77: {g._TODO, "BIT 6 A", 2},
		0x78: {g._TODO, "BIT 7 B", 2},
		0x79: {g._TODO, "BIT 7 C", 2},
		0x7A: {g._TODO, "BIT 7 D", 2},
		0x7B: {g._TODO, "BIT 7 E", 2},
		0x7C: {g._TODO, "BIT 7 H", 2},
		0x7D: {g._TODO, "BIT 7 L", 2},
		0x7E: {g._TODO, "BIT 7 HL", 2},
		0x7F: {g._TODO, "BIT 7 A", 2},
		0x80: {g._TODO, "RES 0 B", 2},
		0x81: {g._TODO, "RES 0 C", 2},
		0x82: {g._TODO, "RES 0 D", 2},
		0x83: {g._TODO, "RES 0 E", 2},
		0x84: {g._TODO, "RES 0 H", 2},
		0x85: {g._TODO, "RES 0 L", 2},
		0x86: {g._TODO, "RES 0 HL", 2},
		0x87: {g._TODO, "RES 0 A", 2},
		0x88: {g._TODO, "RES 1 B", 2},
		0x89: {g._TODO, "RES 1 C", 2},
		0x8A: {g._TODO, "RES 1 D", 2},
		0x8B: {g._TODO, "RES 1 E", 2},
		0x8C: {g._TODO, "RES 1 H", 2},
		0x8D: {g._TODO, "RES 1 L", 2},
		0x8E: {g._TODO, "RES 1 HL", 2},
		0x8F: {g._TODO, "RES 1 A", 2},
		0x90: {g._TODO, "RES 2 B", 2},
		0x91: {g._TODO, "RES 2 C", 2},
		0x92: {g._TODO, "RES 2 D", 2},
		0x93: {g._TODO, "RES 2 E", 2},
		0x94: {g._TODO, "RES 2 H", 2},
		0x95: {g._TODO, "RES 2 L", 2},
		0x96: {g._TODO, "RES 2 HL", 2},
		0x97: {g._TODO, "RES 2 A", 2},
		0x98: {g._TODO, "RES 3 B", 2},
		0x99: {g._TODO, "RES 3 C", 2},
		0x9A: {g._TODO, "RES 3 D", 2},
		0x9B: {g._TODO, "RES 3 E", 2},
		0x9C: {g._TODO, "RES 3 H", 2},
		0x9D: {g._TODO, "RES 3 L", 2},
		0x9E: {g._TODO, "RES 3 HL", 2},
		0x9F: {g._TODO, "RES 3 A", 2},
		0xA0: {g._TODO, "RES 4 B", 2},
		0xA1: {g._TODO, "RES 4 C", 2},
		0xA2: {g._TODO, "RES 4 D", 2},
		0xA3: {g._TODO, "RES 4 E", 2},
		0xA4: {g._TODO, "RES 4 H", 2},
		0xA5: {g._TODO, "RES 4 L", 2},
		0xA6: {g._TODO, "RES 4 HL", 2},
		0xA7: {g._TODO, "RES 4 A", 2},
		0xA8: {g._TODO, "RES 5 B", 2},
		0xA9: {g._TODO, "RES 5 C", 2},
		0xAA: {g._TODO, "RES 5 D", 2},
		0xAB: {g._TODO, "RES 5 E", 2},
		0xAC: {g._TODO, "RES 5 H", 2},
		0xAD: {g._TODO, "RES 5 L", 2},
		0xAE: {g._TODO, "RES 5 HL", 2},
		0xAF: {g._TODO, "RES 5 A", 2},
		0xB0: {g._TODO, "RES 6 B", 2},
		0xB1: {g._TODO, "RES 6 C", 2},
		0xB2: {g._TODO, "RES 6 D", 2},
		0xB3: {g._TODO, "RES 6 E", 2},
		0xB4: {g._TODO, "RES 6 H", 2},
		0xB5: {g._TODO, "RES 6 L", 2},
		0xB6: {g._TODO, "RES 6 HL", 2},
		0xB7: {g._TODO, "RES 6 A", 2},
		0xB8: {g._TODO, "RES 7 B", 2},
		0xB9: {g._TODO, "RES 7 C", 2},
		0xBA: {g._TODO, "RES 7 D", 2},
		0xBB: {g._TODO, "RES 7 E", 2},
		0xBC: {g._TODO, "RES 7 H", 2},
		0xBD: {g._TODO, "RES 7 L", 2},
		0xBE: {g._TODO, "RES 7 HL", 2},
		0xBF: {g._TODO, "RES 7 A", 2},
		0xC0: {g._TODO, "SET 0 B", 2},
		0xC1: {g._TODO, "SET 0 C", 2},
		0xC2: {g._TODO, "SET 0 D", 2},
		0xC3: {g._TODO, "SET 0 E", 2},
		0xC4: {g._TODO, "SET 0 H", 2},
		0xC5: {g._TODO, "SET 0 L", 2},
		0xC6: {g._TODO, "SET 0 HL", 2},
		0xC7: {g._TODO, "SET 0 A", 2},
		0xC8: {g._TODO, "SET 1 B", 2},
		0xC9: {g._TODO, "SET 1 C", 2},
		0xCA: {g._TODO, "SET 1 D", 2},
		0xCB: {g._TODO, "SET 1 E", 2},
		0xCC: {g._TODO, "SET 1 H", 2},
		0xCD: {g._TODO, "SET 1 L", 2},
		0xCE: {g._TODO, "SET 1 HL", 2},
		0xCF: {g._TODO, "SET 1 A", 2},
		0xD0: {g._TODO, "SET 2 B", 2},
		0xD1: {g._TODO, "SET 2 C", 2},
		0xD2: {g._TODO, "SET 2 D", 2},
		0xD3: {g._TODO, "SET 2 E", 2},
		0xD4: {g._TODO, "SET 2 H", 2},
		0xD5: {g._TODO, "SET 2 L", 2},
		0xD6: {g._TODO, "SET 2 HL", 2},
		0xD7: {g._TODO, "SET 2 A", 2},
		0xD8: {g._TODO, "SET 3 B", 2},
		0xD9: {g._TODO, "SET 3 C", 2},
		0xDA: {g._TODO, "SET 3 D", 2},
		0xDB: {g._TODO, "SET 3 E", 2},
		0xDC: {g._TODO, "SET 3 H", 2},
		0xDD: {g._TODO, "SET 3 L", 2},
		0xDE: {g._TODO, "SET 3 HL", 2},
		0xDF: {g._TODO, "SET 3 A", 2},
		0xE0: {g._TODO, "SET 4 B", 2},
		0xE1: {g._TODO, "SET 4 C", 2},
		0xE2: {g._TODO, "SET 4 D", 2},
		0xE3: {g._TODO, "SET 4 E", 2},
		0xE4: {g._TODO, "SET 4 H", 2},
		0xE5: {g._TODO, "SET 4 L", 2},
		0xE6: {g._TODO, "SET 4 HL", 2},
		0xE7: {g._TODO, "SET 4 A", 2},
		0xE8: {g._TODO, "SET 5 B", 2},
		0xE9: {g._TODO, "SET 5 C", 2},
		0xEA: {g._TODO, "SET 5 D", 2},
		0xEB: {g._TODO, "SET 5 E", 2},
		0xEC: {g._TODO, "SET 5 H", 2},
		0xED: {g._TODO, "SET 5 L", 2},
		0xEE: {g._TODO, "SET 5 HL", 2},
		0xEF: {g._TODO, "SET 5 A", 2},
		0xF0: {g._TODO, "SET 6 B", 2},
		0xF1: {g._TODO, "SET 6 C", 2},
		0xF2: {g._TODO, "SET 6 D", 2},
		0xF3: {g._TODO, "SET 6 E", 2},
		0xF4: {g._TODO, "SET 6 H", 2},
		0xF5: {g._TODO, "SET 6 L", 2},
		0xF6: {g._TODO, "SET 6 HL", 2},
		0xF7: {g._TODO, "SET 6 A", 2},
		0xF8: {g._TODO, "SET 7 B", 2},
		0xF9: {g._TODO, "SET 7 C", 2},
		0xFA: {g._TODO, "SET 7 D", 2},
		0xFB: {g._TODO, "SET 7 E", 2},
		0xFC: {g._TODO, "SET 7 H", 2},
		0xFD: {g._TODO, "SET 7 L", 2},
		0xFE: {g._TODO, "SET 7 HL", 2},
		0xFF: {g._TODO, "SET 7 A", 2},
	}
}

func (gbz *GBZ80) _TODO() uint {
	panic("Instruction not implemented")
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
