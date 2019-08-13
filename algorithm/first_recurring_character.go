package algorithm

func FirstRecurringCharacter(in string) rune {
	cache := make(map[rune]uint8)
	for _, runeValue := range in {
		if cache[runeValue] == 1 {
			return runeValue
		}
		cache[runeValue] = 1
	}
	return 0;
}
