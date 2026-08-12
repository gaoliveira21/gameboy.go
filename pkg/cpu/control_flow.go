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
