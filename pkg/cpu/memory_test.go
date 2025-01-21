package cpu

import (
	"testing"
)

func TestMemory(t *testing.T) {
	m := &memory{}
	m.Write(0x6, 0x99)
	if m.Read(0x6) != 0x99 {
		t.Errorf("Error reading from memory; expected %d; got %d", 0x99, m.Read(0x6))
	}
}
