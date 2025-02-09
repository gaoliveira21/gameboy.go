package memory

import "github.com/gaoliveira21/gameboy.go/pkg/cartridge"

type Memory struct {
	mem       [0x10000]byte
	cartridge *cartridge.Cartridge
}

func NewMemory(c *cartridge.Cartridge) *Memory {
	return &Memory{
		cartridge: c,
	}
}

func (m *Memory) Read(addr uint16) byte {
	// Reading from ROM memory bank
	if addr >= 0x4000 && addr <= 0x7FFF {
		return m.cartridge.Memory[(addr-0x4000)+uint16(m.cartridge.CurrentROMBank*0x4000)]
	}

	// Reading from RAM memory bank
	if addr >= 0xA000 && addr <= 0xBFFF {
		return m.cartridge.RamBanks[(addr-0xA000)+uint16(m.cartridge.CurrentRamBank*0x2000)]
	}

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
