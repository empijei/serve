package lib

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var form = `<html>
<head>
    <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="{{.Endpoint}}" method="post">
      <input type="file" name="uploadfile" />
      <input type="submit" value="upload" />
</form>
</body>
</html>`

// UploaderEndpoint handles file uploading.
// It responds to GET requests with the file upload form, and to POST
// requests with the actual uploading.
func UploaderEndpoint(name, path, webpath string) http.HandlerFunc {
	webpath = name + webpath
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := make([]byte, 8)
			_, err := rand.Read(data)

			t := template.Must(template.New("uploadform").Parse(form))
			fill := struct {
				Endpoint string
			}{}
			fill.Endpoint = path
			//TODO use a cookie instead
			err = t.Execute(w, fill)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == "POST" {
			_ = r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() { _ = file.Close() }()
			f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer func() { _ = f.Close() }()
			n, err := io.Copy(f, file)
			if err != nil {
				fmt.Fprintf(w, "Errors occurred")
				log.Println(err)
				return
			}
			fmt.Fprintf(w, "<h1> Uploaded %d bytes</h1><a href='"+webpath+"'>Back to dirlist</a>", n)
		} else {
			http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
		}
	}
}
