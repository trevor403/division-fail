package main

import (
	"image/color"
	"machine"
)

var (
	Black = color.RGBA{0x00, 0x00, 0x00, 0xff}
	White = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

var display = machine.Display

func printChar(char byte, offsetX, offsetY int16) {
	charBitmap := BitmapFont[char-32]
	for i := uint32(0); i < 32; i++ {
		bit := (charBitmap >> i) & 1
		y := ((8 - int16(i/4) + (8 * offsetY)) * 2) - 2
		x := ((4 - int16(i%4) + (4 * offsetX)) * 2) - 1
		if bit == 1 {
			display.SetPixel(x, y, White)
			display.SetPixel(x+1, y, White)
			display.SetPixel(x, y+1, White)
			display.SetPixel(x+1, y+1, White)
		} else {
			display.SetPixel(x, y, Black)
			display.SetPixel(x+1, y, Black)
			display.SetPixel(x, y+1, Black)
			display.SetPixel(x+1, y+1, Black)
		}
	}
}

func printString(s string, line int16) {
	for i, char := range s {
		printChar(byte(char), int16(i), line)
	}
}
