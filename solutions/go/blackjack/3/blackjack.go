package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int = 0

	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}

	return value
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	option := ""

	switch {
	case isBlackjack && dealerScore < 10:
		option = "W" // Automatic win
	case isBlackjack && dealerScore >= 10:
		option = "S" // stand
	default:
		option = "P" // split
	}

	return option
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	option := ""

	switch {
	case handScore >= 17:
		option = "S" // stand
	case handScore <= 16 && handScore >= 12 && dealerScore >= 7:
		option = "H" // hit
	case handScore <= 16 && handScore >= 12 && dealerScore < 7:
		option = "S" // stand
	case handScore <= 11:
		option = "H" // hit
	}
	return option
}

func FirstTurn(card1, card2, dealerCard string) string {
	option := ""

	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case playerScore > 20:
		option = LargeHand(IsBlackjack(card1, card2), dealerScore)
	case playerScore < 21:
		option = SmallHand(playerScore, dealerScore)
	}

	return option
}
