package pad

import "strings"

func Left(s string, l int) string {
	if len(s) >= l {
		return s
	}

	return strings.Repeat("*", l-len(s)) + s
}

func Right(s string, l int) string {
	if len(s) >= l {
		return s
	}

	return s + strings.Repeat("*", l-len(s))
}
