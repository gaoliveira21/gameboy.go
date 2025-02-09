package memory_test

import (
	"testing"

	"github.com/gaoliveira21/gameboy.go/pkg/cartridge"
	"github.com/gaoliveira21/gameboy.go/pkg/memory"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	m := memory.NewMemory(cartridge.CreateValidCartridgeForTest(t.TempDir()))
	m.Write(0x8000, 0x99)

	assert.EqualValues(t, 0x99, m.Read(0x8000), "Error reading/writing value on memory")
}

func TestEchoRAM(t *testing.T) {
	m := memory.NewMemory(cartridge.CreateValidCartridgeForTest(t.TempDir()))
	m.Write(0xE000, 0x99)
	m.Write(0xE002, 0x88)
	m.Write(0xFDFF, 0x77)

	assert.EqualValues(t, 0x99, m.Read(0xE000), "Error reading/writing value on memory")
	assert.EqualValues(t, 0x99, m.Read(0xC000), "Writing to Echo RAM did not work properly")
	assert.EqualValues(t, 0x88, m.Read(0xE002), "Error reading/writing value on memory")
	assert.EqualValues(t, 0x88, m.Read(0xC002), "Writing to Echo RAM did not work properly")
	assert.EqualValues(t, 0x77, m.Read(0xFDFF), "Error reading/writing value on memory")
	assert.EqualValues(t, 0x77, m.Read(0xDDFF), "Writing to Echo RAM did not work properly")
}

func TestNotUsableRAM(t *testing.T) {
	m := memory.NewMemory(cartridge.CreateValidCartridgeForTest(t.TempDir()))
	m.Write(0xFEA0, 0x99)
	m.Write(0xFEA1, 0x88)
	m.Write(0xFEFF, 0x77)

	assert.EqualValues(t, 0x00, m.Read(0xFEA0), "Not Usable RAM has been written")
	assert.EqualValues(t, 0x00, m.Read(0xFEA1), "Not Usable RAM has been written")
	assert.EqualValues(t, 0x00, m.Read(0xFEFF), "Not Usable RAM has been written")
}
