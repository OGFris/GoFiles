package main

import (
	"bufio"
	"fmt"
	"github.com/OGFris/GoFiles"
	"github.com/OGFris/GoFiles/server"
	"os"
)

func main() {
	server.Instance.Start("8888")
	html := GoFiles.NewFile("./html/index.html")
	css := GoFiles.NewFile("./css/index.css")
	js := GoFiles.NewFile("./js/index.js")

	html.SetName("")
	html.Host(false)
	html.SetName("index.html")
	html.Host(false)

	css.SetName("css/index.css")
	css.Host(false)
	js.SetName("js/index.js")
	js.Host(false)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press ENTER to close the server.")
	scanner.Scan()
	server.Instance.Http.Close()
}
