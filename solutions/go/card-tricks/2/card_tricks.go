package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var slice []int

	if length > 0 {
		for i := 0; i < length; i++ {
			slice = append(slice, value)
		}
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		pre := slice[:index]
		post := slice[index+1:]
		slice = append(pre, post...)
	}

	return slice
}

func FavoriteCards() []int {
	favoriteCards := []int{2, 6, 9}

	return favoriteCards
}

func PrependItems(slice []int, numbers ...int) []int {
	var newSlice []int

	for i := 0; i < len(numbers); i++ {
		newSlice = append(newSlice, numbers[i])
	}

	for j := 0; j < len(slice); j++ {
		newSlice = append(newSlice, slice[j])
	}

	return newSlice
}
