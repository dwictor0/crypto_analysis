package main

import (
	"bufio"
	"cryptoanalysis/validate"
	"log"
	"os"
)

func main() {
	file, err := os.Open("wordlists/hash.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		hashes := scanner.Text()
		makeHash(hashes)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func makeHash(hashList string) {
	file, err := os.Open("wordlists/wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		validate.IdentifyHash(password, hashList)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
