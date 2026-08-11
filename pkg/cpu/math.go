package cpu

func (gbz *GBZ80) incr16(rh, rl *uint8) {
	r16 := (uint16(*rh) << 8) | uint16(*rl)
	r16++

	*rh = uint8(r16 >> 8)
	*rl = uint8(r16 & 0xFF)
}

func (gbz *GBZ80) decr16(rh, rl *uint8) {
	r16 := (uint16(*rh) << 8) | uint16(*rl)
	r16--

	*rh = uint8(r16 >> 8)
	*rl = uint8(r16 & 0xFF)
}

func (gbz *GBZ80) incr8(r *uint8) {
	orig := *r
	*r++

	gbz.flags.Set(Zero, *r == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, orig&0x0F == 0x0F)
}

func (gbz *GBZ80) decr8(r *uint8) {
	orig := *r
	*r--

	gbz.flags.Set(Zero, *r == 0)
	gbz.flags.Set(Sub, true)
	gbz.flags.Set(HalfCarry, orig&0x0F == 0)
}
