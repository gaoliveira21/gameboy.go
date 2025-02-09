package cartridge

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

const maxMem uint = 0x200000

var OfficialNintendoLogo = []byte{
	0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
	0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
	0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E}

var CartridgeLogoDoesNotMatch = errors.New("cartridge::New: Nintendo logo in this cartridge does not match the original nintendo logo")
var CartridgeFileNotFound = errors.New("cartridge::New: Error opening file")

type MbcType = int

const (
	RomOnly MbcType = iota
	Mbc1    MbcType = iota
	Mbc2    MbcType = iota
)

type Cartridge struct {
	Memory [maxMem]byte
	len    int
	Mbc    MbcType

	CurrentROMBank int

	RamBanks       []byte
	CurrentRamBank int
}

func New(path string) (*Cartridge, error) {
	rom, err := os.ReadFile(path)
	if err != nil {
		return nil, CartridgeFileNotFound
	}

	c := &Cartridge{
		len:            len(rom),
		CurrentROMBank: 1,
		RamBanks:       make([]byte, 0x8000),
		CurrentRamBank: 0,
	}

	for i, v := range rom {
		c.Memory[i] = v
	}

	err = c.validate()
	if err != nil {
		fmt.Println("The selected cartridge is invalid")
		return nil, err
	}

	c.setCartridgeType()

	return c, nil
}

func (c *Cartridge) Length() int {
	return c.len
}

func (c *Cartridge) PrintLogo(w io.StringWriter) {
	matrix := [8][12]string{}
	halfTop := OfficialNintendoLogo[:24]
	halfBottom := OfficialNintendoLogo[24:]

	c.drawLogo(halfTop, &matrix, 0)
	c.drawLogo(halfBottom, &matrix, 4)

	w.WriteString("\n\n")
	for _, row := range matrix {
		for _, item := range row {
			w.WriteString(string(item))
		}
		w.WriteString("\n")
	}
	w.WriteString("\n\n")
}

func (c *Cartridge) Title() string {
	start := 0x0134
	end := 0x0144
	title := []byte{}

	for i := start; i < end; i++ {
		if c.Memory[i] == 0x0 {
			break
		}
		title = append(title, c.Memory[i])
	}

	return string(title)
}

func (c *Cartridge) validate() error {
	cartridgeLogo := c.getNintendoLogo()

	if slices.Equal(cartridgeLogo, OfficialNintendoLogo) {
		return nil
	}

	return CartridgeLogoDoesNotMatch
}

func (c *Cartridge) getNintendoLogo() []byte {
	start := 0x104
	end := 0x133

	cartridgeLogo := c.Memory[start : end+1]

	return cartridgeLogo
}

func (c *Cartridge) drawLogo(lg []byte, matrix *[8][12]string, startRow int) {
	row := startRow
	col := 0
	for i, v := range lg {
		if i%2 == 0 {
			row = startRow
		}

		tile := ""
		for bit := 7; bit >= 0; bit-- {
			pixel := (v >> bit) & 0x1
			if pixel == 1 {
				tile += "XX"
			} else {
				tile += "  "
			}

			if bit == 4 || bit == 0 {
				matrix[row][col] = tile
				tile = ""
				row++
			}
		}

		if i%2 != 0 {
			col++
		}
	}
}

func (c *Cartridge) setCartridgeType() {
	switch c.Memory[0x147] {
	case 0x0:
		c.Mbc = RomOnly
	case 0x1, 0x2, 0x3:
		c.Mbc = Mbc1
	case 0x5, 0x6:
		c.Mbc = Mbc2
	}
}
