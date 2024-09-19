package utils

// ?: I found packages which provide set datatypes, but on forums I saw most people recommend to use map keys as sets
func Uniques(xs []string) []string {
	ys := make(map[string]bool)
	for _, x := range xs {
		ys[x] = true
	}

	keys := make([]string, 0, len(ys))
	for k := range ys {
		keys = append(keys, k)
	}

	return keys
}
