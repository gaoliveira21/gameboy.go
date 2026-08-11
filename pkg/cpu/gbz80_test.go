package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Memory [0x10000]byte

func (m *Memory) Read(addr uint16) byte {
	return m[addr]
}

func (m *Memory) Write(addr uint16, b byte) {
	m[addr] = b
}

func createGBZ() *GBZ80 {
	return NewGBZ80(&Memory{})
}

func resetGBZ(gbz *GBZ80) *GBZ80 {
	gbz.a = 0x01
	gbz.b = 0x00
	gbz.c = 0x13
	gbz.d = 0x00
	gbz.e = 0xd8
	gbz.h = 0x01
	gbz.l = 0x4d
	gbz.flags = &Flags{
		value: 0xb0,
	}
	gbz.pc = 0x100
	gbz.sp = 0xFFFE
	gbz.mem = &Memory{}
	gbz.cycles = 0
	gbz.interruptEnabled = false

	return gbz
}

func createGBZWithOpcode(opcode uint8) *GBZ80 {
	gbz := createGBZ()
	gbz.mem.Write(gbz.pc, byte(opcode))

	return gbz
}

func assertCycles(t *testing.T, gbz *GBZ80, expected uint) {
	t.Helper()
	assert.Equal(t, expected, gbz.cycles, "Expected cycles %d, but got %d", expected, gbz.cycles)
}

func assertPC(t *testing.T, gbz *GBZ80, expected uint16) {
	t.Helper()
	assert.Equal(t, expected, gbz.pc, "Expected PC %04x, but got %04x", expected, gbz.pc)
}

func assertMemory(t *testing.T, gbz *GBZ80, addr uint16, expected byte) {
	t.Helper()
	got := gbz.mem.Read(addr)
	assert.Equal(t, got, expected, "Expected value %02x to be written in memory at %04x, but got %02x", expected, addr, got)
}

func assertRegister(t *testing.T, expected byte, got byte, reg string) {
	t.Helper()
	assert.Equal(t, expected, got, "Expected register %s %02x, but got %02x", reg, expected, got)
}

func assertSP(t *testing.T, expected uint16, got uint16) {
	t.Helper()
	assert.Equal(t, expected, got, "Expected SP %04x, but got %04x", expected, got)
}

func TestNewGBZ80(t *testing.T) {
	gbz := createGBZ()

	assert.EqualValues(t, 0x01, gbz.a, "NewGBZ80 dit not start register A correctly")
	assert.EqualValues(t, 0x00, gbz.b, "NewGBZ80 dit not start register B correctly")
	assert.EqualValues(t, 0x13, gbz.c, "NewGBZ80 dit not start register C correctly")
	assert.EqualValues(t, 0x00, gbz.d, "NewGBZ80 dit not start register D correctly")
	assert.EqualValues(t, 0xD8, gbz.e, "NewGBZ80 dit not start register E correctly")
	assert.EqualValues(t, 0x01, gbz.h, "NewGBZ80 dit not start register H correctly")
	assert.EqualValues(t, 0x4D, gbz.l, "NewGBZ80 dit not start register L correctly")
	assert.EqualValues(t, 0xB0, gbz.flags.value, "NewGBZ80 dit not start flags correctly")
	assert.EqualValues(t, 0x100, gbz.pc, "NewGBZ80 dit not start PC correctly")
	assert.EqualValues(t, 0xFFFE, gbz.sp, "NewGBZ80 dit not start SP correctly")
	assert.EqualValues(t, 0xCF, gbz.mem.Read(0xFF00), "NewGBZ80 dit not boot memory correctly at 0xFF00")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF01), "NewGBZ80 dit not boot memory correctly at 0xFF01")
	assert.EqualValues(t, 0x7E, gbz.mem.Read(0xFF02), "NewGBZ80 dit not boot memory correctly at 0xFF02")
	assert.EqualValues(t, 0xAB, gbz.mem.Read(0xFF04), "NewGBZ80 dit not boot memory correctly at 0xFF04")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF05), "NewGBZ80 dit not boot memory correctly at 0xFF05")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF06), "NewGBZ80 dit not boot memory correctly at 0xFF06")
	assert.EqualValues(t, 0xF8, gbz.mem.Read(0xFF07), "NewGBZ80 dit not boot memory correctly at 0xFF07")
	assert.EqualValues(t, 0xE1, gbz.mem.Read(0xFF0F), "NewGBZ80 dit not boot memory correctly at 0xFF0F")
	assert.EqualValues(t, 0x80, gbz.mem.Read(0xFF10), "NewGBZ80 dit not boot memory correctly at 0xFF10")
	assert.EqualValues(t, 0xBF, gbz.mem.Read(0xFF11), "NewGBZ80 dit not boot memory correctly at 0xFF11")
	assert.EqualValues(t, 0xF3, gbz.mem.Read(0xFF12), "NewGBZ80 dit not boot memory correctly at 0xFF12")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF13), "NewGBZ80 dit not boot memory correctly at 0xFF13")
	assert.EqualValues(t, 0xBF, gbz.mem.Read(0xFF14), "NewGBZ80 dit not boot memory correctly at 0xFF14")
	assert.EqualValues(t, 0x3F, gbz.mem.Read(0xFF16), "NewGBZ80 dit not boot memory correctly at 0xFF16")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF17), "NewGBZ80 dit not boot memory correctly at 0xFF17")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF18), "NewGBZ80 dit not boot memory correctly at 0xFF18")
	assert.EqualValues(t, 0xBF, gbz.mem.Read(0xFF19), "NewGBZ80 dit not boot memory correctly at 0xFF19")
	assert.EqualValues(t, 0x7F, gbz.mem.Read(0xFF1A), "NewGBZ80 dit not boot memory correctly at 0xFF1A")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF1B), "NewGBZ80 dit not boot memory correctly at 0xFF1B")
	assert.EqualValues(t, 0x9F, gbz.mem.Read(0xFF1C), "NewGBZ80 dit not boot memory correctly at 0xFF1C")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF1D), "NewGBZ80 dit not boot memory correctly at 0xFF1D")
	assert.EqualValues(t, 0xBF, gbz.mem.Read(0xFF1E), "NewGBZ80 dit not boot memory correctly at 0xFF1E")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF20), "NewGBZ80 dit not boot memory correctly at 0xFF20")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF21), "NewGBZ80 dit not boot memory correctly at 0xFF21")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF22), "NewGBZ80 dit not boot memory correctly at 0xFF22")
	assert.EqualValues(t, 0xBF, gbz.mem.Read(0xFF23), "NewGBZ80 dit not boot memory correctly at 0xFF23")
	assert.EqualValues(t, 0x77, gbz.mem.Read(0xFF24), "NewGBZ80 dit not boot memory correctly at 0xFF24")
	assert.EqualValues(t, 0xF3, gbz.mem.Read(0xFF25), "NewGBZ80 dit not boot memory correctly at 0xFF25")
	assert.EqualValues(t, 0xF1, gbz.mem.Read(0xFF26), "NewGBZ80 dit not boot memory correctly at 0xFF26")
	assert.EqualValues(t, 0x91, gbz.mem.Read(0xFF40), "NewGBZ80 dit not boot memory correctly at 0xFF40")
	assert.EqualValues(t, 0x85, gbz.mem.Read(0xFF41), "NewGBZ80 dit not boot memory correctly at 0xFF41")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF42), "NewGBZ80 dit not boot memory correctly at 0xFF42")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF43), "NewGBZ80 dit not boot memory correctly at 0xFF43")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF44), "NewGBZ80 dit not boot memory correctly at 0xFF44")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF45), "NewGBZ80 dit not boot memory correctly at 0xFF45")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF46), "NewGBZ80 dit not boot memory correctly at 0xFF46")
	assert.EqualValues(t, 0xFC, gbz.mem.Read(0xFF47), "NewGBZ80 dit not boot memory correctly at 0xFF47")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF48), "NewGBZ80 dit not boot memory correctly at 0xFF48")
	assert.EqualValues(t, 0xFF, gbz.mem.Read(0xFF49), "NewGBZ80 dit not boot memory correctly at 0xFF49")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF4A), "NewGBZ80 dit not boot memory correctly at 0xFF4A")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFF4B), "NewGBZ80 dit not boot memory correctly at 0xFF4B")
	assert.EqualValues(t, 0x00, gbz.mem.Read(0xFFFF), "NewGBZ80 dit not boot memory correctly at 0xFFFF")
}

func Test_NOP(t *testing.T) {
	gbz := createGBZWithOpcode(0x00)
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
}

func TestLdR8N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation       string
		Opcode          uint8
		Register        *uint8
		RegisterInitial string
	}{
		{"_LD_B_n8", 0x06, &gbz.b, "B"},
		{"_LD_C_n8", 0x0E, &gbz.c, "C"},
		{"_LD_D_n8", 0x16, &gbz.d, "D"},
		{"_LD_E_n8", 0x1E, &gbz.e, "E"},
		{"_LD_H_n8", 0x26, &gbz.h, "H"},
		{"_LD_L_n8", 0x2E, &gbz.l, "L"},
		{"_LD_A_n8", 0x3E, &gbz.a, "A"},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			pc := gbz.pc
			expected := byte(0x99)
			gbz.mem.Write(pc+1, expected)

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, expected, *c.Register, c.RegisterInitial)
			resetGBZ(gbz)
		})
	}
}

func TestLdRpn16(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation        string
		Opcode           uint8
		Register1        *uint8
		Register1Initial string
		Register2        *uint8
		Register2Initial string
	}{
		{"_LD_BC_n16", 0x01, &gbz.b, "B", &gbz.c, "C"},
		{"_LD_DE_n16", 0x11, &gbz.d, "D", &gbz.e, "E"},
		{"_LD_HL_n16", 0x21, &gbz.h, "H", &gbz.l, "L"},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			pc := gbz.pc
			lb := byte(0x20)
			hb := byte(0x80)

			gbz.mem.Write(pc+1, lb)
			gbz.mem.Write(pc+2, hb)

			gbz.Run()

			assertCycles(t, gbz, 12)
			assertPC(t, gbz, pc+3)
			assertRegister(t, hb, *c.Register1, c.Register1Initial)
			assertRegister(t, lb, *c.Register2, c.Register2Initial)

			resetGBZ(gbz)
		})
	}
}

func TestLdRpR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation      string
		Opcode         uint8
		Register1      *uint8
		Register2      *uint8
		RegisterTarget *uint8
		Inc            int
	}{
		{"_LD_BC_A", 0x02, &gbz.b, &gbz.c, &gbz.a, 0},
		{"_LD_DE_A", 0x12, &gbz.d, &gbz.e, &gbz.a, 0},
		{"_LD_HLI_A", 0x22, &gbz.h, &gbz.l, &gbz.a, 1},
		{"_LD_HLD_A", 0x32, &gbz.h, &gbz.l, &gbz.a, -1},
		{"_LD_HL_B", 0x70, &gbz.h, &gbz.l, &gbz.b, 0},
		{"_LD_HL_C", 0x71, &gbz.h, &gbz.l, &gbz.c, 0},
		{"_LD_HL_D", 0x72, &gbz.h, &gbz.l, &gbz.d, 0},
		{"_LD_HL_E", 0x73, &gbz.h, &gbz.l, &gbz.e, 0},
		{"_LD_HL_H", 0x74, &gbz.h, &gbz.l, &gbz.h, 0},
		{"_LD_HL_L", 0x75, &gbz.h, &gbz.l, &gbz.l, 0},
		{"_LD_HL_A", 0x77, &gbz.h, &gbz.l, &gbz.a, 0},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.Register1 = 0x20
			*c.Register2 = 0x80
			*c.RegisterTarget = 0x99
			pc := gbz.pc
			addr := int((uint16(*c.Register1)<<8)|uint16(*c.Register2)) + c.Inc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertMemory(t, gbz, uint16(addr), *c.RegisterTarget)

			resetGBZ(gbz)
		})
	}
}

func TestLdR8Rp(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation             string
		Opcode                uint8
		Register1             *uint8
		Register2             *uint8
		RegisterTarget        *uint8
		RegisterTargetInitial string
		Inc                   int
	}{
		{"_LD_A_BC", 0x0A, &gbz.b, &gbz.c, &gbz.a, "A", 0},
		{"_LD_A_DE", 0x1A, &gbz.d, &gbz.e, &gbz.a, "A", 0},
		{"_LD_A_HLI", 0x2A, &gbz.h, &gbz.l, &gbz.a, "A", 1},
		{"_LD_A_HLD", 0x3A, &gbz.h, &gbz.l, &gbz.a, "A", -1},
		{"_LD_B_HL", 0x46, &gbz.h, &gbz.l, &gbz.b, "B", 0},
		{"_LD_C_HL", 0x4E, &gbz.h, &gbz.l, &gbz.c, "C", 0},
		{"_LD_D_HL", 0x56, &gbz.h, &gbz.l, &gbz.d, "D", 0},
		{"_LD_E_HL", 0x5E, &gbz.h, &gbz.l, &gbz.e, "E", 0},
		{"_LD_H_HL", 0x66, &gbz.h, &gbz.l, &gbz.h, "H", 0},
		{"_LD_L_HL", 0x6E, &gbz.h, &gbz.l, &gbz.l, "L", 0},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.Register1 = 0x20
			*c.Register2 = 0x80
			pc := gbz.pc
			addr := int((uint16(*c.Register1)<<8)|uint16(*c.Register2)) + c.Inc
			expected := byte(0x99)
			gbz.mem.Write(uint16(addr), expected)

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertMemory(t, gbz, uint16(addr), *c.RegisterTarget)
			assertRegister(t, expected, *c.RegisterTarget, c.RegisterTargetInitial)

			resetGBZ(gbz)
		})
	}
}

func Test_LD_a16_SP(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0x8020
	pc := gbz.pc

	targetAddr := uint16(0xC000)
	gbz.mem.Write(pc, 0x08)
	gbz.mem.Write(pc+1, byte(targetAddr))
	gbz.mem.Write(pc+2, byte(targetAddr>>8))

	gbz.Run()

	assertCycles(t, gbz, 20)
	assertPC(t, gbz, pc+3)
	assertMemory(t, gbz, targetAddr, byte(gbz.sp&0xFF))
	assertMemory(t, gbz, targetAddr+1, byte(gbz.sp>>8))
}

func Test_LD_SP_n16(t *testing.T) {
	gbz := createGBZ()
	pc := gbz.pc

	expected := uint16(0xC000)
	gbz.mem.Write(pc, 0x31)
	gbz.mem.Write(pc+1, byte(expected))
	gbz.mem.Write(pc+2, byte(expected>>8))

	gbz.Run()
	assertCycles(t, gbz, 12)
	assertPC(t, gbz, pc+3)
	assertSP(t, expected, gbz.sp)
}

func Test_LD_HL_n8(t *testing.T) {
	gbz := createGBZ()
	gbz.h = 0xC0
	gbz.l = 0x00
	pc := gbz.pc

	opcode := byte(0x36)
	expected := byte(0x99)
	gbz.mem.Write(pc, opcode)
	gbz.mem.Write(pc+1, expected)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, pc+2)
	assertMemory(t, gbz, 0xC000, expected)
}

func Test_LdR8R8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation          string
		Opcode             uint8
		RegisterDst        *uint8
		RegisterDstInitial string
		RegisterSrc        *uint8
	}{
		{"_LD_B_B", 0x40, &gbz.b, "B", &gbz.b},
		{"_LD_B_C", 0x41, &gbz.b, "B", &gbz.c},
		{"_LD_B_D", 0x42, &gbz.b, "B", &gbz.d},
		{"_LD_B_E", 0x43, &gbz.b, "B", &gbz.e},
		{"_LD_B_H", 0x44, &gbz.b, "B", &gbz.h},
		{"_LD_B_L", 0x45, &gbz.b, "B", &gbz.l},
		{"_LD_B_A", 0x47, &gbz.b, "B", &gbz.a},
		{"_LD_C_B", 0x48, &gbz.c, "C", &gbz.b},
		{"_LD_C_C", 0x49, &gbz.c, "C", &gbz.c},
		{"_LD_C_D", 0x4A, &gbz.c, "C", &gbz.d},
		{"_LD_C_E", 0x4B, &gbz.c, "C", &gbz.e},
		{"_LD_C_H", 0x4C, &gbz.c, "C", &gbz.h},
		{"_LD_C_L", 0x4D, &gbz.c, "C", &gbz.l},
		{"_LD_C_A", 0x4F, &gbz.c, "C", &gbz.a},
		{"_LD_D_B", 0x50, &gbz.d, "D", &gbz.b},
		{"_LD_D_C", 0x51, &gbz.d, "D", &gbz.c},
		{"_LD_D_D", 0x52, &gbz.d, "D", &gbz.d},
		{"_LD_D_E", 0x53, &gbz.d, "D", &gbz.e},
		{"_LD_D_H", 0x54, &gbz.d, "D", &gbz.h},
		{"_LD_D_L", 0x55, &gbz.d, "D", &gbz.l},
		{"_LD_D_A", 0x57, &gbz.d, "D", &gbz.a},
		{"_LD_E_B", 0x58, &gbz.e, "E", &gbz.b},
		{"_LD_E_C", 0x59, &gbz.e, "E", &gbz.c},
		{"_LD_E_D", 0x5A, &gbz.e, "E", &gbz.d},
		{"_LD_E_E", 0x5B, &gbz.e, "E", &gbz.e},
		{"_LD_E_H", 0x5C, &gbz.e, "E", &gbz.h},
		{"_LD_E_L", 0x5D, &gbz.e, "E", &gbz.l},
		{"_LD_E_A", 0x5F, &gbz.e, "E", &gbz.a},
		{"_LD_H_B", 0x60, &gbz.h, "H", &gbz.b},
		{"_LD_H_C", 0x61, &gbz.h, "H", &gbz.c},
		{"_LD_H_D", 0x62, &gbz.h, "H", &gbz.d},
		{"_LD_H_E", 0x63, &gbz.h, "H", &gbz.e},
		{"_LD_H_H", 0x64, &gbz.h, "H", &gbz.h},
		{"_LD_H_L", 0x65, &gbz.h, "H", &gbz.l},
		{"_LD_H_A", 0x67, &gbz.h, "H", &gbz.a},
		{"_LD_L_B", 0x68, &gbz.l, "L", &gbz.b},
		{"_LD_L_C", 0x69, &gbz.l, "L", &gbz.c},
		{"_LD_L_D", 0x6A, &gbz.l, "L", &gbz.d},
		{"_LD_L_E", 0x6B, &gbz.l, "L", &gbz.e},
		{"_LD_L_H", 0x6C, &gbz.l, "L", &gbz.h},
		{"_LD_L_L", 0x6D, &gbz.l, "L", &gbz.l},
		{"_LD_L_A", 0x6F, &gbz.l, "L", &gbz.a},
		{"_LD_A_B", 0x78, &gbz.a, "A", &gbz.b},
		{"_LD_A_C", 0x79, &gbz.a, "A", &gbz.c},
		{"_LD_A_D", 0x7A, &gbz.a, "A", &gbz.d},
		{"_LD_A_E", 0x7B, &gbz.a, "A", &gbz.e},
		{"_LD_A_H", 0x7C, &gbz.a, "A", &gbz.h},
		{"_LD_A_L", 0x7D, &gbz.a, "A", &gbz.l},
		{"_LD_A_A", 0x7F, &gbz.a, "A", &gbz.a},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.RegisterDst = 0x0
			*c.RegisterSrc = 0x99
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, *c.RegisterSrc, *c.RegisterDst, c.RegisterDstInitial)

			resetGBZ(gbz)
		})
	}
}

func Test_DI(t *testing.T) {
	gbz := createGBZWithOpcode(0xF3)
	gbz.interruptEnabled = true
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assert.Equal(t, false, gbz.interruptEnabled, "DI should disable interrupts")
}

func Test_EI(t *testing.T) {
	gbz := createGBZWithOpcode(0xFB)
	gbz.interruptEnabled = false
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assert.Equal(t, true, gbz.interruptEnabled, "EI should enable interrupts")
}

func Test_INCR16(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation        string
		Opcode           uint8
		Register1        *uint8
		Register1Initial string
		Register2        *uint8
		Register2Initial string
	}{
		{"_INC_BC", 0x03, &gbz.b, "B", &gbz.c, "C"},
		{"_INC_DE", 0x13, &gbz.d, "D", &gbz.e, "E"},
		{"_INC_HL", 0x23, &gbz.h, "H", &gbz.l, "L"},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.Register1 = 0x20
			*c.Register2 = 0x80
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, 0x20, *c.Register1, c.Register1Initial)
			assertRegister(t, 0x81, *c.Register2, c.Register2Initial)

			resetGBZ(gbz)
		})
	}
}

func Test_INC_SP(t *testing.T) {
	gbz := createGBZWithOpcode(0x33)
	pc := gbz.pc
	sp := gbz.sp

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+1)
	assertSP(t, sp+1, gbz.sp)
}
