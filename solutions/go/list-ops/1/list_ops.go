package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (list IntList) Foldl(fn binFunc, accumulator int) int {
	for _, v := range list {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

func (list IntList) Foldr(fn binFunc, accumulator int) int {
	for _, v := range list.Reverse() {
		accumulator = fn(v, accumulator)
	}
	return accumulator
}

func (list IntList) Length() int {
	var length int = 0
	for k := range list {
		length = k + 1
	}

	return length
}

func (list IntList) Map(fn unaryFunc) IntList {
	var new_list IntList = make(IntList, list.Length())

	for k, v := range list {
		new_list[k] = fn(v)
	}
	return new_list

}

func (list IntList) Filter(fn predFunc) IntList {
	new_list := make(IntList, 0)

	for _, v := range list {
		if fn(v) {
			sub_list := make(IntList, 1)
			sub_list[0] = v
			new_list = new_list.Append(sub_list)
		}
	}
	return new_list
}

func (list IntList) Reverse() IntList {
	var reversed_list IntList = make(IntList, list.Length())

	for k, v := range list {
		reversed_list[list.Length()-k-1] = v
	}

	return reversed_list
}

func (list IntList) Concat(lists []IntList) IntList {
	var new_list IntList

	new_list = new_list.Append(list)

	for _, l := range lists {
		new_list = new_list.Append(l)
	}

	return new_list
}

func (list1 IntList) Append(list2 IntList) IntList {
	var new_list IntList = make(IntList, list1.Length()+list2.Length())

	for k, v := range list1 {
		new_list[k] = v
	}

	for k, v := range list2 {
		new_list[k+list1.Length()] = v
	}

	return new_list
}
