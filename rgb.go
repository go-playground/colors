package colors

import "fmt"

const (
	rgbString = "rgb(%d,%d,%d)"
)

type RGBColor struct {
	R uint8
	G uint8
	B uint8
}

// func HEX(s string) *HEXColor,error {

// 	return nil, errors.New("ERROR BAD COLOR")
// 	// h := new(HEXColor)
// 	// h.hex = s

// 	// return h
// // }

func (c *RGBColor) String() string {
	return fmt.Sprintf(rgbString, c.R, c.G, c.B)
}

// func (c *RGBColor255) ToHEX() *RGBColor255 {
// 	return c
// }

// func (c *RGBColor255) ToRGB() *RGBColor255 {
// 	return new(RGBColor255)
// }

// func (c *RGBColor255) ToRGBA() *RGBColor255 {
// 	return new(RGBAColor255)
// }
