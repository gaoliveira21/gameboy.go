package cpu

type Flag = byte

const (
	Zero      Flag = 1 << 7
	Sub       Flag = 1 << 6
	HalfCarry Flag = 1 << 5
	Carry     Flag = 1 << 4
)

type Flags struct {
	value byte
}

func (flags *Flags) Set(f Flag, v bool) {
	if v {
		flags.value |= f
	} else {
		flags.value &= ^f
	}
}

func (flags *Flags) Get(f Flag) bool {
	return flags.value&f != 0
}
