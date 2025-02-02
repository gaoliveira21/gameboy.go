package main

import (
	"flag"
	"fmt"

	"github.com/gaoliveira21/gameboy.go/pkg/cartridge"
	"github.com/gaoliveira21/gameboy.go/pkg/cpu"
	"github.com/gaoliveira21/gameboy.go/pkg/graphics"
	"github.com/gaoliveira21/gameboy.go/pkg/timer"
)

const maxCycles = 69905

func Run() {
	cpu := cpu.NewGBZ80()
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

	fmt.Printf("%d bytes read", c.Length())
}
