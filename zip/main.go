package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	pathToDirectory = "./zip/test"
	pathToZip       = "./zip/test.zip"
)

func main() {
	//This will create a zip file a zip file
	if err := Zip(pathToDirectory, pathToZip); err != nil {
		panic(err)
	}

	fmt.Println("Zip file was created")
}

//Zip method for zip the selected src to specific destination
func Zip(src string, dest string) error {
	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	zipDestFile := zip.NewWriter(destinationFile)
	if err := filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filePath == dest {
			return nil
		}

		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(filePath, src)
		zipFile, err := zipDestFile.Create(relPath)

		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}

		if err := fsFile.Close(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if err := zipDestFile.Close(); err != nil {
		return err
	}

	return nil
}
