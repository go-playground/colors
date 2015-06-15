package colors

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	rgbaString                    = "rgba(%d,%d,%d,%g)"
	rgbaCaptureRegexString        = "^rgba\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*((0.[1-9]*)|[01])\\s*\\)$"
	rgbaCaptureRegexPercentString = "^rgba\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*((0.[1-9]*)|[01])\\s*\\)$"
)

var (
	rgbaCaptureRegex        = regexp.MustCompile(rgbaCaptureRegexString)
	rgbaCapturePercentRegex = regexp.MustCompile(rgbaCaptureRegexPercentString)
)

// RGBAColor represents an RGBA color
type RGBAColor struct {
	R uint8
	G uint8
	B uint8
	A float64
}

// ParseRGBA validates an parses the provided string into an RGBAColor object
// supports both RGBA 255 and RGBA as percentages
func ParseRGBA(s string) (*RGBAColor, error) {

	s = strings.ToLower(s)

	var isPercent bool
	vals := rgbaCaptureRegex.FindAllStringSubmatch(s, -1)

	if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {

		vals = rgbaCapturePercentRegex.FindAllStringSubmatch(s, -1)

		if vals == nil || len(vals) == 0 || len(vals[0]) == 0 {
			return nil, ErrBadColor
		}

		isPercent = true
	}

	r, _ := strconv.ParseUint(vals[0][1], 10, 8)
	g, _ := strconv.ParseUint(vals[0][2], 10, 8)
	b, _ := strconv.ParseUint(vals[0][3], 10, 8)
	a, _ := strconv.ParseFloat(vals[0][4], 64)

	if isPercent {
		r = uint64(math.Floor(float64(r)/100*255 + .5))
		g = uint64(math.Floor(float64(g)/100*255 + .5))
		b = uint64(math.Floor(float64(b)/100*255 + .5))
	}

	return &RGBAColor{R: uint8(r), G: uint8(g), B: uint8(b), A: a}, nil
}

// RGBA validates and returns a new RGBAColor object from the provided r, g, b, a values
func RGBA(r, g, b uint8, a float64) (*RGBAColor, error) {

	if a < 0 || a > 1 {
		return nil, ErrBadColor
	}

	return &RGBAColor{R: r, G: g, B: b, A: a}, nil
}

// String returns the string representation on the RGBAColor
func (c *RGBAColor) String() string {
	return fmt.Sprintf(rgbaString, c.R, c.G, c.B, c.A)
}

// ToHEX converts the RGBAColor to a HEXColor
func (c *RGBAColor) ToHEX() *HEXColor {
	return &HEXColor{hex: fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)}
}

// ToRGB converts the RGBAColor to an RGBColor
func (c *RGBAColor) ToRGB() *RGBColor {
	return &RGBColor{R: c.R, G: c.G, B: c.B}
}

// ToRGBA converts the RGBAColor to an RGBAColor
// it's here to satisfy the Color interface
func (c *RGBAColor) ToRGBA() *RGBAColor {
	return c
}

func (c *RGBAColor) IsLight() bool {

	return c.ToRGB().IsLight()
}

func (c *RGBAColor) IsDark() bool {

	return !c.IsLight()
}
