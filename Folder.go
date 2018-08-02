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
	"encoding/json"
)

type Folder struct {
	Path    string   `json:"path"`
	Files   []File   `json:"files"`
	Folders []Folder `json:"folders"`
}

// AddFolder adds a virtual folder to your Folder
// Shouldn't be used unless you know what you're doing.
func (f *Folder) AddFolder(fadd Folder) {
	f.SetFolders(append(f.GetFolders(), fadd))
}

// AddFile adds a virtual file to your Folder
// Shouldn't be used unless you know what you're doing.
func (f *Folder) AddFile(file File) {
	f.SetFiles(append(f.GetFiles(), file))
}

// GetFolders gives you all the folders on your Folder
func (f *Folder) GetFolders() []Folder {
	return f.Folders
}

// SetFolders sets all the folders on your Folder.
// Shouldn't be used unless you know what you're doing.
func (f *Folder) SetFolders(folders []Folder) {
	f.Folders = folders
}

// GetFiles gives you all the files on your Folder
func (f *Folder) GetFiles() []File {
	return f.Files
}

// SetFiles sets all the files on your Folder.
// Shouldn't be used unless you know what you're doing.
func (f *Folder) SetFiles(files []File) {
	f.Files = files
}

// Gives you the path of your folder.
func (f *Folder) GetPath() string {
	return f.Path
}

// SetPath sets the path of your Folder.
// Shouldn't be used unless you know what you're doing.
func (f *Folder) SetPath(path string) {
	f.Path = path
}

func NewFolder(path string) *Folder {
	return &Folder{Path: path}
}

// SetFolder is a function ran by the compiler.
// Shouldn't be used unless you know what you're doing.
func (f *Folder) SetFolder(folder string) *Folder {
	var v Folder
	json.Unmarshal([]byte(folder), &v)
	f.SetFiles(v.GetFiles())
	f.SetFolders(v.GetFolders())
	return f
}
