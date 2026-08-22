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
	gbz.InterruptEnabled = false

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
	gbz.InterruptEnabled = true
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assert.Equal(t, false, gbz.InterruptEnabled, "DI should disable interrupts")
}

func Test_EI(t *testing.T) {
	gbz := createGBZWithOpcode(0xFB)
	gbz.InterruptEnabled = false
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assert.Equal(t, true, gbz.InterruptEnabled, "EI should enable interrupts")
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

func Test_INCR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation       string
		Opcode          uint8
		Register        *uint8
		RegisterInitial string
		InitialValue    byte
		WantValue       byte
		WantZero        bool
		WantSub         bool
		WantHalfCarry   bool
	}{
		{"_INC_B", 0x04, &gbz.b, "B", 0x00, 0x01, false, false, false},
		{"_INC_C", 0x0C, &gbz.c, "C", 0x00, 0x01, false, false, false},
		{"_INC_D", 0x14, &gbz.d, "D", 0x00, 0x01, false, false, false},
		{"_INC_E", 0x1C, &gbz.e, "E", 0x00, 0x01, false, false, false},
		{"_INC_H", 0x24, &gbz.h, "H", 0x00, 0x01, false, false, false},
		{"_INC_L", 0x2C, &gbz.l, "L", 0x00, 0x01, false, false, false},
		{"_INC_A", 0x3C, &gbz.a, "A", 0x00, 0x01, false, false, false},
		{"_INC_A_Zero", 0x3C, &gbz.a, "A", 0xFF, 0x00, true, false, true},
		{"_INC_A_HalfCarry_0x0F", 0x3C, &gbz.a, "A", 0x0F, 0x10, false, false, true},
		{"_INC_A_HalfCarry_0x1F", 0x3C, &gbz.a, "A", 0x1F, 0x20, false, false, true},
		{"_INC_A_HalfCarry_0x7F", 0x3C, &gbz.a, "A", 0x7F, 0x80, false, false, true},
		{"_INC_A_NoFlags", 0x3C, &gbz.a, "A", 0x01, 0x02, false, false, false},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.Register = c.InitialValue
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.WantValue, *c.Register, c.RegisterInitial)
			assert.Equal(t, c.WantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, c.WantSub, gbz.flags.Get(Sub), "Sub flag mismatch")
			assert.Equal(t, c.WantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")

			resetGBZ(gbz)
		})
	}
}

func Test_INC_HLX(t *testing.T) {
	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		wantValue     byte
		wantZero      bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0xC000, 0x01, 0x02, false, false},
		{"Zero", 0xC000, 0xFF, 0x00, true, true},
		{"HalfCarry_0x0F", 0xC000, 0x0F, 0x10, false, true},
		{"HalfCarry_0x1F", 0xC000, 0x1F, 0x20, false, true},
		{"HalfCarry_0xFF", 0xC000, 0xFF, 0x00, true, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x34)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 12)
			assertPC(t, gbz, pc+1)
			assert.Equal(t, c.wantValue, gbz.mem.Read(c.hl), "Memory value mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
		})
	}
}

func Test_DECR16(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation        string
		Opcode           uint8
		Register1        *uint8
		Register1Initial string
		Register2        *uint8
		Register2Initial string
	}{
		{"_DEC_BC", 0x0B, &gbz.b, "B", &gbz.c, "C"},
		{"_DEC_DE", 0x1B, &gbz.d, "D", &gbz.e, "E"},
		{"_DEC_HL", 0x2B, &gbz.h, "H", &gbz.l, "L"},
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
			assertRegister(t, 0x7F, *c.Register2, c.Register2Initial)

			resetGBZ(gbz)
		})
	}
}

func Test_DEC_SP(t *testing.T) {
	gbz := createGBZWithOpcode(0x3B)
	pc := gbz.pc
	sp := gbz.sp

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+1)
	assertSP(t, sp-1, gbz.sp)
}

func Test_DECR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		Operation       string
		Opcode          uint8
		Register        *uint8
		RegisterInitial string
		InitialValue    byte
		WantValue       byte
		WantZero        bool
		WantSub         bool
		WantHalfCarry   bool
	}{
		{"_DEC_B", 0x05, &gbz.b, "B", 0x01, 0x00, true, true, false},
		{"_DEC_C", 0x0D, &gbz.c, "C", 0x01, 0x00, true, true, false},
		{"_DEC_D", 0x15, &gbz.d, "D", 0x01, 0x00, true, true, false},
		{"_DEC_E", 0x1D, &gbz.e, "E", 0x01, 0x00, true, true, false},
		{"_DEC_H", 0x25, &gbz.h, "H", 0x01, 0x00, true, true, false},
		{"_DEC_L", 0x2D, &gbz.l, "L", 0x01, 0x00, true, true, false},
		{"_DEC_A", 0x3D, &gbz.a, "A", 0x01, 0x00, true, true, false},
		{"_DEC_A_Zero", 0x3D, &gbz.a, "A", 0x00, 0xFF, false, true, true},
		{"_DEC_A_HalfCarry_0x00", 0x3D, &gbz.a, "A", 0x00, 0xFF, false, true, true},
		{"_DEC_A_HalfCarry_0x10", 0x3D, &gbz.a, "A", 0x10, 0x0F, false, true, true},
		{"_DEC_A_NoFlags", 0x3D, &gbz.a, "A", 0x02, 0x01, false, true, false},
	}

	for _, c := range cases {
		t.Run(c.Operation, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, byte(c.Opcode))
			*c.Register = c.InitialValue
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.WantValue, *c.Register, c.RegisterInitial)
			assert.Equal(t, c.WantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, c.WantSub, gbz.flags.Get(Sub), "Sub flag mismatch")
			assert.Equal(t, c.WantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")

			resetGBZ(gbz)
		})
	}
}

func Test_ADD_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		opcode        uint8
		reg           *uint8
		regName       string
		initialA      byte
		initialReg    byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
		cycles        uint
	}{
		{"ADD_A_B_NoFlags", 0x80, &gbz.b, "B", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_B_HalfCarry", 0x80, &gbz.b, "B", 0x0F, 0x01, 0x10, false, true, false, 4},
		{"ADD_A_B_Carry", 0x80, &gbz.b, "B", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_B_Zero", 0x80, &gbz.b, "B", 0x00, 0x00, 0x00, false, false, true, 4},
		{"ADD_A_C_NoFlags", 0x81, &gbz.c, "C", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_C_HalfCarry", 0x81, &gbz.c, "C", 0x0F, 0x01, 0x10, false, true, false, 4},
		{"ADD_A_C_Carry", 0x81, &gbz.c, "C", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_C_Zero", 0x81, &gbz.c, "C", 0x00, 0x00, 0x00, false, false, true, 4},
		{"ADD_A_D_NoFlags", 0x82, &gbz.d, "D", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_D_Carry", 0x82, &gbz.d, "D", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_E_NoFlags", 0x83, &gbz.e, "E", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_E_Carry", 0x83, &gbz.e, "E", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_H_NoFlags", 0x84, &gbz.h, "H", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_H_Carry", 0x84, &gbz.h, "H", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_L_NoFlags", 0x85, &gbz.l, "L", 0x01, 0x02, 0x03, false, false, false, 4},
		{"ADD_A_L_Carry", 0x85, &gbz.l, "L", 0xFF, 0x01, 0x00, true, true, true, 4},
		{"ADD_A_A_Self", 0x87, &gbz.a, "A", 0x01, 0x01, 0x02, false, false, false, 4},
		{"ADD_A_A_Carry", 0x87, &gbz.a, "A", 0x80, 0x80, 0x00, true, false, true, 4},
		{"ADD_A_A_Zero", 0x87, &gbz.a, "A", 0x00, 0x00, 0x00, false, false, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADD_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		initialA      byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"ADD_A_HLX_NoFlags", 0xC000, 0x02, 0x01, 0x03, false, false, false},
		{"ADD_A_HLX_HalfCarry", 0xC000, 0x01, 0x0F, 0x10, false, true, false},
		{"ADD_A_HLX_Carry", 0xC000, 0x01, 0xFF, 0x00, true, true, true},
		{"ADD_A_HLX_Zero", 0xC000, 0x00, 0x00, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x86)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_DEC_HLX(t *testing.T) {
	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		wantValue     byte
		wantZero      bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0xC000, 0x02, 0x01, false, false},
		{"Zero", 0xC000, 0x01, 0x00, true, false},
		{"HalfCarry_0x00", 0xC000, 0x00, 0xFF, false, true},
		{"HalfCarry_0x10", 0xC000, 0x10, 0x0F, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x35)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 12)
			assertPC(t, gbz, pc+1)
			assert.Equal(t, c.wantValue, gbz.mem.Read(c.hl), "Memory value mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
		})
	}
}

func Test_ADC_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		opcode        uint8
		reg           *uint8
		regName       string
		initialA      byte
		initialReg    byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
		cycles        uint
	}{
		{"ADC_A_B_NoCarry_NoFlags", 0x88, &gbz.b, "B", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_B_WithCarry_NoFlags", 0x88, &gbz.b, "B", 0x01, 0x02, true, 0x04, false, false, false, 4},
		{"ADC_A_B_HalfCarry", 0x88, &gbz.b, "B", 0x0F, 0x01, false, 0x10, false, true, false, 4},
		{"ADC_A_B_Carry", 0x88, &gbz.b, "B", 0xFF, 0x01, false, 0x00, true, true, true, 4},
		{"ADC_A_B_Zero", 0x88, &gbz.b, "B", 0x00, 0x00, false, 0x00, false, false, true, 4},
		{"ADC_A_C_NoCarry_NoFlags", 0x89, &gbz.c, "C", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_D_NoCarry_NoFlags", 0x8A, &gbz.d, "D", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_E_NoCarry_NoFlags", 0x8B, &gbz.e, "E", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_H_NoCarry_NoFlags", 0x8C, &gbz.h, "H", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_L_NoCarry_NoFlags", 0x8D, &gbz.l, "L", 0x01, 0x02, false, 0x03, false, false, false, 4},
		{"ADC_A_A_NoCarry_NoFlags", 0x8F, &gbz.a, "A", 0x01, 0x01, false, 0x02, false, false, false, 4},
		{"ADC_A_A_WithCarry", 0x8F, &gbz.a, "A", 0x80, 0x80, true, 0x01, true, false, false, 4},
		{"ADC_A_A_Zero", 0x8F, &gbz.a, "A", 0x00, 0x00, false, 0x00, false, false, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADC_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		initialA      byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"ADC_A_HLX_NoCarry_NoFlags", 0xC000, 0x02, 0x01, false, 0x03, false, false, false},
		{"ADC_A_HLX_HalfCarry", 0xC000, 0x01, 0x0F, false, 0x10, false, true, false},
		{"ADC_A_HLX_Carry", 0xC000, 0x01, 0xFF, false, 0x00, true, true, true},
		{"ADC_A_HLX_Zero", 0xC000, 0x00, 0x00, false, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x8E)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADC_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		imm           byte
		initialA      byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"ADC_A_N8_NoCarry_NoFlags", 0x02, 0x01, false, 0x03, false, false, false},
		{"ADC_A_N8_HalfCarry", 0x01, 0x0F, false, 0x10, false, true, false},
		{"ADC_A_N8_Carry", 0x01, 0xFF, false, 0x00, true, true, true},
		{"ADC_A_N8_Zero", 0x00, 0x00, false, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.mem.Write(gbz.pc, 0xCE)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADC_WithCarry(t *testing.T) {
	gbz := createGBZ()

	gbz.mem.Write(gbz.pc, 0x88)
	gbz.a = 0x01
	gbz.b = 0x02
	gbz.flags.Set(Carry, true)
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assertRegister(t, 0x04, gbz.a, "A")
	assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")
	assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
	assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
	assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
}

func Test_SUB_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		opcode        uint8
		reg           *uint8
		regName       string
		initialA      byte
		initialReg    byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
		cycles        uint
	}{
		{"SUB_A_B_NoFlags", 0x90, &gbz.b, "B", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_B_HalfCarry", 0x90, &gbz.b, "B", 0x10, 0x01, 0x0F, false, true, false, 4},
		{"SUB_A_B_Carry", 0x90, &gbz.b, "B", 0x00, 0x01, 0xFF, true, true, false, 4},
		{"SUB_A_B_Zero", 0x90, &gbz.b, "B", 0x02, 0x02, 0x00, false, false, true, 4},
		{"SUB_A_C_NoFlags", 0x91, &gbz.c, "C", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_D_NoFlags", 0x92, &gbz.d, "D", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_E_NoFlags", 0x93, &gbz.e, "E", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_H_NoFlags", 0x94, &gbz.h, "H", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_L_NoFlags", 0x95, &gbz.l, "L", 0x03, 0x02, 0x01, false, false, false, 4},
		{"SUB_A_A_Zero", 0x97, &gbz.a, "A", 0x02, 0x02, 0x00, false, false, true, 4},
		{"SUB_A_A_Carry", 0x97, &gbz.a, "A", 0x01, 0x01, 0x00, false, false, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_SUB_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		initialA      byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"SUB_A_HLX_NoFlags", 0xC000, 0x02, 0x03, 0x01, false, false, false},
		{"SUB_A_HLX_HalfCarry", 0xC000, 0x01, 0x10, 0x0F, false, true, false},
		{"SUB_A_HLX_Carry", 0xC000, 0x01, 0x00, 0xFF, true, true, false},
		{"SUB_A_HLX_Zero", 0xC000, 0x02, 0x02, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x96)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_SUB_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		imm           byte
		initialA      byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"SUB_A_N8_NoFlags", 0x02, 0x03, 0x01, false, false, false},
		{"SUB_A_N8_HalfCarry", 0x01, 0x10, 0x0F, false, true, false},
		{"SUB_A_N8_Carry", 0x01, 0x00, 0xFF, true, true, false},
		{"SUB_A_N8_Zero", 0x02, 0x02, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xD6)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_SBC_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		opcode        uint8
		reg           *uint8
		regName       string
		initialA      byte
		initialReg    byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
		cycles        uint
	}{
		{"SBC_A_B_NoCarry_NoFlags", 0x98, &gbz.b, "B", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_B_WithCarry_NoFlags", 0x98, &gbz.b, "B", 0x03, 0x02, true, 0x00, false, false, true, 4},
		{"SBC_A_B_HalfCarry", 0x98, &gbz.b, "B", 0x10, 0x01, false, 0x0F, false, true, false, 4},
		{"SBC_A_B_WithCarry_HalfCarry", 0x98, &gbz.b, "B", 0x10, 0x01, true, 0x0E, false, true, false, 4},
		{"SBC_A_B_Carry", 0x98, &gbz.b, "B", 0x00, 0x01, false, 0xFF, true, true, false, 4},
		{"SBC_A_C_NoCarry_NoFlags", 0x99, &gbz.c, "C", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_D_NoCarry_NoFlags", 0x9A, &gbz.d, "D", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_E_NoCarry_NoFlags", 0x9B, &gbz.e, "E", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_H_NoCarry_NoFlags", 0x9C, &gbz.h, "H", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_L_NoCarry_NoFlags", 0x9D, &gbz.l, "L", 0x03, 0x02, false, 0x01, false, false, false, 4},
		{"SBC_A_A_NoCarry_Zero", 0x9F, &gbz.a, "A", 0x02, 0x02, false, 0x00, false, false, true, 4},
		{"SBC_A_A_WithCarry_Zero", 0x9F, &gbz.a, "A", 0x01, 0x01, true, 0xFF, true, true, false, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_SBC_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		initialA      byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"SBC_A_HLX_NoCarry_NoFlags", 0xC000, 0x02, 0x03, false, 0x01, false, false, false},
		{"SBC_A_HLX_WithCarry_NoFlags", 0xC000, 0x02, 0x03, true, 0x00, false, false, true},
		{"SBC_A_HLX_HalfCarry", 0xC000, 0x01, 0x10, false, 0x0F, false, true, false},
		{"SBC_A_HLX_Carry", 0xC000, 0x01, 0x00, false, 0xFF, true, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0x9E)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_SBC_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		imm           byte
		initialA      byte
		initialCarry  bool
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"SBC_A_N8_NoCarry_NoFlags", 0x02, 0x03, false, 0x01, false, false, false},
		{"SBC_A_N8_WithCarry_NoFlags", 0x02, 0x03, true, 0x00, false, false, true},
		{"SBC_A_N8_HalfCarry", 0x01, 0x10, false, 0x0F, false, true, false},
		{"SBC_A_N8_Carry", 0x01, 0x00, false, 0xFF, true, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.flags.Set(Carry, c.initialCarry)
			gbz.mem.Write(gbz.pc, 0xDE)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_AND_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name       string
		opcode     uint8
		reg        *uint8
		regName    string
		initialA   byte
		initialReg byte
		wantValue  byte
		wantZero   bool
		cycles     uint
	}{
		{"AND_A_B_NoFlags", 0xA0, &gbz.b, "B", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_B_Zero", 0xA0, &gbz.b, "B", 0xFF, 0x00, 0x00, true, 4},
		{"AND_A_C_NoFlags", 0xA1, &gbz.c, "C", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_D_NoFlags", 0xA2, &gbz.d, "D", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_E_NoFlags", 0xA3, &gbz.e, "E", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_H_NoFlags", 0xA4, &gbz.h, "H", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_L_NoFlags", 0xA5, &gbz.l, "L", 0xFF, 0x0F, 0x0F, false, 4},
		{"AND_A_A_NoFlags", 0xA7, &gbz.a, "A", 0xFF, 0xFF, 0xFF, false, 4},
		{"AND_A_A_Zero", 0xA7, &gbz.a, "A", 0x00, 0x00, 0x00, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, true, gbz.flags.Get(HalfCarry), "HalfCarry flag should be true")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_AND_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		hl        uint16
		memValue  byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"AND_A_HLX_NoFlags", 0xC000, 0x0F, 0xFF, 0x0F, false},
		{"AND_A_HLX_Zero", 0xC000, 0x00, 0xFF, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0xA6)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, true, gbz.flags.Get(HalfCarry), "HalfCarry flag should be true")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_AND_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		imm       byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"AND_A_N8_NoFlags", 0x0F, 0xFF, 0x0F, false},
		{"AND_A_N8_Zero", 0x00, 0xFF, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xE6)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, true, gbz.flags.Get(HalfCarry), "HalfCarry flag should be true")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_XOR_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name       string
		opcode     uint8
		reg        *uint8
		regName    string
		initialA   byte
		initialReg byte
		wantValue  byte
		wantZero   bool
		cycles     uint
	}{
		{"XOR_A_B_NoFlags", 0xA8, &gbz.b, "B", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_B_Zero", 0xA8, &gbz.b, "B", 0xFF, 0xFF, 0x00, true, 4},
		{"XOR_A_C_NoFlags", 0xA9, &gbz.c, "C", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_D_NoFlags", 0xAA, &gbz.d, "D", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_E_NoFlags", 0xAB, &gbz.e, "E", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_H_NoFlags", 0xAC, &gbz.h, "H", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_L_NoFlags", 0xAD, &gbz.l, "L", 0xFF, 0x0F, 0xF0, false, 4},
		{"XOR_A_A_Zero", 0xAF, &gbz.a, "A", 0xFF, 0xFF, 0x00, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_XOR_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		hl        uint16
		memValue  byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"XOR_A_HLX_NoFlags", 0xC000, 0x0F, 0xFF, 0xF0, false},
		{"XOR_A_HLX_Zero", 0xC000, 0xFF, 0xFF, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0xAE)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_XOR_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		imm       byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"XOR_A_N8_NoFlags", 0x0F, 0xFF, 0xF0, false},
		{"XOR_A_N8_Zero", 0xFF, 0xFF, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xEE)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_OR_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name       string
		opcode     uint8
		reg        *uint8
		regName    string
		initialA   byte
		initialReg byte
		wantValue  byte
		wantZero   bool
		cycles     uint
	}{
		{"OR_A_B_NoFlags", 0xB0, &gbz.b, "B", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_B_Zero", 0xB0, &gbz.b, "B", 0x00, 0x00, 0x00, true, 4},
		{"OR_A_C_NoFlags", 0xB1, &gbz.c, "C", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_D_NoFlags", 0xB2, &gbz.d, "D", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_E_NoFlags", 0xB3, &gbz.e, "E", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_H_NoFlags", 0xB4, &gbz.h, "H", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_L_NoFlags", 0xB5, &gbz.l, "L", 0x0F, 0xF0, 0xFF, false, 4},
		{"OR_A_A_NoFlags", 0xB7, &gbz.a, "A", 0x42, 0x42, 0x42, false, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_OR_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		hl        uint16
		memValue  byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"OR_A_HLX_NoFlags", 0xC000, 0xF0, 0x0F, 0xFF, false},
		{"OR_A_HLX_Zero", 0xC000, 0x00, 0x00, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0xB6)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_OR_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		imm       byte
		initialA  byte
		wantValue byte
		wantZero  bool
	}{
		{"OR_A_N8_NoFlags", 0xF0, 0x0F, 0xFF, false},
		{"OR_A_N8_Zero", 0x00, 0x00, 0x00, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xF6)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_CP_AR8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		opcode        uint8
		reg           *uint8
		regName       string
		initialA      byte
		initialReg    byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
		cycles        uint
	}{
		{"CP_A_B_NoFlags", 0xB8, &gbz.b, "B", 0x03, 0x02, false, false, false, 4},
		{"CP_A_B_Carry", 0xB8, &gbz.b, "B", 0x00, 0x01, true, true, false, 4},
		{"CP_A_B_Zero", 0xB8, &gbz.b, "B", 0x02, 0x02, false, false, true, 4},
		{"CP_A_B_HalfCarry", 0xB8, &gbz.b, "B", 0x10, 0x01, false, true, false, 4},
		{"CP_A_C_NoFlags", 0xB9, &gbz.c, "C", 0x03, 0x02, false, false, false, 4},
		{"CP_A_D_NoFlags", 0xBA, &gbz.d, "D", 0x03, 0x02, false, false, false, 4},
		{"CP_A_E_NoFlags", 0xBB, &gbz.e, "E", 0x03, 0x02, false, false, false, 4},
		{"CP_A_H_NoFlags", 0xBC, &gbz.h, "H", 0x03, 0x02, false, false, false, 4},
		{"CP_A_L_NoFlags", 0xBD, &gbz.l, "L", 0x03, 0x02, false, false, false, 4},
		{"CP_A_A_Zero", 0xBF, &gbz.a, "A", 0x02, 0x02, false, false, true, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, c.opcode)
			gbz.a = c.initialA
			*c.reg = c.initialReg
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.initialA, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_CP_A_HLX(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		hl            uint16
		memValue      byte
		initialA      byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"CP_A_HLX_NoFlags", 0xC000, 0x02, 0x03, false, false, false},
		{"CP_A_HLX_Carry", 0xC000, 0x01, 0x00, true, true, false},
		{"CP_A_HLX_Zero", 0xC000, 0x02, 0x02, false, false, true},
		{"CP_A_HLX_HalfCarry", 0xC000, 0x01, 0x10, false, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.h = byte(c.hl >> 8)
			gbz.l = byte(c.hl & 0xFF)
			gbz.mem.Write(c.hl, c.memValue)
			gbz.mem.Write(gbz.pc, 0xBE)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.initialA, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_CP_A_N8(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name          string
		imm           byte
		initialA      byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"CP_A_N8_NoFlags", 0x02, 0x03, false, false, false},
		{"CP_A_N8_Carry", 0x01, 0x00, true, true, false},
		{"CP_A_N8_Zero", 0x02, 0x02, false, false, true},
		{"CP_A_N8_HalfCarry", 0x01, 0x10, false, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xFE)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.initialA, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")

			resetGBZ(gbz)
		})
	}
}

func Test_ADD_HL_BC(t *testing.T) {
	cases := []struct {
		name          string
		opcode        uint8
		b, c          uint8
		h, l          uint8
		wantH, wantL  uint8
		wantCarry     bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0x09, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, false, false},
		{"Carry", 0x09, 0xFF, 0xFF, 0x00, 0x01, 0x00, 0x00, true, true},
		{"CarryHalfCarry", 0x09, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, true, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.b, gbz.c = c.b, c.c
			gbz.h, gbz.l = c.h, c.l
			gbz.mem.Write(gbz.pc, c.opcode)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantH, gbz.h, "H")
			assertRegister(t, c.wantL, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
		})
	}
}

func Test_ADD_HL_DE(t *testing.T) {
	cases := []struct {
		name          string
		opcode        uint8
		d, e          uint8
		h, l          uint8
		wantH, wantL  uint8
		wantCarry     bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0x19, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, false, false},
		{"Carry", 0x19, 0xFF, 0xFF, 0x00, 0x01, 0x00, 0x00, true, true},
		{"CarryHalfCarry", 0x19, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, true, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.d, gbz.e = c.d, c.e
			gbz.h, gbz.l = c.h, c.l
			gbz.mem.Write(gbz.pc, c.opcode)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantH, gbz.h, "H")
			assertRegister(t, c.wantL, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
		})
	}
}

func Test_ADD_HL_HL(t *testing.T) {
	cases := []struct {
		name          string
		h, l          uint8
		wantH, wantL  uint8
		wantCarry     bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0x00, 0x00, 0x00, 0x00, false, false},
		{"HalfCarry_0x0800", 0x08, 0x00, 0x10, 0x00, false, true},
		{"Carry", 0x80, 0x00, 0x00, 0x00, true, false},
		{"CarryHalfCarry", 0xFF, 0xFF, 0xFF, 0xFE, true, true},
		{"HalfCarry_0x0FFF", 0x0F, 0xFF, 0x1F, 0xFE, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.h, gbz.l = c.h, c.l
			gbz.mem.Write(gbz.pc, 0x29)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantH, gbz.h, "H")
			assertRegister(t, c.wantL, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
		})
	}
}

func Test_ADD_HL_SP(t *testing.T) {
	cases := []struct {
		name          string
		sp            uint16
		h, l          uint8
		wantH, wantL  uint8
		wantCarry     bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0x0001, 0x00, 0x00, 0x00, 0x01, false, false},
		{"Carry", 0xFFFF, 0x00, 0x01, 0x00, 0x00, true, true},
		{"CarryHalfCarry", 0xFFFF, 0xFF, 0xFF, 0xFF, 0xFE, true, true},
		{"HalfCarry_0x0FFF", 0x0001, 0x0F, 0xFF, 0x10, 0x00, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.sp = c.sp
			gbz.h, gbz.l = c.h, c.l
			gbz.mem.Write(gbz.pc, 0x39)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantH, gbz.h, "H")
			assertRegister(t, c.wantL, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
		})
	}
}

func Test_JR_E8(t *testing.T) {
	cases := []struct {
		name   string
		offset int8
		wantPC uint16
		cycles uint
	}{
		{"PositiveOffset", 0x10, 0x0112, 12},
		{"NegativeOffset", -16, 0x00F2, 12},
		{"ZeroOffset", 0x00, 0x0102, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0x18)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JR_NZ_E8(t *testing.T) {
	cases := []struct {
		name       string
		zeroFlag   bool
		offset     int8
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenNotZero_Positive", false, 0x10, 0x0112, 12},
		{"JumpWhenNotZero_Negative", false, -16, 0x00F2, 12},
		{"NoJumpWhenZero", true, 0x10, 0x0102, 8},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.mem.Write(gbz.pc, 0x20)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JR_Z_E8(t *testing.T) {
	cases := []struct {
		name       string
		zeroFlag   bool
		offset     int8
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenZero_Positive", true, 0x10, 0x0112, 12},
		{"JumpWhenZero_Negative", true, -16, 0x00F2, 12},
		{"NoJumpWhenNotZero", false, 0x10, 0x0102, 8},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.mem.Write(gbz.pc, 0x28)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JR_NC_E8(t *testing.T) {
	cases := []struct {
		name       string
		carryFlag  bool
		offset     int8
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenNotCarry_Positive", false, 0x10, 0x0112, 12},
		{"JumpWhenNotCarry_Negative", false, -16, 0x00F2, 12},
		{"NoJumpWhenCarry", true, 0x10, 0x0102, 8},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.mem.Write(gbz.pc, 0x30)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JR_C_E8(t *testing.T) {
	cases := []struct {
		name       string
		carryFlag  bool
		offset     int8
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenCarry_Positive", true, 0x10, 0x0112, 12},
		{"JumpWhenCarry_Negative", true, -16, 0x00F2, 12},
		{"NoJumpWhenNotCarry", false, 0x10, 0x0102, 8},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.mem.Write(gbz.pc, 0x38)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JP_A16(t *testing.T) {
	gbz := createGBZ()
	targetAddr := uint16(0x1234)
	gbz.mem.Write(gbz.pc, 0xC3)
	gbz.mem.Write(gbz.pc+1, byte(targetAddr))
	gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, targetAddr)
}

func Test_JP_NZ_A16(t *testing.T) {
	cases := []struct {
		name       string
		zeroFlag   bool
		targetAddr uint16
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenNotZero", false, 0x1234, 0x1234, 16},
		{"NoJumpWhenZero", true, 0x1234, 0x0103, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.mem.Write(gbz.pc, 0xC2)
			gbz.mem.Write(gbz.pc+1, byte(c.targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(c.targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JP_Z_A16(t *testing.T) {
	cases := []struct {
		name       string
		zeroFlag   bool
		targetAddr uint16
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenZero", true, 0x1234, 0x1234, 16},
		{"NoJumpWhenNotZero", false, 0x1234, 0x0103, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.mem.Write(gbz.pc, 0xCA)
			gbz.mem.Write(gbz.pc+1, byte(c.targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(c.targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JP_NC_A16(t *testing.T) {
	cases := []struct {
		name       string
		carryFlag  bool
		targetAddr uint16
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenNotCarry", false, 0x1234, 0x1234, 16},
		{"NoJumpWhenCarry", true, 0x1234, 0x0103, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.mem.Write(gbz.pc, 0xD2)
			gbz.mem.Write(gbz.pc+1, byte(c.targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(c.targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JP_C_A16(t *testing.T) {
	cases := []struct {
		name       string
		carryFlag  bool
		targetAddr uint16
		wantPC     uint16
		wantCycles uint
	}{
		{"JumpWhenCarry", true, 0x1234, 0x1234, 16},
		{"NoJumpWhenNotCarry", false, 0x1234, 0x0103, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.mem.Write(gbz.pc, 0xDA)
			gbz.mem.Write(gbz.pc+1, byte(c.targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(c.targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			resetGBZ(gbz)
		})
	}
}

func Test_JP_HL(t *testing.T) {
	gbz := createGBZ()
	gbz.h = 0x12
	gbz.l = 0x34
	gbz.mem.Write(gbz.pc, 0xE9)

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, 0x1234)
}

func Test_RET(t *testing.T) {
	gbz := createGBZ()
	returnAddr := uint16(0x1234)
	gbz.sp = 0xFFFE
	gbz.mem.Write(gbz.sp, byte(returnAddr))
	gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
	gbz.mem.Write(gbz.pc, 0xC9)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, returnAddr)
	assertSP(t, 0x0000, gbz.sp)
}

func Test_RET_NZ(t *testing.T) {
	returnAddr := uint16(0x1234)

	cases := []struct {
		name       string
		zeroFlag   bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"ReturnWhenNotZero", false, returnAddr, 20, 0x0000},
		{"NoReturnWhenZero", true, 0x0101, 8, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.sp, byte(returnAddr))
			gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
			gbz.mem.Write(gbz.pc, 0xC0)

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_RET_Z(t *testing.T) {
	returnAddr := uint16(0x1234)

	cases := []struct {
		name       string
		zeroFlag   bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"ReturnWhenZero", true, returnAddr, 20, 0x0000},
		{"NoReturnWhenNotZero", false, 0x0101, 8, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.sp, byte(returnAddr))
			gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
			gbz.mem.Write(gbz.pc, 0xC8)

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_RET_NC(t *testing.T) {
	returnAddr := uint16(0x1234)

	cases := []struct {
		name       string
		carryFlag  bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"ReturnWhenNotCarry", false, returnAddr, 20, 0x0000},
		{"NoReturnWhenCarry", true, 0x0101, 8, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.sp, byte(returnAddr))
			gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
			gbz.mem.Write(gbz.pc, 0xD0)

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_RET_C(t *testing.T) {
	returnAddr := uint16(0x1234)

	cases := []struct {
		name       string
		carryFlag  bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"ReturnWhenCarry", true, returnAddr, 20, 0x0000},
		{"NoReturnWhenNotCarry", false, 0x0101, 8, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.sp, byte(returnAddr))
			gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
			gbz.mem.Write(gbz.pc, 0xD8)

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_RETI(t *testing.T) {
	gbz := createGBZ()
	returnAddr := uint16(0x1234)
	gbz.sp = 0xFFFE
	gbz.InterruptEnabled = false
	gbz.mem.Write(gbz.sp, byte(returnAddr))
	gbz.mem.Write(gbz.sp+1, byte(returnAddr>>8))
	gbz.mem.Write(gbz.pc, 0xD9)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, returnAddr)
	assertSP(t, 0x0000, gbz.sp)
	assert.Equal(t, true, gbz.InterruptEnabled, "RETI should enable interrupts")
}

func Test_CALL(t *testing.T) {
	gbz := createGBZ()
	targetAddr := uint16(0x1234)
	gbz.sp = 0xFFFE
	gbz.mem.Write(gbz.pc, 0xCD)
	gbz.mem.Write(gbz.pc+1, byte(targetAddr))
	gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

	gbz.Run()

	assertCycles(t, gbz, 24)
	assertPC(t, gbz, targetAddr)
	assertSP(t, 0xFFFC, gbz.sp)
	assertMemory(t, gbz, 0xFFFC, byte(0x103&0xFF))
	assertMemory(t, gbz, 0xFFFD, byte(0x103>>8))
}

func Test_CALL_NZ(t *testing.T) {
	targetAddr := uint16(0x1234)

	cases := []struct {
		name       string
		zeroFlag   bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"CallWhenNotZero", false, targetAddr, 24, 0xFFFC},
		{"NoCallWhenZero", true, 0x0103, 12, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.pc, 0xC4)
			gbz.mem.Write(gbz.pc+1, byte(targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_CALL_Z(t *testing.T) {
	targetAddr := uint16(0x1234)

	cases := []struct {
		name       string
		zeroFlag   bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"CallWhenZero", true, targetAddr, 24, 0xFFFC},
		{"NoCallWhenNotZero", false, 0x0103, 12, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Zero, c.zeroFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.pc, 0xCC)
			gbz.mem.Write(gbz.pc+1, byte(targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_CALL_NC(t *testing.T) {
	targetAddr := uint16(0x1234)

	cases := []struct {
		name       string
		carryFlag  bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"CallWhenNotCarry", false, targetAddr, 24, 0xFFFC},
		{"NoCallWhenCarry", true, 0x0103, 12, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.pc, 0xD4)
			gbz.mem.Write(gbz.pc+1, byte(targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_CALL_C(t *testing.T) {
	targetAddr := uint16(0x1234)

	cases := []struct {
		name       string
		carryFlag  bool
		wantPC     uint16
		wantCycles uint
		wantSP     uint16
	}{
		{"CallWhenCarry", true, targetAddr, 24, 0xFFFC},
		{"NoCallWhenNotCarry", false, 0x0103, 12, 0xFFFE},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.flags.Set(Carry, c.carryFlag)
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.pc, 0xDC)
			gbz.mem.Write(gbz.pc+1, byte(targetAddr))
			gbz.mem.Write(gbz.pc+2, byte(targetAddr>>8))

			gbz.Run()

			assertCycles(t, gbz, c.wantCycles)
			assertPC(t, gbz, c.wantPC)
			assertSP(t, c.wantSP, gbz.sp)
			resetGBZ(gbz)
		})
	}
}

func Test_RST(t *testing.T) {
	cases := []struct {
		name       string
		opcode     uint8
		targetAddr uint16
	}{
		{"RST_00", 0xC7, 0x0000},
		{"RST_08", 0xCF, 0x0008},
		{"RST_10", 0xD7, 0x0010},
		{"RST_18", 0xDF, 0x0018},
		{"RST_20", 0xE7, 0x0020},
		{"RST_28", 0xEF, 0x0028},
		{"RST_30", 0xF7, 0x0030},
		{"RST_38", 0xFF, 0x0038},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.sp = 0xFFFE
			gbz.mem.Write(gbz.pc, c.opcode)

			gbz.Run()

			assertCycles(t, gbz, 16)
			assertPC(t, gbz, c.targetAddr)
			assertSP(t, 0xFFFC, gbz.sp)
			assertMemory(t, gbz, 0xFFFC, byte(0x101&0xFF))
			assertMemory(t, gbz, 0xFFFD, byte(0x101>>8))
			resetGBZ(gbz)
		})
	}
}

func Test_LDH_A8_A(t *testing.T) {
	gbz := createGBZ()
	gbz.a = 0x42
	offset := byte(0x10)
	gbz.mem.Write(gbz.pc, 0xE0)
	gbz.mem.Write(gbz.pc+1, offset)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x102)
	assertMemory(t, gbz, 0xFF00+uint16(offset), 0x42)
}

func Test_LDH_A_A8(t *testing.T) {
	gbz := createGBZ()
	offset := byte(0x10)
	gbz.mem.Write(0xFF00+uint16(offset), 0x42)
	gbz.mem.Write(gbz.pc, 0xF0)
	gbz.mem.Write(gbz.pc+1, offset)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x102)
	assertRegister(t, 0x42, gbz.a, "A")
}

func Test_LDH_C_A(t *testing.T) {
	gbz := createGBZ()
	gbz.a = 0x42
	gbz.c = 0x10
	gbz.mem.Write(gbz.pc, 0xE2)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, 0x101)
	assertMemory(t, gbz, 0xFF00+uint16(gbz.c), 0x42)
}

func Test_LDH_A_C(t *testing.T) {
	gbz := createGBZ()
	gbz.c = 0x10
	gbz.mem.Write(0xFF00+uint16(gbz.c), 0x42)
	gbz.mem.Write(gbz.pc, 0xF2)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, 0x101)
	assertRegister(t, 0x42, gbz.a, "A")
}

func Test_LD_SP_HL(t *testing.T) {
	gbz := createGBZ()
	gbz.h = 0x12
	gbz.l = 0x34
	gbz.mem.Write(gbz.pc, 0xF9)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, 0x101)
	assertSP(t, 0x1234, gbz.sp)
}

func Test_LD_HL_SP_E8(t *testing.T) {
	cases := []struct {
		name          string
		sp            uint16
		offset        int8
		wantHL        uint16
		wantCarry     bool
		wantHalfCarry bool
		cycles        uint
	}{
		{"NoFlags", 0x0001, 0x01, 0x0002, false, false, 12},
		{"HalfCarryOnly", 0x000F, 0x01, 0x0010, false, true, 12},
		{"CarryFromFFFF", 0xFFFF, 0x01, 0x0000, true, true, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.sp = c.sp
			gbz.mem.Write(gbz.pc, 0xF8)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, c.cycles)
			assertPC(t, gbz, 0x102)
			assert.Equal(t, c.wantHL, (uint16(gbz.h)<<8)|uint16(gbz.l), "HL mismatch")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			resetGBZ(gbz)
		})
	}
}

func Test_POP_BC(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFC
	gbz.mem.Write(gbz.sp, 0x34)
	gbz.mem.Write(gbz.sp+1, 0x12)
	gbz.mem.Write(gbz.pc, 0xC1)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x101)
	assertRegister(t, 0x12, gbz.b, "B")
	assertRegister(t, 0x34, gbz.c, "C")
	assertSP(t, 0xFFFE, gbz.sp)
}

func Test_POP_DE(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFC
	gbz.mem.Write(gbz.sp, 0x56)
	gbz.mem.Write(gbz.sp+1, 0x78)
	gbz.mem.Write(gbz.pc, 0xD1)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x101)
	assertRegister(t, 0x78, gbz.d, "D")
	assertRegister(t, 0x56, gbz.e, "E")
	assertSP(t, 0xFFFE, gbz.sp)
}

func Test_POP_HL(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFC
	gbz.mem.Write(gbz.sp, 0x9A)
	gbz.mem.Write(gbz.sp+1, 0xBC)
	gbz.mem.Write(gbz.pc, 0xE1)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x101)
	assertRegister(t, 0xBC, gbz.h, "H")
	assertRegister(t, 0x9A, gbz.l, "L")
	assertSP(t, 0xFFFE, gbz.sp)
}

func Test_POP_AF(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFC
	gbz.mem.Write(gbz.sp, 0x00)
	gbz.mem.Write(gbz.sp+1, 0x12)
	gbz.mem.Write(gbz.pc, 0xF1)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, 0x101)
	assertRegister(t, 0x12, gbz.a, "A")
	assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag mismatch")
	assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag mismatch")
	assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
	assert.Equal(t, false, gbz.flags.Get(Carry), "Carry flag mismatch")
	assertSP(t, 0xFFFE, gbz.sp)
}

func Test_PUSH_BC(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFE
	gbz.b = 0x12
	gbz.c = 0x34
	gbz.mem.Write(gbz.pc, 0xC5)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, 0x101)
	assert.Equal(t, uint8(0x34), gbz.mem.Read(gbz.sp), "Stack low byte mismatch")
	assert.Equal(t, uint8(0x12), gbz.mem.Read(gbz.sp+1), "Stack high byte mismatch")
	assertSP(t, 0xFFFC, gbz.sp)
}

func Test_PUSH_DE(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFE
	gbz.d = 0xAB
	gbz.e = 0xCD
	gbz.mem.Write(gbz.pc, 0xD5)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, 0x101)
	assert.Equal(t, uint8(0xCD), gbz.mem.Read(gbz.sp), "Stack low byte mismatch")
	assert.Equal(t, uint8(0xAB), gbz.mem.Read(gbz.sp+1), "Stack high byte mismatch")
	assertSP(t, 0xFFFC, gbz.sp)
}

func Test_PUSH_HL(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFE
	gbz.h = 0xDE
	gbz.l = 0xAD
	gbz.mem.Write(gbz.pc, 0xE5)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, 0x101)
	assert.Equal(t, uint8(0xAD), gbz.mem.Read(gbz.sp), "Stack low byte mismatch")
	assert.Equal(t, uint8(0xDE), gbz.mem.Read(gbz.sp+1), "Stack high byte mismatch")
	assertSP(t, 0xFFFC, gbz.sp)
}

func Test_PUSH_AF(t *testing.T) {
	gbz := createGBZ()
	gbz.sp = 0xFFFE
	gbz.a = 0x12
	gbz.flags.Set(Zero, true)
	gbz.flags.Set(Sub, true)
	gbz.flags.Set(HalfCarry, true)
	gbz.flags.Set(Carry, true)
	gbz.mem.Write(gbz.pc, 0xF5)

	gbz.Run()

	assertCycles(t, gbz, 16)
	assertPC(t, gbz, 0x101)
	assert.Equal(t, uint8(0xF0), gbz.mem.Read(gbz.sp), "Stack low byte (flags) mismatch")
	assert.Equal(t, uint8(0x12), gbz.mem.Read(gbz.sp+1), "Stack high byte (A) mismatch")
	assert.Equal(t, uint16(0xFFFC), gbz.sp, "SP mismatch")
}

func Test_RLCA(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialA  byte
		wantValue byte
		wantCarry bool
	}{
		{"Rotate_0x00_NoCarry", 0x00, 0x00, false},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true},
		{"Rotate_0x81_Carry", 0x81, 0x03, true},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0x07)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRCA(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialA  byte
		wantValue byte
		wantCarry bool
	}{
		{"Rotate_0x00_NoCarry", 0x00, 0x00, false},
		{"Rotate_0x01_Carry", 0x01, 0x80, true},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0x0F)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLA(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name         string
		initialA     byte
		initialCarry bool
		wantValue    byte
		wantCarry    bool
	}{
		{"Rotate_0x00_NoCarry_In", 0x00, false, 0x00, false},
		{"Rotate_0x00_Carry_In", 0x00, true, 0x01, false},
		{"Rotate_0x01_NoCarry_In", 0x01, false, 0x02, false},
		{"Rotate_0x01_Carry_In", 0x01, true, 0x03, false},
		{"Rotate_0x80_NoCarry_In", 0x80, false, 0x00, true},
		{"Rotate_0x80_Carry_In", 0x80, true, 0x01, true},
		{"Rotate_0xFF_NoCarry_In", 0xFF, false, 0xFE, true},
		{"Rotate_0xFF_Carry_In", 0xFF, true, 0xFF, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0x17)
			gbz.a = c.initialA
			gbz.flags.Set(Carry, c.initialCarry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRA(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name         string
		initialA     byte
		initialCarry bool
		wantValue    byte
		wantCarry    bool
	}{
		{"Rotate_0x00_NoCarry_In", 0x00, false, 0x00, false},
		{"Rotate_0x00_Carry_In", 0x00, true, 0x80, false},
		{"Rotate_0x01_NoCarry_In", 0x01, false, 0x00, true},
		{"Rotate_0x01_Carry_In", 0x01, true, 0x80, true},
		{"Rotate_0x80_NoCarry_In", 0x80, false, 0x40, false},
		{"Rotate_0x80_Carry_In", 0x80, true, 0xC0, false},
		{"Rotate_0xFF_NoCarry_In", 0xFF, false, 0x7F, true},
		{"Rotate_0xFF_Carry_In", 0xFF, true, 0xFF, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0x1F)
			gbz.a = c.initialA
			gbz.flags.Set(Carry, c.initialCarry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_B(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialB  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x00)
			gbz.b = c.initialB
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.b, "B")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_C(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialC  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x01)
			gbz.c = c.initialC
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.c, "C")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_D(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialD  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x02)
			gbz.d = c.initialD
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.d, "D")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_E(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialE  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x03)
			gbz.e = c.initialE
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.e, "E")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_H(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialH  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x04)
			gbz.h = c.initialH
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.h, "H")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_L(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialL  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x05)
			gbz.l = c.initialL
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_HL(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialHL byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x06)
			hl := uint16(gbz.h)<<8 | uint16(gbz.l)
			gbz.mem.Write(hl, c.initialHL)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 16)
			assertPC(t, gbz, pc+2)
			assert.Equal(t, c.wantValue, gbz.mem.Read(hl), "Memory at HL mismatch")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RLC_A(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialA  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_NoCarry", 0x01, 0x02, false, false},
		{"Rotate_0x7F_NoCarry", 0x7F, 0xFE, false, false},
		{"Rotate_0x80_Carry", 0x80, 0x01, true, false},
		{"Rotate_0x81_Carry", 0x81, 0x03, true, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x07)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_HALT(t *testing.T) {
	gbz := createGBZ()
	gbz.Halt = false
	pc := gbz.pc

	gbz.mem.Write(gbz.pc, 0x76)

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+1)
	assert.Equal(t, true, gbz.Halt, "Halt flag should be true")
}

func Test_STOP(t *testing.T) {
	gbz := createGBZ()
	gbz.IsStopped = false
	pc := gbz.pc

	gbz.mem.Write(gbz.pc, 0x10)
	gbz.mem.Write(gbz.pc+1, 0x00)

	gbz.Run()

	assertCycles(t, gbz, 4)
	assertPC(t, gbz, pc+2)
	assert.Equal(t, true, gbz.IsStopped, "IsStopped flag should be true")
}

func Test_DAA(t *testing.T) {
	cases := []struct {
		name           string
		initialA       byte
		sub, halfCarry, carry bool
		wantValue      byte
		wantZero       bool
	}{
		{"Add_0x01_0x02_NoFlags", 0x03, false, false, false, 0x03, false},
		{"Add_0x09_0x01_HalfCarry", 0x0A, false, true, false, 0x10, false},
		{"Add_0x80_0x80_Carry", 0x00, false, false, true, 0x60, false},
		{"Add_0x01_0x01_NoCarry", 0x02, false, false, false, 0x02, false},
		{"Add_0x0F_0x01_HalfCarry", 0x10, false, true, false, 0x16, false},
		{"Add_0x99_0x01_Carry", 0x9A, false, false, true, 0x00, true},
		{"Add_0x9A_0x01_Carry", 0x9B, false, true, true, 0x01, false},
		{"Sub_0x00_0x01_Carry_HalfCarry", 0xFF, true, true, true, 0x99, false},
		{"Sub_0x01_0x01_NoFlags", 0x00, true, false, false, 0x00, true},
		{"Sub_0x10_0x01_HalfCarry", 0x0F, true, true, false, 0x09, false},
		{"Sub_0x30_0x01_HalfCarry", 0x2F, true, true, false, 0x29, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0x27)
			gbz.a = c.initialA
			gbz.flags.Set(Sub, c.sub)
			gbz.flags.Set(HalfCarry, c.halfCarry)
			gbz.flags.Set(Carry, c.carry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")
		})
	}
}

func Test_CPL(t *testing.T) {
	cases := []struct {
		name         string
		initialA     byte
		wantValue    byte
	}{
		{"Complement_0x00", 0x00, 0xFF},
		{"Complement_0xFF", 0xFF, 0x00},
		{"Complement_0x42", 0x42, 0xBD},
		{"Complement_0xAA", 0xAA, 0x55},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0x2F)
			gbz.a = c.initialA
			initialCarry := gbz.flags.Get(Carry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, true, gbz.flags.Get(Sub), "Sub flag should be true")
			assert.Equal(t, true, gbz.flags.Get(HalfCarry), "HalfCarry flag should be true")
			assert.Equal(t, initialCarry, gbz.flags.Get(Carry), "Carry flag should be unchanged")

			resetGBZ(gbz)
		})
	}
}

func Test_SCF(t *testing.T) {
	cases := []struct {
		name         string
		initialCarry bool
	}{
		{"SetCarry_FromClear", false},
		{"SetCarry_FromSet", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0x37)
			gbz.flags.Set(Carry, c.initialCarry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assert.Equal(t, true, gbz.flags.Get(Carry), "Carry flag should be true")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_CCF(t *testing.T) {
	cases := []struct {
		name         string
		initialCarry bool
		wantCarry    bool
	}{
		{"ToggleCarry_ClearToSet", false, true},
		{"ToggleCarry_SetToClear", true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0x3F)
			gbz.flags.Set(Carry, c.initialCarry)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 4)
			assertPC(t, gbz, pc+1)
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADD_A_N8(t *testing.T) {
	cases := []struct {
		name          string
		imm           byte
		initialA      byte
		wantValue     byte
		wantCarry     bool
		wantHalfCarry bool
		wantZero      bool
	}{
		{"ADD_N8_NoFlags", 0x02, 0x01, 0x03, false, false, false},
		{"ADD_N8_HalfCarry", 0x01, 0x0F, 0x10, false, true, false},
		{"ADD_N8_Carry", 0x01, 0xFF, 0x00, true, true, true},
		{"ADD_N8_Zero", 0x00, 0x00, 0x00, false, false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.mem.Write(gbz.pc, 0xC6)
			gbz.mem.Write(gbz.pc+1, c.imm)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_ADD_SP_E8(t *testing.T) {
	cases := []struct {
		name          string
		sp            uint16
		offset        int8
		wantSP        uint16
		wantCarry     bool
		wantHalfCarry bool
	}{
		{"NoFlags", 0x0001, 0x01, 0x0002, false, false},
		{"HalfCarryOnly", 0x000F, 0x01, 0x0010, false, true},
		{"CarryFromFFFF", 0xFFFF, 0x01, 0x0000, true, true},
		{"NegativeOffset", 0x0010, -0x01, 0x000F, false, true},
		{"NegativeCarryAndHalf", 0x0000, -0x01, 0xFFFF, true, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz := createGBZ()
			gbz.sp = c.sp
			gbz.mem.Write(gbz.pc, 0xE8)
			gbz.mem.Write(gbz.pc+1, byte(c.offset))

			gbz.Run()

			assertCycles(t, gbz, 16)
			assertPC(t, gbz, 0x102)
			assert.Equal(t, c.wantSP, gbz.sp, "SP mismatch")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantHalfCarry, gbz.flags.Get(HalfCarry), "HalfCarry flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Zero), "Zero flag should be false")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
		})
	}
}

func Test_RRC_B(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialB  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x08)
			gbz.b = c.initialB
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.b, "B")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_C(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialC  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x09)
			gbz.c = c.initialC
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.c, "C")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_D(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialD  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0A)
			gbz.d = c.initialD
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.d, "D")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_E(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialE  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0B)
			gbz.e = c.initialE
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.e, "E")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_H(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialH  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0C)
			gbz.h = c.initialH
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.h, "H")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_L(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialL  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0D)
			gbz.l = c.initialL
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.l, "L")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_HL(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialHL byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0E)
			hl := uint16(gbz.h)<<8 | uint16(gbz.l)
			gbz.mem.Write(hl, c.initialHL)
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 16)
			assertPC(t, gbz, pc+2)
			assert.Equal(t, c.wantValue, gbz.mem.Read(hl), "Memory at HL mismatch")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}

func Test_RRC_A(t *testing.T) {
	gbz := createGBZ()

	cases := []struct {
		name      string
		initialA  byte
		wantValue byte
		wantCarry bool
		wantZero  bool
	}{
		{"Rotate_0x00_Zero", 0x00, 0x00, false, true},
		{"Rotate_0x01_Carry", 0x01, 0x80, true, false},
		{"Rotate_0x02_NoCarry", 0x02, 0x01, false, false},
		{"Rotate_0x80_NoCarry", 0x80, 0x40, false, false},
		{"Rotate_0xFF_Carry", 0xFF, 0xFF, true, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gbz.mem.Write(gbz.pc, 0xCB)
			gbz.mem.Write(gbz.pc+1, 0x0F)
			gbz.a = c.initialA
			pc := gbz.pc

			gbz.Run()

			assertCycles(t, gbz, 8)
			assertPC(t, gbz, pc+2)
			assertRegister(t, c.wantValue, gbz.a, "A")
			assert.Equal(t, c.wantCarry, gbz.flags.Get(Carry), "Carry flag mismatch")
			assert.Equal(t, c.wantZero, gbz.flags.Get(Zero), "Zero flag mismatch")
			assert.Equal(t, false, gbz.flags.Get(Sub), "Sub flag should be false")
			assert.Equal(t, false, gbz.flags.Get(HalfCarry), "HalfCarry flag should be false")

			resetGBZ(gbz)
		})
	}
}
