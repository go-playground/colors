package colors

import "gopkg.in/bluesuncorp/validator.v5"

var (
	validate = validator.New("validate", validator.BakedInValidators)
)

type Color interface {
	ToHEX() *HEXColor
	ToRGB() *RGBColor
	ToRGBA() *RGBAColor
	String() string
}
