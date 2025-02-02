package cartridge

import (
	"fmt"
	"os"
)

const maxMem uint = 0x200000

type Cartridge struct {
	memory [maxMem]byte
	len    int
}

func New(path string) (*Cartridge, error) {
	rom, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cartridge::New: Error opening file %s", err.Error())
	}

	c := &Cartridge{
		len: len(rom),
	}

	for i, v := range rom {
		c.memory[i] = v
	}

	return c, nil
}

func (c *Cartridge) Length() int {
	return c.len
}
