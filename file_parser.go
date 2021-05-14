package main

import (
	"log"
	"os"
)

func GetFile(dir string) ([]string) {
	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("[ERROR] - %s", err)
	}

	files, err := file.Readdirnames(-1)
	if err != nil {
		log.Fatalf("[ERROR] - %s", err)
	}

	return files
}

func ParseFile(fileName string) string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("[ERROR] - %s", err)
	}

	defer file.Close()
	return file.Name()
}

func main() {
	_dir := "."
	_fileName := "microservice-patterns-in-2021.md"
	GetFile(_dir)
	ParseFile(_fileName)
}
