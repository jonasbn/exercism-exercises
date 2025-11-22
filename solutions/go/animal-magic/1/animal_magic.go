package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	seconds := time.Now().UnixMicro()
	rand.Seed(seconds)
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return float64(rand.Intn(12))
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	var shuffledAnimals []string

	for i := len(animals); i > 0; i-- {
		randomIndex := rand.Intn(len(animals))
		animal := animals[randomIndex]

		newAnimals := animals[:randomIndex]
		randomIndex++
		var tempAnimals = animals[randomIndex:]

		for j := 0; j < len(tempAnimals); j++ {
			newAnimals = append(newAnimals, tempAnimals[j])
		}
		animals = newAnimals

		shuffledAnimals = append(shuffledAnimals, animal)
	}

	return shuffledAnimals
}
