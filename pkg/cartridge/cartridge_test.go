package cartridge_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gaoliveira21/gameboy.go/pkg/cartridge"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func createValidROM() []byte {
	return append(make([]byte, 0x104), cartridge.OfficialNintendoLogo...)
}

func createFakeGbFile(dir string, content []byte) (path string) {
	pth := dir + "/fake.gb"
	f, err := os.Create(pth)
	if err != nil {
		panic(err)
	}

	f.Write(content)

	return pth
}

func TestCartridgeFileNotFound(t *testing.T) {
	c, err := cartridge.New("./invalid/path.gb")

	assert.Nil(t, c)
	assert.ErrorIs(t, err, cartridge.CartridgeFileNotFound)
}

func TestCartridgeLogoDoesNotMatch(t *testing.T) {
	path := createFakeGbFile(t.TempDir(), make([]byte, 200))
	c, err := cartridge.New(path)

	assert.Nil(t, c)
	assert.ErrorIs(t, err, cartridge.CartridgeLogoDoesNotMatch)
}

func TestNewCartridge(t *testing.T) {
	path := createFakeGbFile(t.TempDir(), createValidROM())
	c, err := cartridge.New(path)

	assert.NotNil(t, c)
	assert.Nil(t, err)
}

func TestLength(t *testing.T) {
	rom := createValidROM()
	path := createFakeGbFile(t.TempDir(), rom)
	c, _ := cartridge.New(path)

	assert.Equal(t, len(rom), c.Length())
}

func TestTitle(t *testing.T) {
	title := "TETRIS"
	rom := append(createValidROM(), []byte(title)...)
	path := createFakeGbFile(t.TempDir(), rom)
	c, _ := cartridge.New(path)

	assert.Equal(t, title, c.Title())
}

func TestPrintNintendoLogo(t *testing.T) {
	path := createFakeGbFile(t.TempDir(), createValidROM())
	c, _ := cartridge.New(path)
	buf := new(bytes.Buffer)

	c.PrintLogo(buf)

	snaps.MatchSnapshot(t, buf.String())
}

func TestMbcTypeRomOnly(t *testing.T) {
	rom := append(createValidROM(), make([]byte, 50)...)
	rom[0x147] = 0x0
	path := createFakeGbFile(t.TempDir(), rom)
	c, _ := cartridge.New(path)

	assert.Equal(t, cartridge.RomOnly, c.Mbc)
}

func FuzzMbcTypeMbc1(f *testing.F) {
	f.Add(0x1)
	f.Add(0x2)
	f.Add(0x3)

	f.Fuzz(func(t *testing.T, b int) {
		rom := append(createValidROM(), make([]byte, 50)...)
		rom[0x147] = byte(b)
		path := createFakeGbFile(t.TempDir(), rom)
		c, _ := cartridge.New(path)

		assert.Equal(t, cartridge.Mbc1, c.Mbc)
	})
}

func FuzzMbcTypeMbc2(f *testing.F) {
	f.Add(0x5)
	f.Add(0x6)

	f.Fuzz(func(t *testing.T, b int) {
		rom := append(createValidROM(), make([]byte, 50)...)
		rom[0x147] = byte(b)
		path := createFakeGbFile(t.TempDir(), rom)
		c, _ := cartridge.New(path)

		assert.Equal(t, cartridge.Mbc2, c.Mbc)
	})
}
