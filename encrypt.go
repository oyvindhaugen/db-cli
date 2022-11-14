package main

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//both bytes and encStr should be in .env file
var bytes = []byte{87, 97, 92, 64, 6, 25, 73, 3, 42, 22, 93, 35, 14, 29, 11, 53}

const encStr string = "h4so9(?^l[f^27jium^="

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Encrypt(text, encStr string) (string, error) {
	block, err := aes.NewCipher([]byte(encStr))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func EncryptString(s string) string {

	encText, err := Encrypt(s, encStr)
	if err != nil {
		log.Fatal("Error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	return encText
}

// func Decode(s string) []byte {
// 	data, err := base64.StdEncoding.DecodeString(s)
// 	if err != nil {
// 		log.Fatal("Could not decode data: ", err)
// 	}
// 	return data
// }
// func Decrypt(text, encStr string) (string, error) {
// 	block, err := aes.NewCipher([]byte(encStr))
// 	if err != nil {
// 		log.Fatal("Could not decrypt: ", err)
// 	}
// 	cipherText := Decode(text)
// 	cfb := cipher.NewCFBDecrypter(block, bytes)
// 	plainText := make([]byte, len(cipherText))
// 	cfb.XORKeyStream(plainText, cipherText)
// 	return string(plainText), nil
// }
// func DecryptString(s string) {
// 	decText, err := Decrypt(s, encStr)
// 	if err != nil {
// 		log.Fatal("Error decrypting your encrypted text: ", err)
// 	}
// 	fmt.Println(decText)
// }
func newUser(user string, pass string) {
	pass = EncryptString(pass)
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Could not secure connection to database: ", err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user, pass)
	if err != nil {
		log.Fatal("Could not execute query: ", err)
	}
}
func login(user string, pass string) bool {
	pass = EncryptString(pass)
	var checkPass string
	psqlconn := fmt.Sprintf("host= localhost port = 5432 user = oyvind password = iktfag dbname = test_db sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Could not secure connection to database: ", err)
	}
	defer db.Close()
	row := db.QueryRow("SELECT password FROM users WHERE username = $1", user)
	switch err := row.Scan(&checkPass); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
	case nil:
		if checkPass != pass {
			return false
		} else {
			return true
		}
	default:
		log.Fatal("Error")
	}
	return false
}
