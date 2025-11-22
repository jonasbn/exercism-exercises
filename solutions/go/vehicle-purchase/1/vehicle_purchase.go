package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	case "bike":
		return false
	case "stroller":
		return false
	case "e-scooter":
		return false
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	s := []string{option1, option2}
	sort.Strings(s)

	return fmt.Sprintf("%s is clearly the better choice.", s[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	var resellPrice float64

	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age < 10 {
		resellPrice = originalPrice * 0.7
	} else {
		resellPrice = originalPrice * 0.5
	}

	return resellPrice
}
