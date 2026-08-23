package cpu

func (gbz *GBZ80) rlc(r8 *uint8) {
	r := *r8

	*r8 = *r8<<1 | *r8>>7

	gbz.flags.Set(Zero, *r8 == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, r > 0x7F)
}

func (gbz *GBZ80) rrc(r8 *uint8) {
	r := *r8
	*r8 = r>>1 | (r&0x1)<<7

	gbz.flags.Set(Zero, *r8 == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, r&0x1 == 1)
}

func (gbz *GBZ80) rl(r8 *uint8) {
	c := uint8(0)
	if gbz.flags.Get(Carry) {
		c = 1
	}

	r := *r8
	*r8 = *r8<<1 | c

	gbz.flags.Set(Zero, *r8 == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, r>>7 == 1)
}

func (gbz *GBZ80) rr(r8 *uint8) {
	c := uint8(0)
	if gbz.flags.Get(Carry) {
		c = 1
	}

	r := *r8
	*r8 = c<<7 | *r8>>1

	gbz.flags.Set(Zero, *r8 == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, r&0x1 == 1)
}
