package main

func main() {
	display.Configure()

	printString("                             |", 0)
	printString("                             |", 1)
	printString("                             |", 2)
	printString("_____________________________|", 3)
	printString("                             |", 4)
	printString("                             |", 5)
	printString("                             |", 6)
	printString("                             |", 7)
	printString("                             |", 8)
	printString("                             |", 9)

	// loop past 100 to pass fastSmalls
	for num := uint32(0); ; num++ {
		numString := FixedFormatUint(num)
		printString("num "+numString, 9)
	}

}
