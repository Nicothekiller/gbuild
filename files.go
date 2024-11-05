package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// get cpp source files to compile, returns a []string
func getFiles() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return find(path, ".cpp")
}

// func to get files with an extension in a dir
func find(root, ext string) (res []string) {
	filepath.WalkDir(root, func(fileName string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(file.Name()) == ext {
			res = append(res, fileName)
		}
		return nil
	})
	return
}
