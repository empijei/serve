package main

import (
	"flag"
	"fmt"
	"log"
	"net"
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

func Upload(localfolder, webpath string) (err error) {
	http.HandleFunc(webpath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html>
		<title>Go upload</title>
		<body>
		<form action="http://{{myip}}{{*localport}}/xyzzyupload" method="post" enctype="multipart/form-data">
		<label for="file">Filename:</label>
		<input type="file" name="file" id="file">
		<input type="submit" name="submit" value="Submit">
		</form>
		
		</body>
		</html>
		`)
	})
	return
}

var ip string

func myip() string {
	if ip != "" {
		return ip
	}
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {
		log.Println(err)
		return "localhost"
	}
	for _, i := range ifaces {
		if i.Name == "lo" && len(ifaces) > 1 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			return ip.String()
		}
	}
	return "localhost"
}
