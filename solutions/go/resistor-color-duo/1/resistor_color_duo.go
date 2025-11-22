package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {

	values := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

	var resistance int

	factor := 10 // factor to multiply the resistance by position 2 is 10, position 1 is 1
	for i := 0; i < 2; i++ {

		color := colors[i]
		for idx, value := range values {
			if value == color {
				resistance += idx * factor
			}
		}
		factor = 1 // second band position is 1
	}

	return resistance
}
