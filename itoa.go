package main

func FixedFormatUint(value uint32) string {
	var powers = [...]uint32{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}

	buf := make([]byte, 10)
	bufIndex := 0

	foundFront := false

	for i := 8; i >= 0; i-- {
		var digit byte = '0'
		power := powers[i]
		for value >= power {
			value -= power
			digit++
		}

		// skip leading zeros
		if digit != '0' {
			foundFront = true
		} else if !foundFront {
			continue
		}

		buf[bufIndex] = digit
		bufIndex++
	}
	buf[bufIndex] = byte(value + '0')
	bufIndex++

	return string(buf[:bufIndex])
}
