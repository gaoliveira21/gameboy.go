package cpu_test

import (
	"testing"

	"github.com/gaoliveira21/gameboy.go/pkg/cpu"
)

func TestSetFlagsToTrue(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, true)
	flags.Set(cpu.Sub, true)
	flags.Set(cpu.HalfCarry, true)
	flags.Set(cpu.Carry, true)

	if flags.Get(cpu.Zero) != true {
		t.Errorf("cpu.Zero flag is not set")
	}

	if flags.Get(cpu.Sub) != true {
		t.Errorf("cpu.Sub flag is not set")
	}

	if flags.Get(cpu.HalfCarry) != true {
		t.Errorf("cpu.HalfCarry flag is not set")
	}

	if flags.Get(cpu.Carry) != true {
		t.Errorf("cpu.Carry flag is not set")
	}
}

func TestSetFlagsToFalse(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, false)
	flags.Set(cpu.Sub, false)
	flags.Set(cpu.HalfCarry, false)
	flags.Set(cpu.Carry, false)

	if flags.Get(cpu.Zero) != false {
		t.Errorf("cpu.Zero flag is not set")
	}

	if flags.Get(cpu.Sub) != false {
		t.Errorf("cpu.Sub flag is not set")
	}

	if flags.Get(cpu.HalfCarry) != false {
		t.Errorf("cpu.HalfCarry flag is not set")
	}

	if flags.Get(cpu.Carry) != false {
		t.Errorf("cpu.Carry flag is not set")
	}
}

func TestSetFlagsToMixedValues(t *testing.T) {
	flags := &cpu.Flags{}

	flags.Set(cpu.Zero, false)
	flags.Set(cpu.Sub, true)
	flags.Set(cpu.HalfCarry, false)
	flags.Set(cpu.Carry, true)

	if flags.Get(cpu.Zero) != false {
		t.Errorf("cpu.Zero flag is not set")
	}

	if flags.Get(cpu.Sub) != true {
		t.Errorf("cpu.Sub flag is not set")
	}

	if flags.Get(cpu.HalfCarry) != false {
		t.Errorf("cpu.HalfCarry flag is not set")
	}

	if flags.Get(cpu.Carry) != true {
		t.Errorf("cpu.Carry flag is not set")
	}
}
