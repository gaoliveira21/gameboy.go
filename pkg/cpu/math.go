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

func (gbz *GBZ80) add(v byte, carry uint16) {
	r := uint16(gbz.a) + uint16(v) + carry

	gbz.flags.Set(Carry, r > 0xFF)
	gbz.flags.Set(HalfCarry, ((gbz.a^uint8(r)^v)&0x10) > 0)

	gbz.a = uint8(r & 0xFF)

	gbz.flags.Set(Sub, false)
	gbz.flags.Set(Zero, gbz.a == 0)
}

func (gbz *GBZ80) adc(v byte) {
	cv := uint16(0)
	if gbz.flags.Get(Carry) {
		cv = 1
	}

	gbz.add(v, cv)
}

func (gbz *GBZ80) sub(v byte, carry uint16) {
	r := uint16(gbz.a) - uint16(v) - carry

	gbz.flags.Set(Carry, r>>8 > 0)
	gbz.flags.Set(HalfCarry, ((gbz.a^uint8(r)^v)&0x10) > 0)

	gbz.a = uint8(r & 0xFF)

	gbz.flags.Set(Sub, true)
	gbz.flags.Set(Zero, gbz.a == 0)
}

func (gbz *GBZ80) sbc(v byte) {
	cv := uint16(0)
	if gbz.flags.Get(Carry) {
		cv = 1
	}

	gbz.sub(v, cv)
}

func (gbz *GBZ80) and(v byte) {
	gbz.a &= v

	gbz.flags.Set(Zero, gbz.a == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, true)
	gbz.flags.Set(Carry, false)
}

func (gbz *GBZ80) xor(v byte) {
	gbz.a ^= v

	gbz.flags.Set(Zero, gbz.a == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, false)
}

func (gbz *GBZ80) or(v byte) {
	gbz.a |= v

	gbz.flags.Set(Zero, gbz.a == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, false)
}

func (gbz *GBZ80) cp(v byte) {
	r := uint16(gbz.a) - uint16(v)

	gbz.flags.Set(Carry, r>>8 > 0)
	gbz.flags.Set(HalfCarry, ((gbz.a^uint8(r)^v)&0x10) > 0)
	gbz.flags.Set(Zero, uint8(r) == 0)
	gbz.flags.Set(Sub, true)
}
