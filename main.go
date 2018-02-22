package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var (
	dpi        = flag.Float64("dpi", 50, "screen resolution in Dots Per Inch")
	fontfile   = flag.String("fontfile", "./font.ttf", "filename of the ttf font")
	avatarfile = flag.String("avatarfile", "./avatar.png", "filename of the avatar file")
	text       = flag.String("text", "Me", "the avatar content")
	hinting    = flag.String("hinting", "full", "none | full")
	size       = flag.Float64("size", 40, "font size in points")
	width      = flag.Int("width", 64, "avatar width")
	hight      = flag.Int("hight", 64, "avatar hight")
)

func main() {
	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, *width, *hight))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the text.
	fontSize := int(c.PointToFixed(*size) >> 6)
	xOffset := (*width - fontSize) / 2
	yOffset := (*hight + fontSize) / 2
	pt := freetype.Pt(xOffset, yOffset)
	_, err = c.DrawString(*text, pt)
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
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
}
