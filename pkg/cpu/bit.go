package cpu

func (gbz *GBZ80) rlcr8(r8 *uint8) {
	r := *r8

	*r8 = *r8<<1 | *r8>>7

	gbz.flags.Set(Zero, *r8 == 0)
	gbz.flags.Set(Sub, false)
	gbz.flags.Set(HalfCarry, false)
	gbz.flags.Set(Carry, r > 0x7F)
}
