package style

import (
	"fmt"
	"sort"
)

var code = map[string]int {
	"bold":      1,  
	"dim":       2,  
	"underline": 4,  
	"blink":     5,  
	"reverse":   7,  
	"hidden":    8,  
}

var reset = map[string]int {
	"all":       0,
	"bold":      21,  
	"dim":       22,  
	"underline": 24,  
	"blink":     25,  
	"reverse":   27,  
	"hidden":    28,  
}

func format(c string, s ...string) string {
	prefix := Code(c)

	var out string
	for _, v :=  range s {
		out += prefix + v
	}
	out += Reset(c)

	return out
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

func Code(s string) string {
	return fmt.Sprintf("\x1b[%dm", code[s])
}

func Reset(s string) string {
	return fmt.Sprintf("\x1b[%dm", reset[s])
}

// Returns the strings joined and formatted bold.
func Bold (s ...string) string {
	return format("bold", s...)
}

// Returns the strings joined and dimmed.
func Dim (s ...string) string {
	return format("dim", s...)
}

// Returns the strings joined and underlined.
func Underline (s ...string) string {
	return format("underline", s...)
}

// Returns the strings joined and blinking.
// Does not work on most terminal emulators.
func Blink (s ...string) string {
	return format("blink", s...)
}

// Returns the strings joined and the colors reversed.
func Reverse (s ...string) string {
	return format("reverse", s...)
}

// Returns the strings joined and hidden. 
// I have no idea what this does.
func Hidden (s ...string) string {
	return format("hidden", s...)
}



