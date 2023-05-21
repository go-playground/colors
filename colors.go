package colors

import (
	"errors"
	"strings"
)

var (
	// ErrBadColor is the default bad color error
	ErrBadColor = errors.New("parsing of color failed, Bad Color")
)

// Color is the base color interface from which all others ascribe to
type Color interface {
	// ToHEX converts the Color interface to a concrete HEXColor
	ToHEX() *HEXColor

	// ToRGB converts the Color interface to a concrete RGBColor
	ToRGB() *RGBColor

	// ToRGBA converts the Color interface to a concrete RGBAColor
	ToRGBA() *RGBAColor

	// String returns the string representation of the Color
	String() string

	// IsLight returns whether the color is perceived to be a light color
	// http://stackoverflow.com/a/24213274/3158232 and http://www.nbdtech.com/Blog/archive/2008/04/27/Calculating-the-Perceived-Brightness-of-a-Color.aspx
	IsLight() bool

	// IsDark returns whether the color is perceived to be a dark color
	//for perceived luminance, not strict math
	IsDark() bool

	// RGBA implements std-lib color.Color interface.
	// It returns the red, green, blue and alpha values for the color. Each value ranges within [0, 0xffff]
	RGBA() (r, g, b, a uint32)

	// Equal reports whether the colors are the same
	Equal(Color) bool
}

// Parse parses an unknown color type to it's appropriate type, or returns a ErrBadColor
func Parse(s string) (Color, error) {

	if len(s) < 4 {
		return nil, ErrBadColor
	}

	s = strings.ToLower(s)

	if s[:1] == "#" {
		return ParseHEX(s)
	} else if s[:4] == "rgba" {
		return ParseRGBA(s)
	} else if s[:3] == "rgb" {
		return ParseRGB(s)
	}

	return nil, ErrBadColor
}
