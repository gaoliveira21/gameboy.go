package cpu

func (gbz *GBZ80) jumpE8() {
	e8 := int8(gbz.mem.Read(gbz.pc))
	gbz.pc++
	addr := int32(gbz.pc) + int32(e8)
	gbz.pc = uint16(addr)
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

func (gbz *GBZ80) rst(addr uint16) {
	gbz.sp--
	gbz.mem.Write(gbz.sp, byte(gbz.pc>>8))
	gbz.sp--
	gbz.mem.Write(gbz.sp, byte(gbz.pc&0xFF))

	gbz.pc = addr
}

func (gbz *GBZ80) pop() (hb byte, lb byte) {
	lob, hib := gbz.mem.Read(gbz.sp), gbz.mem.Read(gbz.sp+1)
	gbz.sp += 2

	return hib, lob
}

func (gbz *GBZ80) push(hb byte, lb byte) {
	gbz.sp--
	gbz.mem.Write(gbz.sp, hb)
	gbz.sp--
	gbz.mem.Write(gbz.sp, lb)
}
