package pad

import (
	"strings"
)

func LeftPad(s string, pad string, num int) string {
	if num <= 0 {
		return s
	}

	b := strings.Builder{}

	for num > 0 {
		b.WriteString(pad)
		num--
	}

	b.WriteString(s)
	return b.String()
}

func RightPad(s string, pad string, num int) string {
	if num <= 0 {
		return s
	}

	b := strings.Builder{}
	b.WriteString(s)

	for num > 0 {
		b.WriteString(pad)
		num--
	}
	return b.String()
}
