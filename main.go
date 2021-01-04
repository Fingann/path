package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func fixPrefix(filename string) string {
	if filename[0] == '.' && filename[1] == '\\' {
		return string(os.PathSeparator) + filename[2:]
	} else if filename[0] == os.PathSeparator {
		return filename
	}
	return string(os.PathSeparator) + filename
}

func main() {
	file := strings.TrimSpace(os.Args[1])
	file = fixPrefix(file)
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(path + file)

}
