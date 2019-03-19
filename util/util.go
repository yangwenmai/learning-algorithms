package util

import (
	"io/ioutil"
	"log"
	"os"
)

// ReadFile read file
func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}

// WriteFile write file
func WriteFile(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
