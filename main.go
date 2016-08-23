package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/empijei/serve/lib"
)

var localport = flag.String("p", "8080", "The port to run on")
var localfolder = flag.String("f", "", "The folder to serve (default PWD)")
var uploadpath = flag.String("u", "xyzzy/", "The webpath where the upload form is hosted")
var webpath = flag.String("w", "/", "The root webpath")
var tls = flag.Bool("s", false, "Use TLS")

func init() {
	flag.Parse()
	flag.Arg(1)
	if !strings.HasPrefix(*localport, ":") {
		*localport = ":" + *localport
	}
}

func main() {
	fmt.Println("Serve is availble at https://github.com/empijei/serve")
	fmt.Println("Local ip address: " + lib.MyIP().String())
	err := lib.DirList(*localfolder, *webpath)
	if err != nil {
		log.Fatal(err)
	}
	wd, err := os.Getwd()
	fmt.Printf("Serving local folder %v on \"%s\"\n", wd+string(os.PathSeparator)+*localfolder, Name())

	name := Name()
	http.HandleFunc(*webpath+*uploadpath, lib.UploaderEndpoint(name, *uploadpath, *webpath))
	fmt.Printf("File Upload form is available at \"%s%s\"\n", name, *uploadpath)

	fmt.Println("Press Control+C to stop")
	if !*tls {
		err = http.ListenAndServe(*localport, nil)
	} else {
		certpath := os.Getenv("HOME") + "/.servepj/"
		err = http.ListenAndServeTLS(*localport, certpath+"cert.pem", certpath+"key.pem", nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Name returns the full URL for the local port and webpath passed from the command line
func Name() string {
	b := bytes.Buffer{}
	_, _ = b.WriteString("http")
	if *tls {
		_, _ = b.WriteString("s")
	}
	_, _ = b.WriteString("://")
	_, _ = b.WriteString(lib.MyName())
	_, _ = b.WriteString(*localport)
	_, _ = b.WriteString(*webpath)
	return b.String()
}
