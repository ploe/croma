package bg

import (
	"fmt"
	"sort"
)

// Returns s joined and color formatted.
// Each element in s is prefixed with the color code for c.
// The idea is that codes can be nested within each other.
// The last thing is appended is a return to the default color.
func format(c string, s ...string) string {
	color := Code(c)

	var out string
	for _, v :=  range s {
		out += color + v
	}
	out += Code("default")

	return out
}

var code = map[string]int {
	"black":         40,
	"red":           41,
	"green":         42,	
	"yellow":        43,	
	"blue":          44,	
	"magenta":       45,	
	"cyan":          46,
	"light gray":    47,
	"default":       49,
	"dark gray":     100,
	"light red":     101,
	"light green":   102,
	"light yellow":  103,
	"light blue":    104,
	"light magenta": 105,
	"light cyan":    106,
	"white":         107,
}

func Demo() string {
	var keys []string
	for k,_ := range code {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var out string
	for _, v := range keys {
		out += format(v, v) + "\n"
	}
	return out
}


// Returns the color code for the 8/16 colors.
// s is a lowercase string.
func Code(s string) string {
	return fmt.Sprintf("\x1b[%dm", code[s])
}

// Returns the strings joined and with the background color as Black.
func Black (s ...string) string {
	return format("black", s...)
}


// Returns the strings joined and with the background color as Red.
func Red (s ...string) string {
	return format("red", s...)
}

// Returns the strings joined and with the background color as Green.
func Green (s ...string) string {
	return format("green", s...)
}

// Returns the strings joined and with the background color as Yellow.
func Yellow (s ...string) string {
	return format("yellow", s...)
}

// Returns the strings joined and with the background color as Blue.
func Blue (s ...string) string {
	return format("blue", s...)
}

// Returns the strings joined and with the background color as Magenta.
func Magenta (s ...string) string {
	return format("magenta", s...)
}

// Returns the strings joined and with the background color as Cyan.
func Cyan (s ...string) string {
	return format("cyan", s...)
}

// Returns the strings joined and with the background color as Light Gray.
func LightGray (s ...string) string {
	return format("light gray", s...)
}

// Returns the strings joined and with the background color as Dark Gray.
func DarkGray (s ...string) string {
	return format("dark gray", s...)
}

// Returns the strings joined and with the background color as Light Red.
func LightRed (s ...string) string {
	return format("light red", s...)
}

// Returns the strings joined and with the background color as Light Green.
func LightGreen (s ...string) string {
	return format("light green", s...)
}

// Returns the strings joined and with the background color as Light Yellow.
func LightYellow (s ...string) string {
	return format("light yellow", s...)
}

// Returns the strings joined and with the background color as Light Blue.
func LightBlue (s ...string) string {
	return format("light blue", s...)
}

// Returns the strings joined and with the background color as Light Magenta.
func LightMagenta (s ...string) string {
	return format("light magenta", s...)
}

// Returns the strings joined and with the background color as Light Cyan.
func LightCyan (s ...string) string {
	return format("light cyan", s...)
}

// Returns the strings joined and with the background color as White.
func White (s ...string) string {
	return format("white", s...)
}

