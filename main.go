package main

import (
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var password string

	fmt.Print("Digite a senha desejada: ")
	fmt.Scanln(&password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar hash:", err)
		return
	}

	fmt.Println("Hash da senha encriptada com bcrypt:")
	fmt.Println(string(hash))

	sha512Hash := sha512.New()
	sha512Hash.Write(hash)
	sha512HashBytes := sha512Hash.Sum(nil)

	fmt.Println("Hash da senha encriptada com bcrypt convertido para SHA-512:")
	fmt.Printf("%x\n", sha512HashBytes)
}
