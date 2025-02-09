package cartridge_test

import (
	"os"
	"testing"

	"github.com/gaoliveira21/gameboy.go/pkg/cartridge"
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

	assert.Equal(t, c.Length(), len(rom))
}

func TestTitle(t *testing.T) {
	title := "TETRIS"
	rom := append(createValidROM(), []byte(title)...)
	path := createFakeGbFile(t.TempDir(), rom)

	c, _ := cartridge.New(path)

	assert.Equal(t, c.Title(), title)
}
