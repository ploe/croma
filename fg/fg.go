package fg

import (
	"fmt"
)

func format(c int, s string) string {
	return fmt.Sprintf("\x1b[%d;1m%s\x1b[0m", c, s)
}

var fgcode = map[string]int {
	"black":   30,
	"red":     31,
	"green":   32,	
	"yellow":  33,	
	"blue":    34,	
	"magenta": 35,	
	"cyan":    36,
	"blink":    5,
} 

// Return the string with the foreground color as Black.
func Black (s string) string {
	return format(fgcode["black"], s)
}

// Return the string with the foreground color as Red.
func Red (s string) string {
	return format(fgcode["red"], s)
}

// Return the string with the foreground color as Green.
func Green (s string) string {
	return format(fgcode["green"], s)
}

// Return the string with the foreground color as Yellow.
func Yellow (s string) string {
	return format(fgcode["yellow"], s)
}

// Return the string with the foreground color as Blue.
func Blue (s string) string {
	return format(fgcode["blue"], s)
}

// Return the string with the foreground color as Magenta.
func Magenta (s string) string {
	return format(fgcode["magenta"], s)
}

// Return the string with the foreground color as Cyan.
func Cyan (s string) string {
	return format(fgcode["cyan"], s)
}

