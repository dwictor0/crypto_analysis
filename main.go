package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("hash.txt")
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
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == hashList {
			fmt.Printf("[+] Hash encontrado (MD5): %s\n", password)
		}
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if hash == hashList {
			fmt.Printf("[+] Hash encontrado (SHA-256): %s\n", password)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
