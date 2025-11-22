package diamond

import "fmt"

func Gen(char byte) (string, error) {
	const separator string = " "

	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid character: %c", char)
	}
	var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var diamond string
	var index int = 0

	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == string(char) {
			index = i
			break
		}
	}

	// Lines
	for i := 0; i <= index; i++ {
		// prefix spaces for each line
		for j := 0; j < index-i; j++ {
			diamond += separator
		}
		// first letter
		diamond += alphabet[i]
		// middle spaces
		if i > 0 {
			for j := 0; j < 2*i-1; j++ {
				diamond += separator
			}
			diamond += alphabet[i]
		}
		// postfix spaces
		if i < len(alphabet) {
			for j := index - i; j > 0; j-- {
				diamond += separator
			}
		}
		if index > 0 {
			diamond += "\n"
		}
	}

	// We subtract 1 from the index to avoid repeating the middle row
	for i := index - 1; i >= 0; i-- {
		for j := 0; j < index-i; j++ {
			diamond += separator
		}
		diamond += alphabet[i]
		if i > 0 {
			for j := 0; j < 2*i-1; j++ {
				diamond += separator
			}
			diamond += alphabet[i]
		}
		// postfix spaces
		if i < len(alphabet)-1 {
			for j := i + 1; j <= index; j++ {
				diamond += separator
			}
		}

		if i > 0 {
			diamond += "\n"
		}
	}

	return diamond, nil
}
