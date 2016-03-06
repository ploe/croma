package fg

import (
	"fmt"
	"strings"
)

func format(c int, s ...string) string {
	j := strings.Join(s, "")
	return fmt.Sprintf("\x1b[%d;1m%s\x1b[39m", c, j)
}

var code = map[string]int {
	"black":         30,
	"red":           31,
	"green":         32,	
	"yellow":        33,	
	"blue":          34,	
	"magenta":       35,	
	"cyan":          36,
	"light gray":    37,
	"light grey":    37,
	"default":       39,
	"dark gray":     90,
	"dark grey":     90,
	"light red":     91,
	"light green":   92,
	"light yellow":  93,
	"light blue":    94,
	"light magenta": 95,
	"light cyan":    96,
	"white":         97,
	
}

// Returns the color code for the 8/16 colors.
// s is a lowercase string.
func Code(s string) int {
	return code[s]
}

// Returns the strings joined and with the foreground color as Black.
func Black (s ...string) string {
	return format(code["black"], s...)
}

// Returns the strings joined and with the foreground color as Red.
func Red (s ...string) string {
	return format(code["red"], s...)
}

// Returns the strings joined and with the foreground color as Green.
func Green (s ...string) string {
	return format(code["green"], s...)
}

// Returns the strings joined and with the foreground color as Yellow.
func Yellow (s ...string) string {
	return format(code["yellow"], s...)
}

// Returns the strings joined and with the foreground color as Blue.
func Blue (s ...string) string {
	return format(code["blue"], s...)
}

// Returns the strings joined and with the foreground color as Magenta.
func Magenta (s ...string) string {
	return format(code["magenta"], s...)
}

// Returns the strings joined and with the foreground color as Cyan.
func Cyan (s ...string) string {
	return format(code["cyan"], s...)
}

