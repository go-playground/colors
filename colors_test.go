package colors

import (
	"fmt"
	"image/color"
	"path"
	"reflect"
	"runtime"
	"testing"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//

func IsEqual(t *testing.T, val1, val2 interface{}) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}

	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			goto CASE4
		} else {
			goto CASE3
		}
	} else {
		if v2 == v2Underlying {
			goto CASE2
		} else {
			goto CASE1
		}
	}

CASE1:
	// fmt.Println("CASE 1")
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
CASE2:
	// fmt.Println("CASE 2")
	return reflect.DeepEqual(v1.Interface(), v2)
CASE3:
	// fmt.Println("CASE 3")
	return reflect.DeepEqual(v1, v2.Interface())
CASE4:
	// fmt.Println("CASE 4")
	return reflect.DeepEqual(v1, v2)
}

func Equal(t *testing.T, val1, val2 interface{}) {
	EqualSkip(t, 2, val1, val2)
}

func EqualSkip(t *testing.T, skip int, val1, val2 interface{}) {

	if !IsEqual(t, val1, val2) {

		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.FailNow()
	}
}

func NotEqual(t *testing.T, val1, val2 interface{}) {
	NotEqualSkip(t, 2, val1, val2)
}

func NotEqualSkip(t *testing.T, skip int, val1, val2 interface{}) {

	if IsEqual(t, val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v should not be equal %v\n", path.Base(file), line, val1, val2)
		t.FailNow()
	}
}

func PanicMatches(t *testing.T, fn func(), matches string) {
	PanicMatchesSkip(t, 2, fn, matches)
}

func PanicMatchesSkip(t *testing.T, skip int, fn func(), matches string) {

	_, file, line, _ := runtime.Caller(skip)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%s", r)

			if err != matches {
				fmt.Printf("%s:%d Panic...  expected [%s] received [%s]", path.Base(file), line, matches, err)
				t.FailNow()
			}
		}
	}()

	fn()
}

func TestColorConversionFromHEX(t *testing.T) {

	hex, _ := ParseHEX("#5f55f5")

	Equal(t, hex.ToHEX().String(), "#5f55f5")
	Equal(t, hex.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, hex.ToRGBA().String(), "rgba(95,85,245,1)")

	hex, _ = ParseHEX("#5f5")
	Equal(t, hex.ToRGB().String(), "rgb(85,255,85)")

	hex, _ = ParseHEX("Bad Hex color!")
	Equal(t, hex, nil)
}

func TestColorConversionFromRGB(t *testing.T) {

	rgb, _ := ParseRGB("rgb(95%,85%,50%)")

	Equal(t, rgb.ToRGB().String(), "rgb(242,217,128)")
	Equal(t, rgb.ToRGBA().String(), "rgba(242,217,128,1)")
	Equal(t, rgb.ToHEX().String(), "#f2d980")

	rgb, _ = ParseRGB("rgb(95,85,245)")
	Equal(t, rgb.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, rgb.ToRGBA().String(), "rgba(95,85,245,1)")
	Equal(t, rgb.ToHEX().String(), "#5f55f5")

	rgb, _ = RGB(95, 85, 245)
	Equal(t, rgb.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, rgb.ToRGBA().String(), "rgba(95,85,245,1)")
	Equal(t, rgb.ToHEX().String(), "#5f55f5")

	rgb, _ = ParseRGB("BAD RGB COLOR")
	Equal(t, rgb, nil)

	rgb, _ = ParseRGB("rgb(95%,85%,245)")
	Equal(t, rgb, nil)
}

func TestColorConversionFromRGBA(t *testing.T) {

	rgba, _ := ParseRGBA("rgba(95%,85%,50%,1)")

	Equal(t, rgba.ToRGB().String(), "rgb(242,217,128)")
	Equal(t, rgba.ToRGBA().String(), "rgba(242,217,128,1)")
	Equal(t, rgba.ToHEX().String(), "#f2d980")

	rgba, _ = ParseRGBA("rgba(95,85,245,1)")
	Equal(t, rgba.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, rgba.ToRGBA().String(), "rgba(95,85,245,1)")
	Equal(t, rgba.ToHEX().String(), "#5f55f5")

	rgba, _ = RGBA(95, 85, 245, 1)
	Equal(t, rgba.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, rgba.ToRGBA().String(), "rgba(95,85,245,1)")
	Equal(t, rgba.ToHEX().String(), "#5f55f5")

	rgba, _ = RGBA(95, 85, 245, 6)
	Equal(t, rgba, nil)

	rgba, _ = RGBA(95, 85, 245, -1)
	Equal(t, rgba, nil)

	rgba, _ = ParseRGBA("BAD RGBA COLOR")
	Equal(t, rgba, nil)

	rgba, _ = ParseRGBA("rgba(95%,85%,245,1)")
	Equal(t, rgba, nil)
}

func TestColorConversionFromStdColor(t *testing.T) {
	rgba := FromStdColor(color.RGBA{242, 217, 128, 255})
	Equal(t, rgba.ToRGB().String(), "rgb(242,217,128)")
	Equal(t, rgba.ToRGBA().String(), "rgba(242,217,128,1)")
	Equal(t, rgba.ToHEX().String(), "#f2d980")

	rgba = FromStdColor(color.RGBA{95, 85, 245, 255})
	Equal(t, rgba.ToRGB().String(), "rgb(95,85,245)")
	Equal(t, rgba.ToRGBA().String(), "rgba(95,85,245,1)")
	Equal(t, rgba.ToHEX().String(), "#5f55f5")
}

func TestColorConversionFromToStdColor(t *testing.T) {
	// verify that colors are equals
	equalColors := func(t *testing.T, color Color, stdColor color.Color) {
		r, g, b, a := color.RGBA()
		stdR, stdG, stdB, stdA := stdColor.RGBA()
		Equal(t, r, stdR)
		Equal(t, g, stdG)
		Equal(t, b, stdB)
		Equal(t, a, stdA)
	}

	hex, _ := ParseHEX("#5f55f5")
	r, g, b, a := hex.RGBA()

	Equal(t, r, uint32(24415))
	Equal(t, g, uint32(21845))
	Equal(t, b, uint32(62965))
	Equal(t, a, uint32(65535))
	equalColors(t, hex, &color.RGBA{R: 95, G: 85, B: 245, A: 255})

	rgba, _ := RGBA(242, 217, 128, 0.4)
	r, g, b, a = rgba.RGBA()

	Equal(t, r, uint32(62194))
	Equal(t, g, uint32(55769))
	Equal(t, b, uint32(32896))
	Equal(t, a, uint32(26214))
	equalColors(t, rgba, &color.RGBA{R: 242, G: 217, B: 128, A: 102})

	rgb, _ := RGB(242, 217, 128)
	r, g, b, a = rgb.RGBA()

	Equal(t, r, uint32(62194))
	Equal(t, g, uint32(55769))
	Equal(t, b, uint32(32896))
	Equal(t, a, uint32(65535))
	equalColors(t, rgb, &color.RGBA{R: 242, G: 217, B: 128, A: 255})
}

func TestColorEqual(t *testing.T) {

	hex, _ := ParseHEX("#5f55f5")
	rgb, _ := RGB(95, 85, 245)
	rgba, _ := RGBA(95, 85, 245, 1)

	Equal(t, hex.Equal(hex), true)
	Equal(t, hex.Equal(rgb), true)
	Equal(t, hex.Equal(rgba), true)
	Equal(t, rgb.Equal(rgb), true)
	Equal(t, rgb.Equal(hex), true)
	Equal(t, rgb.Equal(rgba), true)
	Equal(t, rgba.Equal(rgba), true)
	Equal(t, rgba.Equal(rgb), true)
	Equal(t, rgba.Equal(hex), true)

	hex2, _ := ParseHEX("#5f55f4")
	rgb2, _ := RGB(95, 87, 245)
	rgba2, _ := RGBA(93, 85, 245, 1)

	Equal(t, hex2.Equal(rgb2), false)
	Equal(t, hex2.Equal(rgba2), false)
	Equal(t, rgb2.Equal(hex2), false)
	Equal(t, rgb2.Equal(rgba2), false)
	Equal(t, rgba2.Equal(rgb2), false)
	Equal(t, rgba2.Equal(hex2), false)

}

func TestParseColor(t *testing.T) {

	color, _ := Parse("#FFF")
	NotEqual(t, color, nil)
	Equal(t, reflect.TypeOf(color) == reflect.TypeOf(&HEXColor{}), true)

	color, _ = Parse("rgb(95,85,245)")
	NotEqual(t, color, nil)
	Equal(t, reflect.TypeOf(color) == reflect.TypeOf(&RGBColor{}), true)

	color, _ = Parse("rgba(95,85,245,1)")
	NotEqual(t, color, nil)
	Equal(t, reflect.TypeOf(color) == reflect.TypeOf(&RGBAColor{}), true)

	color, _ = Parse("#ff")
	Equal(t, color, nil)

	color, _ = Parse("garbage-data")
	Equal(t, color, nil)

	c, err := Parse("rgba(127,34,94,0.534556634531)")
	Equal(t, err, nil)
	Equal(t, reflect.TypeOf(c) == reflect.TypeOf(&RGBAColor{}), true)
}

func TestIsLightIsDark(t *testing.T) {

	rgb, _ := RGB(0, 0, 0)
	Equal(t, rgb.IsLight(), false)
	Equal(t, rgb.IsDark(), true)

	rgb, _ = RGB(255, 255, 255)
	Equal(t, rgb.IsLight(), true)
	Equal(t, rgb.IsDark(), false)

	rgba, _ := RGBA(0, 0, 0, 1)
	Equal(t, rgba.IsLight(), false)
	Equal(t, rgba.IsDark(), true)

	rgba, _ = RGBA(255, 255, 255, 1)
	Equal(t, rgba.IsLight(), true)
	Equal(t, rgba.IsDark(), false)

	hex, _ := ParseHEX("#99FF33")
	Equal(t, hex.IsLight(), true)
	Equal(t, hex.IsDark(), false)

	hex, _ = ParseHEX("#3300FF")
	Equal(t, hex.IsLight(), false)
	Equal(t, hex.IsDark(), true)
}

func TestIsLightAlphaIsDarkAlpha(t *testing.T) {

	bg, _ := RGB(255, 255, 255)

	rgba, _ := RGBA(0, 0, 0, 1)
	Equal(t, rgba.IsLightAlpha(bg), false)
	Equal(t, rgba.IsDarkAlpha(bg), true)

	rgba, _ = RGBA(0, 0, 0, 0)
	Equal(t, rgba.IsLightAlpha(bg), true)
	Equal(t, rgba.IsDarkAlpha(bg), false)

	rgba, _ = RGBA(255, 255, 255, 1)
	Equal(t, rgba.IsLightAlpha(bg), true)
	Equal(t, rgba.IsDarkAlpha(bg), false)

	rgba, _ = RGBA(0, 0, 0, 0.5)
	Equal(t, rgba.IsLightAlpha(bg), false)
	Equal(t, rgba.IsDarkAlpha(bg), true)

	rgba, _ = RGBA(0, 0, 0, 0.3)
	Equal(t, rgba.IsLightAlpha(bg), true)
	Equal(t, rgba.IsDarkAlpha(bg), false)

	rgba, _ = RGBA(240, 100, 20, 0.5)
	Equal(t, rgba.IsLightAlpha(bg), true)
	Equal(t, rgba.IsDarkAlpha(bg), false)

	bg, _ = RGB(0, 0, 0)

	rgba, _ = RGBA(132, 100, 50, 0.5)
	Equal(t, rgba.IsLightAlpha(bg), false)
	Equal(t, rgba.IsDarkAlpha(bg), true)

	rgba, _ = RGBA(132, 100, 50, 0.7)
	Equal(t, rgba.IsLightAlpha(bg), false)
	Equal(t, rgba.IsDarkAlpha(bg), true)
}

func TestInterfaceTypes(t *testing.T) {

	fn := func(c Color) string {

		if c == nil {
			return ""
		}

		c.IsDark()
		c.IsLight()

		return c.String()
	}

	hex, _ := ParseHEX("#FFF")
	rgb, _ := ParseRGB("rgb(95,85,245)")
	rgba, _ := ParseRGBA("rgba(95,85,245,1)")

	fn(hex)
	fn(rgb)
	fn(rgba)
}

func BenchmarkSpeed(b *testing.B) {

	for n := 0; n < b.N; n++ {
		h, _ := ParseHEX("#FFFFFF")
		h.ToRGBA()
	}
}
