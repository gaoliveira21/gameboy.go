package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	m := &memory{}
	m.Write(0x6, 0x99)

	assert.EqualValues(t, 0x99, m.Read(0x6), "Error reading from memory")
}
