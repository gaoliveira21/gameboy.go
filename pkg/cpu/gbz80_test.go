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

func Test_LD_BC_n16(t *testing.T) {
	gbz := createGBZWithOpcode(0x01)
	pc := gbz.pc
	lb := byte(0x20)
	hb := byte(0x80)

	gbz.mem.Write(pc+1, lb)
	gbz.mem.Write(pc+2, hb)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, pc+3)
	assertRegister(t, lb, gbz.c, "C")
	assertRegister(t, hb, gbz.b, "B")
}

func Test_LD_BC_A(t *testing.T) {
	gbz := createGBZWithOpcode(0x02)
	gbz.b = 0x20
	gbz.c = 0x80
	gbz.a = 0x99
	pc := gbz.pc
	addr := (uint16(gbz.b) << 8) | uint16(gbz.c)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+1)
	assertMemory(t, gbz, addr, gbz.a)
}

func Test_LD_B_n8(t *testing.T) {
	gbz := createGBZWithOpcode(0x06)
	pc := gbz.pc
	expected := byte(0x99)
	gbz.mem.Write(pc+1, expected)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+2)
	assertRegister(t, expected, gbz.b, "B")
}

func Test_LD_n16_SP(t *testing.T) {
	gbz := createGBZWithOpcode(0x08)
	gbz.sp = 0x8020
	pc := gbz.pc

	gbz.Run()

	assertCycles(t, gbz, 20)
	assertPC(t, gbz, pc+3)
	assertMemory(t, gbz, pc+1, byte(gbz.sp&0xFF))
	assertMemory(t, gbz, pc+2, byte(gbz.sp>>8))
}

func Test_LD_A_BC(t *testing.T) {
	gbz := createGBZWithOpcode(0x0A)
	gbz.b = 0x20
	gbz.c = 0x80
	pc := gbz.pc
	addr := (uint16(gbz.b) << 8) | uint16(gbz.c)
	expected := byte(0x99)
	gbz.mem.Write(addr, expected)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+1)
	assertMemory(t, gbz, addr, gbz.a)
	assertRegister(t, expected, gbz.a, "A")
}

func Test_LD_C_n8(t *testing.T) {
	gbz := createGBZWithOpcode(0x0E)
	pc := gbz.pc
	expected := byte(0x99)
	gbz.mem.Write(pc+1, expected)

	gbz.Run()

	assertCycles(t, gbz, 8)
	assertPC(t, gbz, pc+2)
	assertRegister(t, expected, gbz.c, "C")
}

func Test_LD_DE_n16(t *testing.T) {
	gbz := createGBZWithOpcode(0x11)
	pc := gbz.pc
	lb := byte(0x20)
	hb := byte(0x80)

	gbz.mem.Write(pc+1, lb)
	gbz.mem.Write(pc+2, hb)

	gbz.Run()

	assertCycles(t, gbz, 12)
	assertPC(t, gbz, pc+3)
	assertRegister(t, lb, gbz.e, "E")
	assertRegister(t, hb, gbz.d, "D")
}
