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

## Phase 1: CPU Opcodes

### What You Need to Do

The `Run()` method in `pkg/cpu/gbz80.go` currently returns 0. You need to:

1. **Implement an instruction fetch cycle**
2. **Decode the opcode** (and CB-prefixed opcodes)
3. **Execute the instruction**
4. **Return the number of cycles used**

### Step 1: Fetch-Decode-Execute Loop

Start by implementing the basic loop structure:

```go
func (gbz *GBZ80) Run() uint {
    opcode := gbz.fetch()
    cycles := gbz.execute(opcode)
    return cycles
}

func (gbz *GBZ80) fetch() byte {
    opcode := gbz.mem.Read(gbz.pc)
    gbz.pc++
    return opcode
}
```

### Step 2: Categorize Opcodes

Organize opcodes by their first byte (0x00-0xFF):

| Range | Opcodes | Example Instructions |
|-------|---------|---------------------|
| 0x00-0x3F | Misc/Load | NOP, LD, PUSH, POP, ADD, SUB, CALL, RET |
| 0x40-0x7F | Load between registers | LD A,B, LD B,A, etc. (no immediate values) |
| 0x80-0xBF | Arithmetic | ADD A, SUB A, AND, OR, XOR, CP |
| 0xC0-0xDF | Control flow | JP, CALL, PUSH, POP, prefixed ops |
| 0xE0-0xEF | Special | LD (nn), SP, RETI, etc. |
| 0xF0-0xFF | Special | LD A, (FF00+n), EI, DI, etc. |

### Step 3: Implement by Category

**Category 1: 8-bit Load Instructions (0x40-0xBF)**
These are the simplest - copy a value from one register to another.

Example: `LD B, C` (0x41)
- Read value from C register
- Write value to B register
- Return 4 cycles
- PC doesn't change (single byte instruction)

Example: `LD A, n` (0x3E nn)
- Read immediate byte nn
- Write to A register
- Return 8 cycles (2 bytes)
- PC advances by 2

**Category 2: 16-bit Load Instructions (PUSH/POP)**
These affect the stack.

Example: `PUSH BC` (0xC5)
- SP -= 2
- Write B to SP+1, C to SP
- Return 16 cycles

Example: `POP DE` (0xD1)
- Read C from SP, B from SP+1
- SP += 2
- Write to D (high) and E (low)
- Return 12 cycles

**Category 3: Arithmetic Instructions**

Example: `ADD A, B` (0x80)
- Add B to A
- Set flags (Z, N=0, H if half-carry, C if carry)
- Return 4 cycles

Example: `INC A` (0x3C)
- Increment A
- Set Z flag if result is 0
- Set N=0, H if half-carry
- Return 4 cycles

**Category 4: Control Flow (JP, JR, CALL, RET)**

Example: `JP nn` (0xC3 nn nn)
- Read 16-bit immediate
- Set PC to nn
- Return 16 cycles

Example: `JR NZ, n` (0x20 nn)
- If Z flag is 0, add signed byte n to PC
- Return 12 cycles if taken, 8 if not

Example: `CALL nn` (0xCD nn nn)
- Push current PC to stack
- Set PC to nn
- Return 24 cycles

Example: `RET` (0xC9)
- Pop PC from stack
- Return 16 cycles

**Category 5: CB-Prefixed Opcodes (0xCB xx)**

When you encounter 0xCB:
1. Fetch the next byte
2. Use it to index into CB instruction table
3. Execute the CB instruction

CB instructions include:
- `BIT b, r` - test if bit b of register r is 0 (sets Z flag)
- `SET b, r` - set bit b of register r to 1
- `RES b, r` - reset bit b of register r to 0
- `RLC r` - rotate left through carry
- `RRC r` - rotate right through carry
- `RL r` - rotate left through carry
- `RR r` - rotate right through carry
- `SLA r` - shift left arithmetic
- `SRA r` - shift right arithmetic
- `SRL r` - shift right logical
- `SWAP r` - swap upper and lower nibbles

### Step 4: Implement Flag Logic

The flag register F has these bits:
- Bit 7: Z (Zero) - set when result is 0
- Bit 6: N (Subtract) - set when last operation was subtraction
- Bit 5: H (Half-carry) - set when lower 4 bits overflowed
- Bit 4: C (Carry) - set when there was a carry/borrow

For each arithmetic operation, you must update these correctly.

**Example - ADD A, B:**
```
temp = A + B
if (temp > 0xFF) set C flag
if ((A & 0x0F) + (B & 0x0F) > 0x0F) set H flag
A = temp
if (A == 0) set Z flag
```

### Reference: Important Opcodes to Implement

You need about 156 unique instructions. Use this checklist:

**Loads (about 50 instructions):**
- LD r, r' (register to register)
- LD r, n (immediate)
- LD (HL), r and LD r, (HL)
- LD (nn), A and LD A, (nn)
- LD (FF00+n), A and LD A, (FF00+n)
- LD (FF00+C), A
- LD SP, nn
- LD SP, HL
- PUSH rr, POP rr
- LDI (HL), A, LDI A, (HL)
- LDD (HL), A, LDD A, (HL)

**Arithmetic/Logic (about 40 instructions):**
- ADD A, r, ADD A, n
- ADC A, r (add with carry)
- SUB r, SUB n
- SBC A, r (subtract with carry)
- AND r, AND n
- OR r, OR n
- XOR r, XOR n
- CP r, CP n (compare)
- INC r, DEC r
- INC (HL), DEC (HL)

**16-bit Arithmetic (about 10):**
- ADD HL, SP
- ADD HL, rr
- INC rr, DEC rr
- LD HL, SP+n

**Control Flow (about 25):**
- JP nn, JP cc, nn
- JR n, JR cc, n
- CALL nn, CALL cc, nn
- RET, RET cc, RETI
- RST n (restart)

**Block Instructions (about 10):**
- LDIR, LDDR, CPIR, CPDR, etc. (can skip for Tetris initially)

**Prefix CB Instructions (about 50):**
- RLC, RRC, RL, RR
- SLA, SRA, SRL
- SWAP
- BIT b, r
- SET b, r
- RES b, r

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

Update the CPU to handle interrupts properly.

### Step 1: Track Interrupt State

In your CPU struct:
```go
type GBZ80 struct {
    // ... existing fields ...
    ime bool // interrupt master enable
}
```

Initialize ime to false (boot ROM has IME=0).

### Step 2: Implement EI and DI

- `EI` (0xFB): Set ime = true AFTER next instruction
- `DI` (0xF3): Set ime = false immediately

### Step 3: Implement Interrupt Check

At the end of each instruction, check for interrupts:

```go
func (gbz *GBZ80) checkInterrupts() {
    if !gbz.ime {
        return
    }

    // Check each interrupt
    if (gbz.mem.Read(0xFFFF) & 0x01) != 0 && (gbz.mem.Read(0xFF0F) & 0x01) != 0 {
        // V-Blank interrupt
        gbz.handleInterrupt(0x40)
        return
    }
    if (gbz.mem.Read(0xFFFF) & 0x02) != 0 && (gbz.mem.Read(0xFF0F) & 0x02) != 0 {
        // Timer interrupt
        gbz.handleInterrupt(0x48)
        return
    }
    // ... etc
}
```

### Step 4: Implement Interrupt Handler

```go
func (gbz *GBZ80) handleInterrupt(vector uint16) {
    // Push current PC to stack
    gbz.sp--
    gbz.mem.Write(gbz.sp, byte(gbz.pc>>8))
    gbz.sp--
    gbz.mem.Write(gbz.sp, byte(gbz.pc))

    // Jump to vector
    gbz.pc = vector

    // Disable interrupts
    gbz.ime = false

    // Clear the interrupt flag
    // ... (clear appropriate bit in IF)
}
```

### Step 5: Implement RETI

`RETI` (0xD9) is like `RET` but also sets ime = true.

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

Fix the main loop in `main.go` to run properly.

### Step 1: Frame-Based Timing

The game should run like this:
1. Run CPU until enough cycles for one frame (70224 cycles)
2. During that time, PPU and timer advance
3. At end of frame, display the rendered frame

### Step 2: Proper Loop Structure

```go
func Run(c *cartridge.Cartridge) {
    memory := memory.NewMemory(c)
    cpu := cpu.NewGBZ80(memory)
    ppu := graphics.NewPPU(memory)
    timer := timer.NewTimer(memory)

    for {
        // Run for one scanline's worth of cycles
        cycles := 0
        for cycles < 456 {
            insCycles := cpu.Run()
            cycles += insCycles
            timer.Update(insCycles)
            ppu.Update(insCycles)
        }

        // Check for interrupts
        cpu.Interrupt()

        // If on scanline 144, we're in V-Blank - safe to render
        if ppu.LY() == 144 {
            graphics.Render(ppu.Pixels())
        }
    }
}
```

### Step 3: Handle Graceful Exit

Add a way to quit (e.g., press ESC or close window).

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

- [ ] CPU: All 156 opcodes implemented with correct cycle counts
- [ ] Memory: MBC1 banking works for Tetris (RomOnly is fine for Tetris)
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