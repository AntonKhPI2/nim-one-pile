package game

func isLosingStart(variant Variant, N, k int) bool {
	m := k + 1
	if variant == Normal {
		return N%m == 0
	}
	return N%m == 1
}

func BestResponse(variant Variant, remaining, k int) int {
	if remaining <= 0 {
		return 0
	}
	m := k + 1
	if variant == Normal {
		r := remaining % m
		if r == 0 {
			return 0
		}
		return r
	}
	r := (remaining - 1) % m
	if r == 0 {
		return 0
	}
	if r == remaining {
		if r > 1 {
			return r - 1
		}
		return 0
	}
	return r
}
