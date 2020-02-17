package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var dir string

func main() {

	dir = getCurrentDir()
	fmt.Println(fmt.Sprintf("Default directory path is: %s", dir))

	//We switch pointer to the root directory for control the path from which we need to generate test-data file-paths
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(fmt.Sprintf("The package path is: %s", path.Dir(filename)))
	dir = path.Join(path.Dir(filename))

	if err := os.Chdir(dir); err != nil {
		panic(err)
	}

	fmt.Println("The path was switched")

	dir = getCurrentDir()
	fmt.Println(fmt.Sprintf("This is the current directory: %s", dir))
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}
