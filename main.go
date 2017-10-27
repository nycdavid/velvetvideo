package main

import (
	"fmt"
	"github.com/spf13/afero"
)

func main() {
	AppFs := afero.NewOsFs()
	files, _ := afero.Glob(AppFs, "files/*")

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
