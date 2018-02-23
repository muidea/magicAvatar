package main

import (
	"flag"
	"log"

	"muidea.com/magicAvatar/factory"
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

func main() {
	flag.Parse()

	if !factory.MakeTextAvatar(*text, *avatarfile, *size, uint32(*backcolor), uint32(*fontcolor), *width, *hight) {
		log.Print("make test avatar failed.")
	}
}
