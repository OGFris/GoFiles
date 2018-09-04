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

package GoFiles

import (
	"github.com/OGFris/GoFiles/server"
)

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Contents string `json:"contents"`
}

// NewFile Creates a new File instance.
func NewFile(path string) *File {
	return &File{Path: path}
}

// SetContent sets the content of the virtual file.
// Don't recommend using it because trying to edit a virtually compiled file is pointless.
func (f *File) SetContents(content string) *File {
	f.Contents = content
	return f
}

// Host adds the file to the http server.
func (f *File) Host(private bool, contentType string) {
	server.AddFile(f.Name, []byte(f.Contents), private, contentType)
}

// Url gives the http url for the file.
func (f *File) Url() string {
	return "/" + f.Name
}

// Hosted gives a boolean whether the file is hosted on the http server or not.
func (f *File) Hosted() bool {
	return server.Routes[f.Name].Url != ""
}
