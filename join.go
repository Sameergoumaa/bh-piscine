package piscine

func Join(elems []string, sep string) string {
	words := ""
	for i := 0; i < len(elems); i++ {
		if i == len(elems)-1 {
			words = words + elems[i]
		} else {
			words = words + elems[i] + sep
		}
	}
	return words
}
