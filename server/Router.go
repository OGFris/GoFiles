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

package server

import (
	"net/http"
)

type Router struct {
}

type Route struct {
	Url      string
	Contents []byte
	Private  bool
}

var Routes map[string]Route

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	found := false
	for _, route := range Routes {
		if route.Url == path {
			found = true
			if route.Private {
				if req.Header.Get("x-forwarded-for") == "" {
					w.Write(route.Contents)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				w.Write(route.Contents)
			}
			break
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
}

func AddFile(name string, contents []byte, private bool) {
	if !Instance.Running {
		Instance.Start("8080")
	}
	Routes[name] = Route{Url: "/" + name, Contents: contents, Private: private}
}
