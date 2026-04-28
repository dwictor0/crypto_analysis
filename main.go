package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
	}
}
