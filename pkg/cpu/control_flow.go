package cpu

func (gbz *GBZ80) jumpE8() {
	e8 := int8(gbz.mem.Read(gbz.pc))
	gbz.pc++
	gbz.pc = uint16(int16(gbz.pc) + int16(e8))
}

func (gbz *GBZ80) jumpA16() {
	lb := gbz.mem.Read(gbz.pc)
	hb := gbz.mem.Read(gbz.pc + 1)

	gbz.pc = (uint16(hb) << 8) | uint16(lb)
}

func (gbz *GBZ80) ret() {
	lb := gbz.mem.Read(gbz.sp)
	hb := gbz.mem.Read(gbz.sp + 1)
	gbz.sp += 2

	gbz.pc = (uint16(hb) << 8) | uint16(lb)
}

func (gbz *GBZ80) call() {
	lb := gbz.mem.Read(gbz.pc)
	hb := gbz.mem.Read(gbz.pc + 1)
	gbz.pc += 2

	addr := (uint16(hb) << 8) | uint16(lb)

	gbz.sp--
	gbz.mem.Write(gbz.sp, byte(gbz.pc>>8))
	gbz.sp--
	gbz.mem.Write(gbz.sp, byte(gbz.pc&0xFF))

	gbz.pc = addr
}
