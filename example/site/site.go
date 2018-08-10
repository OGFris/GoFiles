//  MIT License
//
//  Copyright (c) 2018 Fris
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.

package main

import (
	"bufio"
	"fmt"
	"github.com/OGFris/GoFiles"
	"github.com/OGFris/GoFiles/server"
	"os"
)

func main() {
	//If the server isn't started, it will be automatically hosted on port 80.
	server.Instance.Start("8888")
	html := GoFiles.NewFile("./html/index.html")
	css := GoFiles.NewFile("./css/index.css")
	js := GoFiles.NewFile("./js/index.js")

	html.Name = ""
	html.Host(false, server.ContentTypeHtml)
	html.Name = "index.html"
	html.Host(false, server.ContentTypeHtml)

	css.Name = "css/index.css"
	css.Host(false, server.ContentTypeCss)
	js.Name = "js/index.js"
	js.Host(false, server.ContentTypePlain)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press ENTER to close the server.")
	scanner.Scan()
	server.Instance.Http.Close()
}
