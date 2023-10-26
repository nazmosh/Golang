package main

import (
	"fmt"
	"path"
)

var width float32
var height float32

func main() {

	dir, file := path.Split("css/main.css")
	fmt.Println("file:" , file)
	fmt.Println("dir:", dir)
}