package utils

import (
	"fmt"
	"strconv"
)

func leftPadding(s string, char string, length int) string {
	ns := s
	for i := len(s); i < length; i += 1 {
		ns = char + ns
	}
	return ns
}

func AddStrings(a, b string) (string, error) {
	na := a
	nb := b

	if len(a) < len(b) {
		na = leftPadding(na, "0", len(b))
	} else if len(b) < len(a) {
		nb = leftPadding(nb, "0", len(a))
	}

	r := ""

	ra := []rune(na)
	rb := []rune(nb)

	carry := 0

	for i := len(ra) - 1; i >= 0; i -= 1 {
		sa := fmt.Sprintf("%c", ra[i])
		ia, err := strconv.Atoi(sa)
		if err != nil {
			return "0", err
		}

		sb := fmt.Sprintf("%c", rb[i])
		ib, err := strconv.Atoi(sb)
		if err != nil {
			return "0", err
		}

		ic := ia + ib + carry

		if ic >= 10 {
			ic -= 10
			carry = 1
		} else {
			carry = 0
		}

		cc := strconv.Itoa(ic)

		r = cc + r
	}

	if carry == 1 {
		r = "1" + r
	}

	return r + "0", nil
}
