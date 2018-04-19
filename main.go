package main

import (
	"fmt"
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("./wordlist.txt")
	checkError(err)

	lines := strings.Split(string(data), "\n")
	minimumLength := 64

	passphrase := []string{}
	passphraseText := ""

	for len(passphraseText) < minimumLength {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(lines))))
		checkError(err)

		word := lines[index.Int64()]
		if !strings.Contains(word, "'") {
			passphrase = append(passphrase, word)
			passphraseText += word
		}
	}

	fmt.Println(strings.Join(passphrase, " "))
}
