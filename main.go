package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

type HashDisponiveis struct {
	MD5Hash    string
	SHA256Hash string
}

var md5hash = "81dc9bdb52d04dc20036dbd8313ed055"
var sha256hash = "914420a9b210195dea7e8a1fdc5234fb1f413c04dba3b5eaabed9df6adb47f51"

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
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if hash == sha256hash {
			fmt.Printf("[+] Hash encontrado (SHA-256): %s\n", password)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
