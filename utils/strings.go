package utils

import "strings"

// FixedWidthAppend returns a string with an exact length, appending whitespace if extra characters are needed
func FixedWidthAppend(s string, width int) string {
	if len(s) > width {
		return s[:width]
	}

	return s + strings.Repeat(" ", width-len(s))
}

// FixedWidthPrepend works exactly like FixedWidthAppend, but prepends whitespace instead of appending
func FixedWidthPrepend(s string, width int) string {
	if len(s) > width {
		return s[:width]
	}

	return strings.Repeat(" ", width-len(s)) + s
}
