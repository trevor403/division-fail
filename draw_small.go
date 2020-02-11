package main

import "strings"

func printCharSmall(char byte, offsetX, offsetY int16) {
	charBitmap := BitmapFont[char-32]
	for i := uint32(0); i < 32; i++ {
		bit := (charBitmap >> i) & 1
		y := (8 - int16(i/4) + (8 * offsetY)) - 1
		x := (4 - int16(i%4) + (4 * offsetX))
		if bit == 1 {
			display.SetPixel(x, y, White)
		} else {
			display.SetPixel(x, y, Black)
		}
	}
}

func printStringSmall(s string, line int16) {
	for i, char := range s {
		printCharSmall(byte(char), int16(i), line)
	}
}

// term is 60x18

func printLorem() {
	const lorem = `[    0.000000] Linux version 5.4.0-12-generic (buildd@lcy01-amd64-006) (gcc version 9.2.1 20200117 (Ubuntu 9.2.1-24ubuntu1)) #15-Ubuntu SMP Tue Jan 21 15:12:29 UTC 2020 (Ubuntu 5.4.0-12.15-generic 5.4.8)
[    0.000000] Command line: BOOT_IMAGE=/vmlinuz-5.4.0-12-generic root=/dev/mapper/vgubuntu-root ro quiet splash acpi_rev_override=1 snd_hda_intel.dmic_detect=0 vt.handoff=7
[    0.000000] KERNEL supported cpus:
[    0.000000]   Intel GenuineIntel
[    0.000000]   AMD AuthenticAMD
[    0.000000]   Hygon HygonGenuine
[    0.000000]   Centaur CentaurHauls
[    0.000000]   zhaoxin   Shanghai  
[    0.000000] x86/fpu: Supporting XSAVE feature 0x001: 'x87 floating point registers'
[    0.000000] x86/fpu: Supporting XSAVE feature 0x002: 'SSE registers'
[    0.000000] x86/fpu: Supporting XSAVE feature 0x004: 'AVX
`
	loremLen := len(lorem)

	const maxSmallColumns = 60

	pos := loremLen
	for i := int16(0); pos > 0; i++ {
		realPos := loremLen - pos

		s := lorem[realPos : realPos+maxSmallColumns]
		newline := strings.Index(s, "\n")
		if newline == -1 {
			if pos < maxSmallColumns {
				printStringSmall(lorem[realPos:], i)
			} else {
				printStringSmall(s, i)
			}
		} else {
			printStringSmall(lorem[realPos:realPos+newline], i)
			pos -= newline + 1
			continue
		}

		pos -= maxSmallColumns
	}
}

func printLoremBig() {
	const lorem = `[    0.000000] Linux version 5.4.0-12-generic (buildd@lcy01-amd64-006) (gcc version 9.2.1 20200117 (Ubuntu 9.2.1-24ubuntu1)) #15-Ubuntu SMP Tue Jan 21 15:12:29 UTC 2020 (Ubuntu 5.4.0-12.15-generic 5.4.8)
[    0.000000] Command line: BOOT_IMAGE=/vmlinuz-5.4.0-12-generic root=/dev/mapper/vgubunt`

	loremLen := len(lorem)

	const maxSmallColumns = 30

	pos := loremLen
	for i := int16(0); pos > 0; i++ {
		realPos := loremLen - pos

		s := lorem[realPos : realPos+maxSmallColumns]
		newline := strings.Index(s, "\n")
		if newline == -1 {
			if pos < maxSmallColumns {
				printString(lorem[realPos:], i)
			} else {
				printString(s, i)
			}
		} else {
			printString(lorem[realPos:realPos+newline], i)
			pos -= newline + 1
			continue
		}

		pos -= maxSmallColumns
	}
}
