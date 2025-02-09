package memory

type Memory struct {
	mem [0x10000]byte
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Read(addr uint16) byte {
	return m.mem[addr]
}

func (m *Memory) Write(addr uint16, b byte) {
	// ROM Bank
	if addr <= 0x7FFF {
		return
	}

	// Echo RAM
	if addr >= 0xE000 && addr <= 0xFDFF {
		m.mem[addr] = b
		m.Write(addr-0x2000, b)
		return
	}

	// Not Usable
	if addr >= 0xFEA0 && addr <= 0xFEFF {
		return
	}

	m.mem[addr] = b
}
