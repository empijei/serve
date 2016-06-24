package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var localport = flag.String("p", "8080", "The port to run on")
var localfolder = flag.String("f", "", "The folder to serve (default PWD)")

func init() {
	flag.Parse()
	flag.Arg(1)
	if !strings.HasPrefix(*localport, ":") {
		*localport = ":" + *localport
	}
}
func main() {
	err := DirList(*localfolder, "/")
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(*localport, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func DirList(localfolder, webpath string) (err error) {
	fs := http.FileServer(http.Dir(localfolder))
	wd, err := os.Getwd()
	http.Handle(webpath, http.StripPrefix(webpath, fs))
	fmt.Printf("Serving local folder %v on http://localhost%s%s\nPress Control+C to stop.\n", wd+string(os.PathSeparator)+localfolder, *localport, webpath)
	return
}

//TODO add file upload!
