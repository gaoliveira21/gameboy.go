package cpu_test

import (
	"testing"

	"github.com/gaoliveira21/gameboy.go/pkg/cpu"
	"github.com/stretchr/testify/assert"
)

func TestSetFlagsToTrue(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, true)
	flags.Set(cpu.Sub, true)
	flags.Set(cpu.HalfCarry, true)
	flags.Set(cpu.Carry, true)

	assert.Equal(t, true, flags.Get(cpu.Zero), "cpu.Zero flag is not set")
	assert.Equal(t, true, flags.Get(cpu.Sub), "cpu.Sub flag is not set")
	assert.Equal(t, true, flags.Get(cpu.HalfCarry), "cpu.HalfCarry flag is not set")
	assert.Equal(t, true, flags.Get(cpu.Carry), "cpu.Carry flag is not set")
}

func TestSetFlagsToFalse(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, false)
	flags.Set(cpu.Sub, false)
	flags.Set(cpu.HalfCarry, false)
	flags.Set(cpu.Carry, false)

	assert.Equal(t, false, flags.Get(cpu.Zero), "cpu.Zero flag is set")
	assert.Equal(t, false, flags.Get(cpu.Sub), "cpu.Sub flag is set")
	assert.Equal(t, false, flags.Get(cpu.HalfCarry), "cpu.HalfCarry flag is set")
	assert.Equal(t, false, flags.Get(cpu.Carry), "cpu.Carry flag is set")
}

func TestSetFlagsToMixedValues(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, false)
	flags.Set(cpu.Sub, true)
	flags.Set(cpu.HalfCarry, false)
	flags.Set(cpu.Carry, true)

	assert.Equal(t, false, flags.Get(cpu.Zero), "cpu.Zero flag is set")
	assert.Equal(t, true, flags.Get(cpu.Sub), "cpu.Sub flag is not set")
	assert.Equal(t, false, flags.Get(cpu.HalfCarry), "cpu.HalfCarry flag is set")
	assert.Equal(t, true, flags.Get(cpu.Carry), "cpu.Carry flag is not set")
}

func TestValue(t *testing.T) {
	flags := &cpu.Flags{}

	if flags.Value() != 0 {
		t.Errorf("cpu.Value is not starting with 0")
	}

	flags.Set(cpu.Zero, true)

	assert.EqualValues(t, 128, flags.Value(), "cpu.Value do not return correctly")
}
