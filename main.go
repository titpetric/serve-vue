package main

import (
	"flag"
	"log"
	"os"
	"path"

	"net/http"
)

// Serves index.html in case the requested file isn't found (or some other os.Stat error)
func serveIndex(assetPath string, serve http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := path.Join(assetPath, "index.html")
		requestedPage := path.Join(assetPath, r.URL.Path)
		_, err := os.Stat(requestedPage)
		if err != nil {
			http.ServeFile(w, r, indexPage)
			return
		}
		serve.ServeHTTP(w, r)
	}
}

func main() {
	var (
		port     = flag.String("port", "8080", "Port for server")
		contents = flag.String("contents", ".", "Folder for display")
	)
	flag.Parse()

	if folder := flag.Arg(0); folder != "" {
		*contents = folder
	}

	// Set absolute path to contents folder
	cwd, _ := os.Getwd()
	contentsPath := path.Join(cwd, *contents)

	http.HandleFunc("/", serveIndex(contentsPath, http.FileServer(http.Dir(contentsPath))))

	log.Println("Started listening on port", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}
