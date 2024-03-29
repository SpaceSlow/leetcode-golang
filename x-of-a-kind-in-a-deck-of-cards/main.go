package x_of_a_kind_in_a_deck_of_cards

// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/description/

func hasGroupsSizeX(deck []int) bool {
	freqs := make(map[int]int)

	for _, card := range deck {
		freqs[card]++
	}

	gcd := freqs[deck[0]]
	for _, value := range freqs {
		if value == 1 {
			return false
		}
		gcd = GCDEuclidean(gcd, value)
	}

	return gcd != 1
}

func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
