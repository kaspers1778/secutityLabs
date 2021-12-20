package main

import (
	"Cripto_L3/Lab4/internal"
	"bufio"
	"log"
	"os"
)

func main() {
	pg:=internal.PasswordsGenerator{
		Top100:   10,
		Top100K:  50,
		Modified: 35,
		Random:   5,
	}
	passwords:=pg.GeneratePasswords(100000);

	hs:=internal.Hasher{Passwords: passwords}
	mp5Hash:=hs.HashMD5()
	sha1Hash:=hs.HashBCrypt()
	writeToFile("Lab4/output/hash1.txt",mp5Hash)
	writeToFile("Lab4/output/hash2.txt",sha1Hash)

}

func writeToFile(fileName string,data []string){
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range data {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
