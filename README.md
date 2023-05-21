Package  colors
================
![Project status](https://img.shields.io/badge/version-1.3.1-green.svg)
[![GoDoc](https://godoc.org/github.com/go-playground/colors?status.svg)](https://pkg.go.dev/github.com/go-playground/colors)

Go color manipulation, conversion and printing library/utility

this library is currently in development, not all color types such as HSL, HSV and CMYK will be included in the first release; pull requests are welcome.

Installation
============

Use go get.

	go get github.com/go-playground/colors

Then import the validator package into your own code.

	import "github.com/go-playground/colors"
	
Usage and documentation
=======================

#Example
```go
hex, err := colors.ParseHEX("#fff")
rgb, err := colors.ParseRGB("rgb(0,0,0)")
rgb, err := colors.RGB(0,0,0)
rgba, err := colors.ParseRGBA("rgba(0,0,0,1)")
rgba, err := colors.RGBA(0,0,0,1)

// don't know which color, it was user selectable
color, err := colors.Parse("#000")

color.ToRGB()   // rgb(0,0,0)
color.ToRGBA()  // rgba(0,0,0,1)
color.ToHEX()   // #000000
color.IsLight() // false
color.IsDark()  // true

```

How to Contribute
=================

Make a pull request...

License
=======
Distributed under MIT License, please see license file in code for more details.
