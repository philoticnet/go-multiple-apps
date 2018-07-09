package utils

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(name string, text string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, text)
}
