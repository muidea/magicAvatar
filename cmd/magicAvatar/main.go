package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

var (
	port = flag.Int("port", 8020, "server listening port.")
)

func textAvatarView(res http.ResponseWriter, req *http.Request) {

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

	m.Get("/text", textAvatarView)

	m.Run()
}
