package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gaoliveira21/gameboy.go/pkg/cartridge"
	"github.com/gaoliveira21/gameboy.go/pkg/cpu"
	"github.com/gaoliveira21/gameboy.go/pkg/graphics"
	"github.com/gaoliveira21/gameboy.go/pkg/memory"
	"github.com/gaoliveira21/gameboy.go/pkg/timer"
)

const maxCycles = 69905

func Run(c *cartridge.Cartridge) {
	memory := memory.NewMemory(c)
	cpu := cpu.NewGBZ80(memory)
	var cycles uint = 0

	for cycles < maxCycles {
		cycles += cpu.Run()
		timer.Update(cycles)
		graphics.Update(cycles)
		cpu.Interrupt()
	}

	graphics.Render()
}

func main() {
	rPath := flag.String("rom", "", "Path to ROM file")

	flag.Parse()

	c, err := cartridge.New(*rPath)
	if err != nil {
		panic(err)
	}

	c.PrintLogo(os.Stdout)
	os.Stdout.WriteString("ROM Loaded successfully!!\n")
	os.Stdout.WriteString("_______________________\n")
	os.Stdout.WriteString(fmt.Sprintf("[Title]=%s\n", c.Title()))
	os.Stdout.WriteString(fmt.Sprintf("[Size]=%d bytes\n", c.Length()))
}
