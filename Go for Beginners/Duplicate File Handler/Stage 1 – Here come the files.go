package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Directory is not specified")
		return
	}
	dir := os.Args[1]

	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Println(path)
		}
		return err
	})
}
