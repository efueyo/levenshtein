package levenshtein

// LDistance returns the Levenshtein distance between str1 and str2
func LDistance(str1, str2 string) int {
	if str1 == str2 {
		return 0
	}
	if len(str1) == 0 {
		return len(str2)
	}
	if len(str2) == 0 {
		return len(str1)
	}
	r0 := make([]int, len(str2)+1)
	r1 := make([]int, len(str2)+1)
	for i := range str2 {
		r0[i] = i
	}
	for i := range str1 {
		r1[0] = i + 1
		for j := range str2 {
			deletionCost := r0[j+1] + 1
			insertionCost := r1[j] + 1
			var substitutionCost int
			if str1[i] == str2[j] {
				substitutionCost = r0[j]
			} else {
				substitutionCost = r0[j] + 1
			}
			var cost int
			// min(deletionCost, inerstionCost, substitutionCost)
			if deletionCost < insertionCost {
				cost = deletionCost
			} else {
				cost = insertionCost
			}
			if cost > substitutionCost {
				cost = substitutionCost
			}
			r1[j+1] = cost
		}
		for i, j := range r1 {
			r0[i] = j
		}
	}
	return r1[len(str2)]
}
