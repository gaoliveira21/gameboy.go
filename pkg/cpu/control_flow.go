package cpu

func (gbz *GBZ80) jumpE8() {
	e8 := int8(gbz.mem.Read(gbz.pc))
	gbz.pc++
	gbz.pc = uint16(int16(gbz.pc) + int16(e8))
}
