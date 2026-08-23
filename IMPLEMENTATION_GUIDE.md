# Game Boy Emulator Implementation Guide

## Table of Contents
1. [Understanding the Game Boy Hardware](#understanding-the-game-boy-hardware)
2. [Phase 1: CPU Opcodes](#phase-1-cpu-opcodes)
3. [Phase 2: Memory/MMU Improvements](#phase-2-memorymmu-improvements)
4. [Phase 3: Timer Implementation](#phase-3-timer-implementation)
5. [Phase 4: PPU/Graphics](#phase-4-ppugraphics)
6. [Phase 5: Interrupt System](#phase-5-interrupt-system)
7. [Phase 6: Input](#phase-6-input)
8. [Phase 7: Main Loop Integration](#phase-7-main-loop-integration)
9. [Testing Your Implementation](#testing-your-implementation)

---

## Understanding the Game Boy Hardware

### What is a Game Boy?

The Game Boy is a portable gaming console released by Nintendo in 1989. It features:
- **Custom 8-bit CPU** (called GBZ80 or Sharp LR35902) running at ~4.19 MHz
- **64KB addressable memory** (all mapped in a single linear space)
- **Custom graphics processor (PPU)** for sprite and background rendering
- **Timer hardware** for game timing
- **Input handling** for controllers

### Memory Map (Understanding the 64KB Space)

Think of the Game Boy's memory as a single street with 65,536 houses (each holding 1 byte), numbered from 0 to 65535. Each "house" has a specific purpose:

| Address Range | Size | Name | What it is for |
|--------------|------|------|----------------|
| `0x0000` - `0x7FFF` | 32KB | ROM Banks | Your game cartridge - the actual game code! |
| `0x8000` - `0x9FFF` | 8KB | Video RAM (VRAM) | Graphics data that the PPU reads to draw the screen |
| `0xA000` - `0xBFFF` | 8KB | External RAM | Save games (if cartridge has battery backup) |
| `0xC000` - `0xCFFF` | 4KB | Work RAM (Bank 0) | General purpose memory for the CPU |
| `0xD000` - `0xDFFF` | 4KB | Work RAM (Bank 1) | More working memory |
| `0xE000` - `0xFDFF` | 8KB | Echo RAM | A copy of work RAM - avoid using it |
| `0xFE00` - `0xFE9F` | 160B | OAM | Sprite information - tells the PPU where to draw 40 sprites |
| `0xFEA0` - `0xFEFF` | 96B | Not Usable | Protected area - can't use it |
| `0xFF00` - `0xFF7F` | 128B | I/O Registers | Hardware control - buttons, timer, LCD settings |
| `0xFF80` - `0xFFFF` | 128B | High RAM (HRAM) | Fast memory, often used for interrupt handlers |

### The CPU (GBZ80)

Your CPU has these "hands" (registers) to work with:

**8-bit registers** (hold single bytes):
```
A  - The "accumulator" - main register for math operations
B, C, D, E, H, L - General purpose registers
F  - Flags register - tells you about the last operation (Z=zero, C=carry, etc.)
```

**16-bit registers** (hold two bytes):
```
PC - Program Counter - points to the NEXT instruction to execute (like a bookmark)
SP - Stack Pointer - points to where the stack is in memory
```

When the CPU runs a game:
1. It looks at the address in PC
2. Reads the byte at that address (the opcode)
3. Executes the instruction (might read more bytes as arguments)
4. Advances PC to the next instruction
5. Returns result (like changing a register, writing to memory, or jumping)

### The PPU (Picture Processing Unit)

The PPU is like an artist that draws the screen. It works in **scanlines** (rows):

1. **144 visible rows** (line 0 to line 143)
2. **52 invisible rows** during "V-Blank" (line 144 to 153)

Each row takes exactly **456 clock cycles** to draw. The PPU has different "modes" as it draws:

| Mode | Cycles | What it does |
|------|--------|--------------|
| Mode 2 | 80 | Looking at sprite data (OAM) to find what sprites appear on this row |
| Mode 3 | ~168-291 | Actually rendering pixels to the screen |
| Mode 0 | Remaining | Horizontal blank - time to do other things |
| Mode 1 (V-Blank) | 4560 total | Vertical blank - safe time to update graphics |

**How tiles work:**
- The Game Boy doesn't draw individual pixels directly
- It uses **8x8 pixel tiles** (small 8x8 pictures)
- Each tile uses 16 bytes (2 bits per pixel, 8 pixels wide, 8 rows)
- Tiles are stored in VRAM at `0x8000-0x97FF`

**How background works:**
- A background is a 32x32 grid of tile numbers
- Two maps exist: `0x9800-0x9BFF` and `0x9C00-0x9FFF`
- Each byte in the map tells which tile to show

### Timer

The timer is a separate clock that:
- **DIV** register counts up constantly (at ~16384 Hz)
- **TIMA** is the actual timer that triggers interrupts
- **TMA** is the "reload value" - when TIMA overflows, it resets to TMA

### Interrupts

Interrupts are like phone calls that pause the CPU to handle important events:

| Interrupt | Address | When it fires |
|-----------|---------|---------------|
| V-Blank | `0x40` | After all 144 visible scanlines are done |
| Timer | `0x48` | When TIMA overflows |
| LCD STAT | `0x50` | Based on PPU mode conditions |
| Joypad | `0x58` | When a button is pressed |

The CPU has an "IME" flag (interrupt master enable) - when on, interrupts are checked.

---

## Phase 1: CPU Opcodes ✅ COMPLETED

### What You Need to Do

All opcodes are fully implemented and tested.

### Implementation Summary

| Component | Status |
|-----------|--------|
| All 256 opcodes (0x00-0xFF) | ✅ Implemented |
| All 256 CB-prefixed opcodes (0xCBxx) | ✅ Implemented |
| Flag handling | ✅ Implemented |
| Control flow (JP, JR, CALL, RET, RST) | ✅ Implemented |
| Unit tests | ✅ All passing |

### File Structure

| File | Purpose |
|------|---------|
| `pkg/cpu/gbz80.go` | Main CPU struct, fetch/decode/exec cycle, boot ROM |
| `pkg/cpu/instruction_set.go` | All 256 main opcodes |
| `pkg/cpu/prefix.go` | CB-prefixed opcodes |
| `pkg/cpu/load.go` | Load instruction helpers |
| `pkg/cpu/math.go` | Math operations |
| `pkg/cpu/bit.go` | Bit operations |
| `pkg/cpu/control_flow.go` | Jump, call, ret, push, pop |
| `pkg/cpu/flags.go` | Flag register handling |
| `pkg/cpu/gbz80_test.go` | Comprehensive unit tests |

---

## Phase 2: Memory/MMU Improvements

### What You Need to Do

Your current memory implementation is basic. Improve it to handle:
1. MBC ROM banking (switching ROM banks)
2. Proper I/O register handling
3. DMA transfers
4. Proper cartridge ROM/RAM mapping

### Step 1: Implement MBC1 Banking

For MBC1 cartridges (type 0x01, 0x02, 0x03):

**Reading from 0x4000-0x7FFF (switchable ROM):**
```go
if addr >= 0x4000 && addr <= 0x7FFF {
    bank := gbz.cartridge.CurrentROMBank
    if bank == 0 { bank = 1 } // bank 0 is illegal, maps to 1
    offset := (uint32(bank) * 0x4000) + uint32(addr-0x4000)
    return gbz.cartridge.Memory[offset]
}
```

**Writing to 0x6000-0x7FFF (ROM bank number):**
When you write to this range, the lower 5 bits become the new ROM bank number.

```go
if addr >= 0x6000 && addr <= 0x7FFF {
    gbz.cartridge.CurrentROMBank &= 0xE0 // clear lower 5 bits
    gbz.cartridge.CurrentROMBank |= (value & 0x1F)
}
```

**Writing to 0x2000-0x3FFF (low bank number):**
```go
if addr >= 0x2000 && addr <= 0x3FFF {
    bank := value & 0x1F
    if bank == 0 { bank = 1 }
    gbz.cartridge.CurrentROMBank = (gbz.cartridge.CurrentROMBank & 0xE0) | bank
}
```

### Step 2: Implement RAM Enable (MBC1)

For addresses 0x0000-0x1FFF:
- Write 0x0A to enable external RAM
- Write 0x00 to disable external RAM

### Step 3: Implement DMA Transfer (0xFF46)

When writing to DMA:
1. Read source address from bits [7:4] × 0x100
2. Copy 160 bytes from source to OAM (0xFE00-0xFE9F)
3. This takes 160 cycles (CPU is halted during DMA)

### Step 4: Handle Special Registers

In your Read/Write functions, handle these specially:

```go
func (m *Memory) Read(addr uint16) byte {
    switch {
    case addr < 0x8000:
        // ROM - handle banking
    case addr >= 0x8000 && addr < 0xA000:
        // VRAM
    case addr >= 0xA000 && addr < 0xC000:
        // External RAM - handle banking
    case addr >= 0xC000 && addr < 0xE000:
        // Work RAM
    case addr >= 0xE000 && addr < 0xFE00:
        // Echo RAM
    case addr >= 0xFE00 && addr < 0xFEA0:
        // OAM
    case addr == 0xFF00:
        // P1 - Joypad (handle bit 4-5 for button selection)
    case addr == 0xFF04:
        // DIV - just return mem[DIV]
    case addr == 0xFF05:
        // TIMA
    case addr == 0xFF06:
        // TMA
    case addr == 0xFF07:
        // TAC
    case addr == 0xFF0F:
        // IF
    case addr >= 0xFF10 && addr < 0xFF26:
        // Sound registers (can ignore for Tetris)
    case addr == 0xFF40:
        // LCDC
    case addr == 0xFF41:
        // STAT
    case addr == 0xFF44:
        // LY (read only, updated by PPU)
    case addr == 0xFF46:
        // DMA
    case addr == 0xFF47:
        // BGP
    case addr >= 0xFF80:
        // HRAM
    default:
        return m.mem[addr]
    }
}
```

---

## Phase 3: Timer Implementation

### What You Need to Do

Update `pkg/timer/timer.go` to properly track time and generate interrupts.

### Timer Registers

| Register | Address | Description |
|----------|---------|-------------|
| DIV | 0xFF04 | Divider register - always counting |
| TIMA | 0xFF05 | Timer counter |
| TMA | 0xFF06 | Timer modulo (reload value) |
| TAC | 0xFF07 | Timer control |

### Step 1: Create Timer State

```go
type Timer struct {
    div     uint8  // counter that increments at 16384 Hz
    tima    uint8  // timer counter
    tma     uint8  // modulo
    tac     uint8  // control
    // internal
    counter  uint   // tracks when to increment
}
```

### Step 2: Implement Divider

- DIV increments at a rate of `16384 Hz` (CPU clock / 256)
- Writing to DIV always resets it to 0x00

### Step 3: Implement TIMA

- TIMA increments based on TAC's clock select bits:
  - Bit 0-1 of TAC select the clock
  - Bit 2 of TAC enables/disables the timer

| TAC bits 1-0 | Rate |
|--------------|------|
| 00 | 4096 Hz |
| 01 | 262144 Hz |
| 10 | 65536 Hz |
| 11 | 16384 Hz |

- When TIMA overflows (0xFF → 0x00):
  - Set bit 2 of IF register (timer interrupt request)
  - Reload TIMA with TMA value

### Step 4: Integration with CPU

In your main loop, after each instruction:
```go
timer.Update(cycles) // add cycles to timer
```

The timer should check if enough cycles have passed to increment its internal counters.

---

## Phase 4: PPU/Graphics

### What You Need to Do

Rewrite `pkg/graphics/graphics.go` to implement the PPU. This is the most complex part.

### Step 1: Create PPU State

```go
type PPU struct {
    mem *Memory

    // Video RAM (0x8000-0x9FFF)
    vram [0x2000]byte

    // OAM (0xFE00-0xFE9F)
    oam [0xA0]byte

    // LCD Control (0xFF40)
    lcdc uint8

    // Current scanline (0-153)
    ly uint8

    // Scroll positions
    scx uint8  // SCX register (0xFF43)
    scy uint8  // SCY register (0xFF42)

    // Window position
    wy uint8   // Window Y (0xFF4A)
    wx uint8   // Window X - 7 (0xFF4B)

    // Background palette (0xFF47)
    bgp uint8

    // PPU mode (0=HB, 1=VB, 2=OAM, 3=Drawing)
    mode uint8

    // Cycles in current scanline
    scanlineCycles uint

    // Framebuffer (160x144 pixels)
    pixels [160*144]uint8
}
```

### Step 2: Understand Scanline Timing

Each scanline takes 456 cycles:

| Mode | Duration | What happens |
|------|----------|--------------|
| Mode 2 (OAM) | 80 cycles | PPU reads sprite info from OAM |
| Mode 3 (Drawing) | ~168 cycles | PPU renders pixels |
| Mode 0 (H-Blank) | Remaining | Safe to write VRAM |

After 144 scanlines, Mode 1 (V-Blank) lasts 10 scanlines (4560 cycles).

### Step 3: Implement OAM Search

In Mode 2, read all 40 sprite entries from OAM. For each sprite:
- Check if sprite is on current scanline (Y position)
- Sprite is 8x8 or 8x16 (based on LCDC bit 2)
- Store sprites visible on this line for rendering

### Step 4: Implement Pixel Rendering

In Mode 3, actually draw pixels:

**Background rendering:**
1. Get current X coordinate (0-159)
2. Add SCX scroll offset
3. Get tile number from background map
4. Get pixel from tile data

**Sprite rendering:**
- Check if sprite covers this pixel
- If sprite pixel is not transparent (color != 0), draw sprite over background
- Handle sprite palette (OBJ palette 0 or 1)

### Step 5: Handle VRAM Access

During mode 3, the CPU cannot access VRAM. Your memory code should handle this:
- If PPU is in mode 3 and CPU tries to read VRAM, return 0xFF
- If PPU is in mode 3 and CPU tries to write VRAM, ignore the write

### Step 6: Implement STAT Interrupts

STAT register (0xFF41) has bits for:
- LYC=LY coincidence interrupt (bit 6)
- OAM interrupt (bit 5)
- V-Blank interrupt (bit 4)
- H-Blank interrupt (bit 3)

Set the appropriate bits in IF register when these conditions occur.

### Step 7: Generate V-Blank Interrupt

When moving from scanline 143 to 144:
- Set mode to 1 (V-Blank)
- Set bit 0 of IF register (V-Blank interrupt)
- This is the signal for the game to update graphics

---

## Phase 5: Interrupt System

### What You Need to Do

Implement the CPU's interrupt handling. The `InterruptEnabled` flag and EI/DI opcodes exist, but the actual `Interrupt()` method needs to be implemented.

### Step 1: Implement Interrupt() Method

In `pkg/cpu/gbz80.go`, implement the `Interrupt()` method that:

1. Checks if `InterruptEnabled` (IME) is true
2. Reads IE register (0xFFFF) for enabled interrupts
3. Reads IF register (0xFF0F) for pending interrupts
4. If an enabled interrupt is pending, calls `handleInterrupt()`
5. `handleInterrupt()` should:
   - Push current PC to stack
   - Jump to the interrupt vector
   - Clear the interrupt flag in IF
   - Set `InterruptEnabled` to false

### Interrupt Vectors

| Interrupt | Vector | Bit in IE/IF |
|-----------|--------|--------------|
| V-Blank | 0x40 | 0 |
| Timer | 0x48 | 1 |
| LCD STAT | 0x50 | 2 |
| Joypad | 0x58 | 4 |

### Step 2: Update Main Loop

Uncomment and call `cpu.Interrupt()` in main.go after timer/graphics updates.

### Step 3: Test with cpu_instrs ROM

Once timer and interrupts are working, blargg's cpu_instrs test should show results.

---

## Phase 6: Input

### What You Need to Do

Implement joypad handling in memory and connect it to the CPU.

### Step 1: Joypad Register (0xFF00)

Bits of P1 register:
| Bit | When set to 0 | When set to 1 |
|-----|---------------|---------------|
| 5 (P15) | Select button keys (A, B, Select, Start) | Deselect |
| 4 (P14) | Select directional keys (Up, Down, Left, Right) | Deselect |
| 3 (P13) | Button A state | Read as 1 when not pressed |
| 2 (P12) | Button B state | Read as 1 when not pressed |
| 1 (P11) | Select state | Read as 1 when not pressed |
| 0 (P10) | Start state | Read as 1 when not pressed |

### Step 2: Reading Buttons

**To read directional keys:**
1. Write 0x10 to 0xFF00 (set bit 4 to 0, bit 5 to 1)
2. Read 0xFF00 - bits 0-3 now show Up/Down/Left/Right

**To read button keys:**
1. Write 0x20 to 0xFF00 (set bit 5 to 0, bit 4 to 1)
2. Read 0xFF00 - bits 0-3 now show A/B/Select/Start

### Step 3: Joypad Interrupt

When a button is pressed:
1. Set bit 4 of IF register (0xFF0F)
2. If IE has bit 4 set, the game can respond to the button press

---

## Phase 7: Main Loop Integration

### What You Need to Do

The current main loop exists but timer/graphics/ppu are stubs, and `cpu.Interrupt()` is commented out.

### Current main.go Issues

1. `timer.Update()` is empty - needs implementation
2. `graphics.Update()` is empty - PPU not implemented
3. `graphics.Render()` is empty - no frame output
4. `cpu.Interrupt()` is commented out - not implemented yet

### Step 1: Implement Timer First

See Phase 3 - timer must work before interrupts can fire.

### Step 2: Implement Interrupt Method

See Phase 5 - `cpu.Interrupt()` needs implementation.

### Step 3: Wire Everything Together

Once timer and interrupts are ready, uncomment `cpu.Interrupt()` in main.go and ensure proper frame timing.

---

## Testing Your Implementation

### Test Cartridges

You can find ROMs for testing:
- **Tetris** - requires mostly working CPU, timer, and basic graphics
- **CPU Tests** - more comprehensive opcode testing

### What to Look For

**If Tetris shows a black screen:**
- CPU might be stuck in a loop (check your JP/RET instructions)
- Timer might be affecting timing incorrectly

**If Tetris shows garbled graphics:**
- Check your VRAM reading (tiling system)
- Check background map selection

**If Tetris crashes:**
- Check your stack handling (PUSH/POP)
- Check your opcode timing (cycle counts affect everything)

### Minimal Test Sequence

1. Implement NOP (0x00) and verify CPU doesn't hang
2. Implement a few LD instructions
3. Test JP and see if PC moves correctly
4. Test CALL/RET
5. Test arithmetic (ADD, SUB)
6. Test CB-prefixed opcodes
7. Test interrupts (EI, DI, RETI)

---

## Summary Checklist

Before Tetris works, ensure each of these works:

- [x] CPU: All 156 opcodes implemented with correct cycle counts
- [x] Memory: MBC1 banking works for Tetris (RomOnly is fine for Tetris)
- [ ] Memory: I/O registers (0xFF00-0xFF7F) work properly
- [ ] Timer: DIV, TIMA, TMA, TAC function correctly
- [ ] Timer: Timer interrupt fires when TIMA overflows
- [ ] PPU: Scanlines advance correctly (LY increases)
- [ ] PPU: Mode transitions happen at correct cycles
- [ ] PPU: Background tiles render correctly
- [ ] PPU: V-Blank interrupt fires after scanline 144
- [ ] Interrupts: IME flag works with EI/DI
- [ ] Interrupts: V-Blank and Timer interrupts handled
- [ ] Main loop: Runs at proper speed, renders frames

---

## Quick Reference: Game Boy Boot State

After the boot ROM runs, the CPU starts at 0x100 with these register values:

| Register | Value |
|----------|-------|
| A | 0x01 |
| B | 0x00 |
| C | 0x13 |
| D | 0x00 |
| E | 0xD8 |
| H | 0x01 |
| L | 0x4D |
| F | 0xB0 |
| SP | 0xFFFE |
| PC | 0x0100 |

And these I/O register values are set:

| Address | Value | Register |
|---------|-------|----------|
| 0xFF00 | 0xCF | P1 |
| 0xFF04 | 0xAB | DIV |
| 0xFF05 | 0x00 | TIMA |
| 0xFF06 | 0x00 | TMA |
| 0xFF07 | 0xF8 | TAC |
| 0xFF0F | 0xE1 | IF |
| 0xFF40 | 0x91 | LCDC |
| 0xFF41 | 0x85 | STAT |
| 0xFF42 | 0x00 | SCY |
| 0xFF43 | 0x00 | SCX |
| 0xFF44 | 0x00 | LY |
| 0xFF45 | 0x00 | LYC |
| 0xFF46 | 0xFF | DMA |
| 0xFF47 | 0xFC | BGP |
| 0xFF4A | 0x00 | WY |
| 0xFF4B | 0x00 | WX |
| 0xFFFF | 0x00 | IE |

## Quick Reference: Cycle Budget

- **One scanline**: 456 cycles
- **One frame** (154 scanlines): 70,224 cycles
- **Visible lines**: 144 × 456 = 65,664 cycles
- **V-Blank lines**: 10 × 456 = 4,560 cycles
- **Target framerate**: ~59.7 FPS at 4.19 MHz