package fg

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
	"black":         30,
	"red":           31,
	"green":         32,	
	"yellow":        33,	
	"blue":          34,	
	"magenta":       35,	
	"cyan":          36,
	"light gray":    37,
	"default":       39,
	"dark gray":     90,
	"light red":     91,
	"light green":   92,
	"light yellow":  93,
	"light blue":    94,
	"light magenta": 95,
	"light cyan":    96,
	"white":         97,
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

// Returns the strings joined and with the foreground color as Black.
func Black (s ...string) string {
	return format("black", s...)
}


// Returns the strings joined and with the foreground color as Red.
func Red (s ...string) string {
	return format("red", s...)
}

// Returns the strings joined and with the foreground color as Green.
func Green (s ...string) string {
	return format("green", s...)
}

// Returns the strings joined and with the foreground color as Yellow.
func Yellow (s ...string) string {
	return format("yellow", s...)
}

// Returns the strings joined and with the foreground color as Blue.
func Blue (s ...string) string {
	return format("blue", s...)
}

// Returns the strings joined and with the foreground color as Magenta.
func Magenta (s ...string) string {
	return format("magenta", s...)
}

// Returns the strings joined and with the foreground color as Cyan.
func Cyan (s ...string) string {
	return format("cyan", s...)
}

// Returns the strings joined and with the foreground color as Light Gray.
func LightGray (s ...string) string {
	return format("light gray", s...)
}

// Returns the strings joined and with the foreground color as Dark Gray.
func DarkGray (s ...string) string {
	return format("dark gray", s...)
}

// Returns the strings joined and with the foreground color as Light Red.
func LightRed (s ...string) string {
	return format("light red", s...)
}

// Returns the strings joined and with the foreground color as Light Green.
func LightGreen (s ...string) string {
	return format("light green", s...)
}

// Returns the strings joined and with the foreground color as Light Yellow.
func LightYellow (s ...string) string {
	return format("light yellow", s...)
}

// Returns the strings joined and with the foreground color as Light Blue.
func LightBlue (s ...string) string {
	return format("light blue", s...)
}

// Returns the strings joined and with the foreground color as Light Magenta.
func LightMagenta (s ...string) string {
	return format("light magenta", s...)
}

// Returns the strings joined and with the foreground color as Light Cyan.
func LightCyan (s ...string) string {
	return format("light cyan", s...)
}

// Returns the strings joined and with the foreground color as White.
func White (s ...string) string {
	return format("white", s...)
}

