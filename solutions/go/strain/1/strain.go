package strain

// Ints is a collection of integers
type Ints []int

// Lists is a collection of Lists
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep iterates over a collection and returns what satisfies the provided predicate
func (l Ints) Keep(fn func(int) bool) Ints {

	if l != nil {
		result := make(Ints, 0, len(l))
		for i := 0; i < len(l); i++ {
			if fn(l[i]) {
				result = append(result, l[i])
			}
		}
		return result
	}
	return nil
}

// Discard iterates over a collection and returns what does not satisfy the provided predicate
func (l Ints) Discard(fn func(int) bool) Ints {

	if l != nil {
		result := make(Ints, 0, len(l))
		for i := 0; i < len(l); i++ {
			if !fn(l[i]) {
				result = append(result, l[i])
			}
		}
		return result
	}
	return nil
}

// Keep iterates over a collection and returns what satisfies the provided predicate
func (l Strings) Keep(fn func(string) bool) Strings {
	result := make(Strings, 0, len(l))
	for i := 0; i < len(l); i++ {
		if fn(l[i]) {
			result = append(result, l[i])
		}
	}
	return result
}

// Discard iterates over a collection and returns what does not satisfy the provided predicate
func (l Strings) Discard(fn func(string) bool) Strings {
	result := make(Strings, 0, len(l))
	for i := 0; i < len(l); i++ {
		if !fn(l[i]) {
			result = append(result, l[i])
		}
	}
	return result
}

// Keep iterates over a collection and returns what satisfies the provided predicate
func (l Lists) Keep(fn func([]int) bool) Lists {
	result := make(Lists, 0, len(l))
	for i := 0; i < len(l); i++ {
		if fn(l[i]) {
			result = append(result, l[i])
		}
	}
	return result
}
