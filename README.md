# colors
Go color manipulation, conversion and printing library/utility

this library is currently in development, not all color types such as HSL, HSV and CMYK will be included in the first release; pull requests are welcome.

Usage and documentation
=======================

Please see ... for detailed usage docs.

#Example
```go
hex, err := colors.ParseHex("#fff")
rgb, err := colors.ParseRGB("rgb(0,0,0)")
rgb, err := colors.RGB(0,0,0)
rgba, err := colors.ParseRGBA("rgba(0,0,0,1)")
rgba, err := colors.RGBA(0,0,0,1)

color, err := colors.Parse("#000)

color.ToRGB()   // rgb(0,0,0)
color.ToRGBA()  // rgba(0,0,0,1)
color.ToHEX()   // #000000
color.IsLight() // false
color.IsDark()  // true

```

How to Contribute
=================

There will always be a development branch for each version i.e. `v1-development`. In order to contribute, 
please make your pull requests against those branches.

If the changes being proposed or requested are breaking changes, please create an issue, for discussion 
or create a pull request against the highest development branch for example this package has a 
v1 and v1-development branch however, there will also be a v2-development brach even though v2 doesn't exist yet.

I am not a color expert by any means and am sure that there could be better or even more efficient
ways to accomplish the color conversion and so forth and I welcome any suggestions or pull request to help!

License
=======
Distributed under MIT License, please see license file in code for more details.
