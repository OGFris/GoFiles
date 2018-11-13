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

package replacer

import (
	"encoding/json"
	"fmt"
	"github.com/OGFris/GoFiles"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const (
	Prefix         = "[GoFiles] "
	CompilerPrefix = "[Compiler] "
)

var OldFiles map[string][]byte

func Restore() {
	for path, content := range OldFiles {
		ioutil.WriteFile(path, content, 0777)
	}
}

func ReplaceAll(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".go") {
			replace(folder + "/" + file.Name())
		}
	}
}

func replace(file string) {
	f, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	if f.IsDir() {
		ReplaceAll(file)
	} else {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		OldFiles[file] = data
		r := string(data)
		r = replaceFolder(r)
		r = replaceFile(r)
		ioutil.WriteFile(file, []byte(r), 0777)
	}
}

func replaceFolder(code string) (new string) {
	for line, text := range strings.Split(code, "\n") {
		if strings.Contains(text, "GoFiles.NewFolder") {
			arg, err := getArg(text)
			if err != nil {
				fmt.Println(Prefix + CompilerPrefix + "Found an error at the line " + fmt.Sprint(line) + ": " + err.Error())
			}
			folder := GoFiles.Folder{}
			scanDir(arg, &folder)
			data, err := json.Marshal(folder)
			if err != nil {
				panic(err)
			}
			new += text + ".SetFolder(`" + string(data) + "`)\n"
		} else {
			new += text + "\n"
		}
	}
	return
}

func replaceFile(code string) (new string) {
	for line, text := range strings.Split(code, "\n") {
		if strings.Contains(text, "GoFiles.NewFile") {
			arg, err := getArg(text)
			if err != nil {
				fmt.Println(Prefix + CompilerPrefix + "Found an error at the line " + fmt.Sprint(line) + ": " + err.Error())
			}
			bytes, err := ioutil.ReadFile(arg)
			if err != nil {
				panic(err)
			}
			new += text + ".SetContents(`" + string(bytes) + "`)\n"
		} else {
			new += text + "\n"
		}
	}
	return
}

func getArg(code string) (arg string, err error) {
	var (
		s   = false
		str = false
		n   = 0
		txt = strings.Split(code, "")
	)
	for i, c := range txt {
		if i > 0 {
			if s && !str {
				if c == ")" && n == 0 {
					err = errors.New("file or folder directory must be set")
					return
				}
				if c == ")" {
					for i2 := n; i2 < (i - 1); i2++ {
						arg += txt[i2]
					}
					return
				}
			}
			if s {
				if c != "\"" && str {
					continue
				}
				if c == "\"" {
					if str == true {
						str = false
					} else {
						str = true
					}
				}
			}
			if c == "(" && !s {
				s = true
				n = i + 2
			}
		} else {
			if c == "(" {
				err = errors.New("'(' can't be at the beginning of the line")
				return
			}
		}
	}
	if s == false {
		err = errors.New("missing '(' at the beginning of the function call")
	} else {
		// This should never be executed unless there is something wrong with the code
		err = errors.New("function isn't called correctly")
	}
	return
}

func scanDir(path string, f *GoFiles.Folder) {
	file, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !file.IsDir() {
		err := errors.New(path + " isn't a valid directory")
		panic(err)
	}
	files, _ := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file = range files {
		if file.IsDir() {
			folder := GoFiles.Folder{}
			scanDir(path+"/"+file.Name(), &folder)
			f.AddFolder(folder)
		} else {
			bytes, err := ioutil.ReadFile(path + "/" + file.Name())
			if err != nil {
				panic(err)
			}
			content := string(bytes)
			f.AddFile(GoFiles.File{Path: path + "/" + file.Name(), Contents: content})
		}
	}
}
