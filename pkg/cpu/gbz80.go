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

	InterruptEnabled bool
	Halt             bool
	IsStopped        bool

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

	g.initInstructionSet()
	g.boot()

	return g
}

func (gbz *GBZ80) Run() uint {
	opcode := gbz.fetch()
	instr := gbz.decode(opcode)
	cycles := gbz.exec(instr)

	gbz.cycles += cycles

	return cycles
}

func (gbz *GBZ80) fetch() byte {
	opcode := gbz.mem.Read(gbz.pc)
	gbz.pc++
	return opcode
}

func (gbz *GBZ80) decode(opcode byte) *instruction {
	return gbz.instructionSet[opcode]
}

func (gbz *GBZ80) exec(instr *instruction) uint {
	return instr.operation()
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
