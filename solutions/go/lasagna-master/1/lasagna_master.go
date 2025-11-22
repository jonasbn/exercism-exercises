// Lasagna Master helper functions
package lasagna

// PreparationTime calculates the preparation time based on an average and the number of layers requested
func PreparationTime(layers []string, avgPreparationTime int) int {

	if avgPreparationTime == 0 {
		avgPreparationTime = 2
	}

	preparationTime := len(layers) * avgPreparationTime

	return preparationTime
}

// Quantities calculates the required quanties of noodles and sauce
func Quantities(layers []string) (int, float64) {

	var noodles int
	var sauce float64

	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}

	}

	return noodles, sauce
}

// AddSecretIngredient() adds a secret ingredient from another recipe when it needs something special
func AddSecretIngredient(friendsList []string, myList []string) {

	secretIngredient := friendsList[len(friendsList)-1]

	for i, ingredient := range myList {
		if ingredient == "?" {
			myList[i] = secretIngredient
		}
	}
}

// ScaleRecipe function scales quanties based on portions for two by the factor of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {

	newQuanties := make([]float64, len(quantities))

	for i, _ := range quantities {
		newQuanties[i] = float64(portions) * (quantities[i] / 2)
	}

	return newQuanties
}
