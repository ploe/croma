# What is croma?

croma is set of Go packages for colors and formatting within the \*nix terminal.

Its packages provide a clean, declaritive interfaces for generating formatted strings.

It is designed to be lightweight and straightforward.

# How do I get it?

Ensure your GOPATH environment variable is set correctly and use the command:

`go get github.com/ploe/croma`

# What are the packages provided?

## fg and bg

The package used for specifying the foreground and background colors are called `fg` and `bg` respectively. 

To include it in your project you will need to `import` them in to your `package`:

```
import (
	"github.com/ploe/croma/fg"
	"github.com/ploe/croma/bg"
)
```

`fg` and `bg` have the following functions for formatting colors:

```
func Black(s ...string) string
func Blue(s ...string) string
func Cyan(s ...string) string
func DarkGray(s ...string) string
func Green(s ...string) string
func LightBlue(s ...string) string
func LightCyan(s ...string) string
func LightGray(s ...string) string
func LightGreen(s ...string) string
func LightMagenta(s ...string) string
func LightRed(s ...string) string
func LightYellow(s ...string) string
func Magenta(s ...string) string
func Red(s ...string) string
func White(s ...string) string
func Yellow(s ...string) string
```

Each of these functions return a string that is formatted with the specified color. 

They are variadic; What this means is that they can take a list of strings and the output is a concatenated version of that list appropiately formatted. 

Here's an example of how to use these functions:

```
fmt.Println(fg.Red("I print red text!"))

str := bg.Magenta("My background is Magenta! ", "So is mine ", "And mine...")
fmt.Printf("%s\t<--- Magenta", str)

fmt.Println(
	bg.Cyan(
		"\ncyan background?",
		fg.Red("\nred"),
		fg.Yellow("\nyellow"),
		bg.Blue(
			fg.Green("\ngreen with a blue background"),
		),
		fg.LightGray("\nlight grey"),
	),
	
)
```

All text formatted this way is suffixed with a tag returning the input to its default state.

As you can see from the above format functions can be nested. This enables you to keep your formatting succinct, but flexible.

There are two more functions. The first one is `Code`:


```
func Code(s string) string
```

With `Code` You can specify a color with a string and get back the escape code for it.


```
red\_fg := fg.Code("red")
green\_bg := bg.Code("green")
terminate := bg.Code("default") + fg.Code("default")
fmt.Println(red\_fg + green\_bg + "I am red on green" + terminate + " I'm back to normal.")
```

The string values you can pass to `Code` are:

```
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
```

And last but not least we have the function `Demo`:

```
func Demo() string
```

With `Demo` you can return a color formatted demo string. You can use this to ascertain what colors are supported in your terminal emulator. It also looks really pretty! ;)
```
fmt.Printf(fg.Demo())
fmt.Printf(bg.Demo())
```

## style - Style

```Docs coming soon...```

# What license is croma available under?

croma is distributed under a FreeBSD (2-clause) License.

```
Copyright (c) 2016, Myke Atkinson
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
