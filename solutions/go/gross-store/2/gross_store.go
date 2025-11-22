package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	i := units[unit]
	if i != 0 {
		if bill[item] != 0 {
			bill[item] += i
		} else {
			bill[item] = i
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	i := units[unit]
	if i != 0 {
		if bill[item] != 0 {
			price := bill[item] - i

			switch {
			case price == 0:
				delete(bill, item)
				return true
			case price > 0:
				bill[item] -= i
				return true
			}
		}
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if bill[item] != 0 {
		return bill[item], true
	} else {
		return 0, false
	}
}
