package validate

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

type WordlistHash struct{}

func IdentifyHash(password string, hashList string) {
	init := WordlistHash{}
	init.Validate(password, hashList)
}

func (u *WordlistHash) Validate(password string, hashList string) string {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if hash == hashList {
		fmt.Printf("[+] Hash encontrado (MD5): %s\n", password)
	}
	hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	if hash == hashList {
		fmt.Printf("[+] Hash encontrado (SHA-256): %s\n", password)
	}
	return "Nao identificado"
}
