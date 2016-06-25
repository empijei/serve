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
      <input type="hidden" name="token" value="{{.Token}}"/>
      <input type="submit" value="upload" />
</form>
</body>
</html>`

var tokens map[string]struct{}

func init() {
	tokens = make(map[string]struct{})
}

func UploaderEndpoint(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := make([]byte, 8)
			_, err := rand.Read(data)

			t := template.Must(template.New("uploadform").Parse(form))
			fill := struct {
				Token    string
				Endpoint string
			}{}
			fill.Token = fmt.Sprintf("%x", data)
			fill.Endpoint = path
			//TODO use a cookie instead
			tokens[fill.Token] = struct{}{}
			err = t.Execute(w, fill)
			if err != nil {
				log.Println(err)
			}
		} else {
			_ = r.ParseMultipartForm(32 << 20)
			if _, ok := tokens[r.FormValue("token")]; !ok {
				_, _ = w.Write([]byte("<h1>NOPE</h1>"))
				return
			}
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
				fmt.Fprintf(w, "Errors occourred")
				log.Println(err)
				return
			}
			fmt.Fprintf(w, "<h1> Uploaded %d bytes</h1>", n)
		}
	}
}