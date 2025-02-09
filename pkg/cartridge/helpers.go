package cartridge

import "os"

func CreateValidCartridgeForTest(dir string) *Cartridge {
	path := dir + "valid_file.gb"
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	rom := append(make([]byte, 0x104), OfficialNintendoLogo...)
	rom = append(rom, make([]byte, 0x8000-len(rom))...)
	f.Write(rom)

	c, err := New(path)
	if err != nil {
		panic(err)
	}

	return c
}
