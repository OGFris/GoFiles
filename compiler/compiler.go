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
	"fmt"
	"github.com/OGFris/GoFiles/compiler/replacer"
	"os"
	"os/exec"
	"strings"
)

func init() {
	replacer.OldFiles = make(map[string][]byte)
}

func main() {

	if len(os.Args) == 3 {
		switch strings.ToLower(os.Args[1]) {
		case "build":
			file, err := os.Stat(os.Args[2])
			if err != nil {
				panic(err)
			}
			if !file.IsDir() {
				fmt.Println(replacer.Prefix + replacer.CompilerPrefix + "Your project's directory isn't working.")
				return
			}
			replacer.ReplaceAll(os.Args[2])
			cmd := exec.Command("go", "build")
			cmd.Dir = os.Args[2]
			cmd.Wait()
			if err := cmd.Run(); err != nil {
				fmt.Println(replacer.Prefix + replacer.CompilerPrefix + "something unexpected happened: " + err.Error())
			} else {
				fmt.Println(replacer.Prefix + replacer.CompilerPrefix + "Your program was compiled successfully!")
			}
			replacer.Restore()
			break
		default:
			fmt.Println(replacer.Prefix + replacer.CompilerPrefix + "Command syntax error, Usage: compiler build {YourGoProjectDirectory}")
		}
	} else {
		fmt.Println(replacer.Prefix + replacer.CompilerPrefix + "Command syntax error, Usage: compiler build {YourGoProjectDirectory}")
	}
}
