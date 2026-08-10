package cpu

func (gbz *GBZ80) ldr8n8(r *uint8) {
	n := gbz.mem.Read(gbz.pc)
	gbz.pc++

	*r = n
}

func (gbz *GBZ80) ldrpn16(rh, rl *uint8) {
	lb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	hb := gbz.mem.Read(gbz.pc)
	gbz.pc++

	*rh, *rl = hb, lb
}

func (gbz *GBZ80) ldrpr8(rh, rl, rt *uint8) {
	hb := uint16(*rh)
	lb := uint16(*rl)

	addr := (hb << 8) | lb
	gbz.mem.Write(addr, *rt)
}

func (gbz *GBZ80) ldr8rp(rt, rh, rl *uint8) {
	hb := uint16(*rh)
	lb := uint16(*rl)

	addr := (hb << 8) | lb
	*rt = gbz.mem.Read(addr)
}
