package util

// Trim trims src string with toTrim uint8.
func Trim(src string, toTrim uint8) string {
	if src[0] == toTrim {
		src = src[1:]
	}

	if i := len(src) - 1; src[i] == toTrim {
		src = src[:i]
	}

	return src
}
