package colors

import (
	"fmt"
	"strings"

	"gopkg.in/bluesuncorp/validator.v5"
)

const (
	hexFormat      = "#%02x%02x%02x"
	hexShortFormat = "#%1x%1x%1x"
	hexToRGBFactor = 17
)

type HEXColor struct {
	hex string
}

func HEX(s string) (*HEXColor, *validator.FieldError) {

	if err := validate.Field(s, "hexcolor"); err != nil {
		return nil, err
	}

	return &HEXColor{hex: strings.ToLower(s)}, nil
}

func (c *HEXColor) String() string {
	return c.hex
}

func (c *HEXColor) ToHEX() *HEXColor {
	return c
}

func (c *HEXColor) ToRGB() *RGBColor {

	var r, g, b uint8

	if len(c.hex) == 4 {
		fmt.Sscanf(c.hex, hexShortFormat, &r, &g, &b)
		r *= hexToRGBFactor
		g *= hexToRGBFactor
		b *= hexToRGBFactor
	} else {
		fmt.Sscanf(c.hex, hexFormat, &r, &g, &b)
	}

	return &RGBColor{R: r, G: g, B: b}
}

func (c *HEXColor) ToRGBA() *RGBAColor {

	rgb := c.ToRGB()

	return &RGBAColor{R: rgb.R, G: rgb.G, B: rgb.B, A: 1}
}
