package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Encrypt(fileName string, Password string) {

	File, err := ioutil.ReadFile(fileName)
	Data := []byte("(" + fileName + ")" + string(File))
	PasswordLength := len(Password)
	SplitText := "x"
	PasswordLength = 32 - PasswordLength
	for i := 1; i <= PasswordLength; i++ {
		Password += SplitText
	}
	key := []byte(Password)

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile(`(?m)(.*)\.`)
	RegexMatch := re.FindStringSubmatch(fileName)[1]

	err = ioutil.WriteFile(RegexMatch+".ash", gcm.Seal(nonce, nonce, Data, nil), 0777)
	if err != nil {

		fmt.Println(err)
	}
	os.Remove(fileName)
	fmt.Println("File Encrypted as " + RegexMatch + ".ash !")
}

func Decrypt(filename string, Password string) {

	PasswordLength := len(Password)
	SplitText := "x"
	PasswordLength = 32 - PasswordLength
	for i := 1; i <= PasswordLength; i++ {
		Password += SplitText
	}
	key := []byte(Password)
	ciphertext, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	DecryptedBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	decryptedString := string(DecryptedBytes)
	re := regexp.MustCompile(`(?m)\(([^)]+)\)`)
	RegexMatch := re.FindStringSubmatch(decryptedString)[1]
	dataString := strings.Replace(decryptedString, "("+RegexMatch+")", "", 1)
	dataBytes := []byte(dataString)
	err = ioutil.WriteFile(RegexMatch, dataBytes, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File Decrypted Successfully !")
}
func main() {
	// Shima -e -f test.png -p 1234
	// Shima -d -f test.png -p 1234

	encrypt := flag.Bool("e", false, "encrypt")
	decrypt := flag.Bool("d", false, "decrypt")
	fileName := flag.String("f", "", "fileName")
	password := flag.String("p", "", "password")
	flag.Parse()

	if *encrypt && *decrypt || !*encrypt && !*decrypt {
		fmt.Println("You Have to use one of -e or -d to encrypt or decrypt")
		os.Exit(1)
	}

	if len(*password) <= 0 {
		fmt.Println("Password Can't Be Empty , Please Use -p to enter Password")
		os.Exit(1)
	}

	if len(*password) >= 32 {
		fmt.Println("Password Can't Be Empty , Please Use -p to enter Password")
		os.Exit(1)
	}

	if !fileExists(*fileName) {
		fmt.Println("File doesn't exists")
		os.Exit(1)
	}

	if *encrypt {
		Encrypt(*fileName, *password)
	}

	if *decrypt {
		Decrypt(*fileName, *password)
	}

}
