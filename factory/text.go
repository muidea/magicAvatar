package factory

import (
	"bufio"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"github.com/muidea/magicAvatar/base"
)

func MakeTextAvatar(text, avatarFile string, fontSize float64, backColor, fontColor uint32, width, hight int) bool {
	dpi := 72.0
	fontFile := "./font.ttf"

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return false
	}
	fontLibrary, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return false
	}

	avatarVal := []rune(text)
	if len(avatarVal) > 1 {
		avatarVal = avatarVal[0:1]
	}

	// Initialize the context.
	frontGround, backGround := image.Black, image.White
	frontGround = image.NewUniform(base.RGBAColor(uint32(fontColor)))
	backGround = image.NewUniform(base.RGBAColor(uint32(backColor)))

	rgba := image.NewRGBA(image.Rect(0, 0, width, hight))
	draw.Draw(rgba, rgba.Bounds(), backGround, image.ZP, draw.Src)

	fontContext := freetype.NewContext()
	fontContext.SetDPI(dpi)
	fontContext.SetFont(fontLibrary)
	fontContext.SetFontSize(fontSize)
	fontContext.SetClip(rgba.Bounds())
	fontContext.SetDst(rgba)
	fontContext.SetSrc(frontGround)
	fontContext.SetHinting(font.HintingFull)

	// Draw the text.
	reallyfontSize := int(fontContext.PointToFixed(fontSize) >> 6)
	xOffset := (width - reallyfontSize) / 2
	yOffset := (hight+reallyfontSize)/2 - reallyfontSize/10

	if len(avatarVal) > 0 && !unicode.Is(unicode.Scripts["Han"], avatarVal[0]) {
		xOffset = (width-reallyfontSize)/2 + reallyfontSize/6
		yOffset = (hight+reallyfontSize)/2 - reallyfontSize/8
	}

	pt := freetype.Pt(xOffset, yOffset)
	_, err = fontContext.DrawString(string(avatarVal), pt)
	if err != nil {
		log.Println(err)
		return false
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create(avatarFile)
	if err != nil {
		log.Println(err)
		return false
	}
	defer outFile.Close()
	bufferWriter := bufio.NewWriter(outFile)
	err = png.Encode(bufferWriter, rgba)
	if err != nil {
		log.Println(err)
		return false
	}
	err = bufferWriter.Flush()
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
