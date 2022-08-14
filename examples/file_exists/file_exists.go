package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/albertoaer/nipper"
	"github.com/gin-gonic/gin"
)

type File struct {
	Name string `param:"name"`
	Path string
}

type FileResult struct {
	Error  string `json:"error"`
	Result File   `json:"file"`
}

func main() {
	router := gin.Default()
	api := nipper.NewHttpRouter(router.Group("/api"))
	api.Route(":name").Get(func(f File) FileResult {
		if _, err := os.Stat(f.Name); !errors.Is(err, os.ErrNotExist) {
			f.Path, _ = filepath.Abs(f.Name)
			return FileResult{"", f}
		}
		return FileResult{"File does not exists", f}
	})
	router.Run()
}
