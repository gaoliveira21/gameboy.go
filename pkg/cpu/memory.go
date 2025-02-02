package cpu

type memory struct {
	mem [0x10000]byte
}

func (m *memory) Read(i uint16) byte {
	return m.mem[i]
}

func (m *memory) Write(i uint16, b byte) {
	m.mem[i] = b
}
