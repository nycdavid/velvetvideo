package main

import (
	"fmt"
	"net/http"

  "github.com/spf13/afero"
)

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Listening on port 1323")
  http.ListenAndServe(":1323", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  AppFs := afero.NewOsFs()
  files, _ := afero.Glob(AppFs, "files/*")
  fmt.Fprintf(w, files[0])
}
