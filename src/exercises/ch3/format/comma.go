package format

import (
	"bytes"
	"strings"
)

func CommaFloat(s string, recursive bool) string {
	sign := ""
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		sign = "-"
	}
	if i:= strings.LastIndex(s, "."); i >= 0 {
		if recursive {
			return sign + commaIntegerRecursive(s[:i]) + s[i:]
		} else {
			return sign + commaIntegerIterative(s[:i]) + s[i:]
		}
	}
	if recursive {
		return sign + commaIntegerRecursive(s)
	} else {
		return sign + commaIntegerRecursive(s)
	}
}

// Recursively inserts commas in the string that represents a positive integer (e.g. 12345)
func commaIntegerRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaIntegerRecursive(s[:n-3]) + "," + s[n-3:]
}

// Iteratively inserts commas in the string that represents a positive integer (e.g. 12345)
func commaIntegerIterative(s string) string {
	var buf bytes.Buffer
	n := len(s)
	triplets := n / 3
	leftovers := n % 3
	var si int; // string index, starts at 0
	for i := 0; i < leftovers; i++ {
		buf.WriteByte(s[si])
		si++
	}
	buf.WriteByte(',')
	for i := 0;  ; i++ {
		buf.WriteByte(s[si])
		si++
		buf.WriteByte(s[si])
		si++
		buf.WriteByte(s[si])
		si++
		if i != triplets - 1 {
			buf.WriteByte(',')
		} else {
			break
		}
	}
	return buf.String()
}
