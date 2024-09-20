package server

import (
	"fmt"
	"html"
	"net/http"
	"path"

	"github.com/NYTimes/gziphandler"
)

type handlerType func(pathToStaticWebsite string) http.HandlerFunc

func handlerFuncWrapper(pathToStaticWebsite string, handler handlerType) http.Handler {
	return gziphandler.GzipHandler(http.HandlerFunc(handler(pathToStaticWebsite)))
}

func fileServerWrapper(pathToStaticWebsite string, pathToServe string) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf("/%s", pathToServe),
		gziphandler.GzipHandler(http.FileServer(http.Dir(path.Join(pathToStaticWebsite, pathToServe)))),
	)
}

// Route handler
func pageHandler(pathToStaticWebsite string, expectedPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := html.EscapeString(r.URL.Path)
		if urlPath == expectedPath {
			http.ServeFile(w, r, path.Join(pathToStaticWebsite, "index.html"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func homeHandler(pathToStaticWebsite string) http.HandlerFunc {
	return pageHandler(pathToStaticWebsite, "/")
}

func writingsHandler(pathToStaticWebsite string) http.HandlerFunc {
	return pageHandler(pathToStaticWebsite, "/writings")
}
