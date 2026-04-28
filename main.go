package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

var md5hash = ""
var sha256hash = ""

func main() {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		password := scanner.Text()
		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == md5hash {
			fmt.Printf("[+] Hash encontrado (MD5): %s\n", password)
		}
	}
}
