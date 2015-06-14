package colors

import (
	"errors"
	"strings"
)

var (
	// ErrBadColor is the default bad color error
	ErrBadColor = errors.New("Parsing of color failed, Bad Color")
)

// Color is the base color interface from which all others ascribe to
type Color interface {
	ToHEX() *HEXColor
	ToRGB() *RGBColor
	ToRGBA() *RGBAColor
	String() string
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
