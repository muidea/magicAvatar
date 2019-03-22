package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/muidea/magicAvatar/factory"
)

var (
	port = flag.Int("port", 8020, "server listening port.")
)

func textAvatarView(res http.ResponseWriter, req *http.Request) {
	text := req.URL.Query().Get("text")
	fontSizeVal := req.URL.Query().Get("fontSize")
	widthVal := req.URL.Query().Get("width")
	hightVal := req.URL.Query().Get("hight")
	backColorVal := req.URL.Query().Get("backColor")
	fontColorVal := req.URL.Query().Get("fontColor")

	log.Print(req.RemoteAddr)

	avatarFile := "./avatar.png"
	newFileName := avatarFile
	for true {

		newFileName = fmt.Sprintf("%s_avatar.png", req.RemoteAddr)
		fontSize, err := strconv.ParseFloat(fontSizeVal, 64)
		if err != nil {
			log.Printf("illegal fontSize value, val:%s, err:%s", fontSizeVal, err.Error())
			fontSize = 40.00
		}

		width, err := strconv.ParseInt(widthVal, 0, 32)
		if err != nil {
			log.Printf("illegal width value, val:%s, err:%s", widthVal, err.Error())
			width = 64
		}

		hight, err := strconv.ParseInt(hightVal, 0, 32)
		if err != nil {
			log.Printf("illegal hight value, val:%s, err:%s", hightVal, err.Error())
			hight = 64
		}

		backColor, err := strconv.ParseUint(backColorVal, 0, 32)
		if err != nil {
			log.Printf("illegal backColor value, val:%s, err:%s", backColorVal, err.Error())
			backColor = 0xFFFFFFFF
		}

		fontColor, err := strconv.ParseUint(fontColorVal, 0, 32)
		if err != nil {
			log.Printf("illegal fontColor value, val:%s, err:%s", fontColorVal, err.Error())
			fontColor = 0x000000FF
		}

		if !factory.MakeTextAvatar(text, newFileName, fontSize, uint32(backColor), uint32(fontColor), int(width), int(hight)) {
			log.Print("make text avatar failed")
			break
		}

		avatarFile = newFileName
		break
	}

	{
		filePath, fileName := path.Split(avatarFile)
		dir := http.Dir(filePath)
		file, err := dir.Open(fileName)
		if err != nil {
			log.Printf("open file failed, fileName:%s, err:%s", fileName, err.Error())
			return
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil || fi.IsDir() {
			log.Printf("fetch file stat failed, fileName:%s, err:%s", fileName, err.Error())
			return
		}

		http.ServeContent(res, req, "", fi.ModTime(), file)
	}

	log.Print(newFileName)
	os.Remove(newFileName)
}

func mainView(res http.ResponseWriter, req *http.Request) {

}

func main() {
	flag.Parse()

	os.Stdout.WriteString("=================================\n")
	os.Stdout.WriteString("MagicAvatar Server V1.0\n")
	os.Stdout.WriteString("Author:rangh\n")
	os.Stdout.WriteString("EMail:rangh@foxmail.com\n")
	os.Stdout.WriteString("=================================\n")
	svrPort := fmt.Sprintf("%d", *port)
	os.Setenv("PORT", svrPort)

	m := martini.Classic()

	m.Get("/textAvatar", textAvatarView)

	m.Run()
}
