package main

import (
	"fmt"
	"path"
)

func main() {

	dir, file := path.Split("css/main.css")
	fmt.Println("file:" , file)
	fmt.Println("dir:", dir)
}