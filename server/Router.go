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
