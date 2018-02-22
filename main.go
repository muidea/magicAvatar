package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi        = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	backcolor  = flag.Uint("backcolor", 0xFFFF0000, "backcolor for avatar")
	fontcolor  = flag.Uint("fontcolor", 0xFF0000FF, "fontcolor for avatar")
	fontfile   = flag.String("fontfile", "./font.ttf", "filename of the ttf font")
	avatarfile = flag.String("avatarfile", "./avatar.png", "filename of the avatar file")
	text       = flag.String("text", "", "the avatar content")
	size       = flag.Float64("size", 40, "font size in points")
	width      = flag.Int("width", 64, "avatar width")
	hight      = flag.Int("hight", 64, "avatar hight")
)

func rgbaColor(val uint32) color.RGBA {
	ret := color.RGBA{}
	ret.R = uint8(val >> 24)
	ret.G = uint8(val >> 16 & 0x00FF)
	ret.B = uint8(val >> 8 & 0x00FF)
	ret.A = uint8(val & 0x00FF)

	return ret
}

func main() {
	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	fontLibrary, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	avatarVal := []rune(*text)
	if len(avatarVal) > 1 {
		avatarVal = avatarVal[0:1]
	}

	// Initialize the context.
	frontGround, backGround := image.Black, image.White
	frontGround = image.NewUniform(rgbaColor(uint32(*fontcolor)))
	backGround = image.NewUniform(rgbaColor(uint32(*backcolor)))

	rgba := image.NewRGBA(image.Rect(0, 0, *width, *hight))
	draw.Draw(rgba, rgba.Bounds(), backGround, image.ZP, draw.Src)

	fontContext := freetype.NewContext()
	fontContext.SetDPI(*dpi)
	fontContext.SetFont(fontLibrary)
	fontContext.SetFontSize(*size)
	fontContext.SetClip(rgba.Bounds())
	fontContext.SetDst(rgba)
	fontContext.SetSrc(frontGround)
	fontContext.SetHinting(font.HintingFull)

	// Draw the text.
	fontSize := int(fontContext.PointToFixed(*size) >> 6)
	xOffset := (*width - fontSize) / 2
	yOffset := (*hight+fontSize)/2 - fontSize/10

	if !unicode.Is(unicode.Scripts["Han"], avatarVal[0]) {
		xOffset = (*width-fontSize)/2 + fontSize/6
		yOffset = (*hight+fontSize)/2 - fontSize/8
	}

	pt := freetype.Pt(xOffset, yOffset)
	_, err = fontContext.DrawString(string(avatarVal), pt)
	if err != nil {
		log.Println(err)
		return
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create(*avatarfile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	bufferWriter := bufio.NewWriter(outFile)
	err = png.Encode(bufferWriter, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = bufferWriter.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
