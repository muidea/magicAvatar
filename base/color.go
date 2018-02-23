package base

import "image/color"

// RGBAColor 生成RGBA
func RGBAColor(val uint32) color.RGBA {
	ret := color.RGBA{}
	ret.R = uint8(val >> 24)
	ret.G = uint8(val >> 16 & 0x00FF)
	ret.B = uint8(val >> 8 & 0x00FF)
	ret.A = uint8(val & 0x00FF)

	return ret
}
