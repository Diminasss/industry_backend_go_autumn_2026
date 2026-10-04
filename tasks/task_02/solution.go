package main

func rotateRunes(s string, shift int) string {
	runes := []rune(s)
	stringLength := len(runes)

	// Краевые случаи, когда строка будет та же
	if stringLength == shift || stringLength == 1 || stringLength == 0 || shift == 0 {
		return string(runes)
	}

	if shift > 0 {
		// Вращение влево
		shift %= stringLength
	} else {
		// Вращение вправо
		shift = ((shift % stringLength) + stringLength) % stringLength
	}

	result := make([]rune, 0, stringLength)
	result = append([]rune{}, runes[shift:]...)
	result = append(result, runes[:shift]...)
	return string(result)
}
