package internal

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type PasswordsGenerator struct {
	Top100 int
	Top100K int
	Modified int
	Random int
}

func (pg *PasswordsGenerator) GeneratePasswords(amount int) []string{
	passwords:=append(pg.GenerateTop100Passwords(amount*pg.Top100 / 100),pg.GenerateTop100KPasswords(amount*pg.Top100K/100)...)
	passwords=append(passwords,pg.GenerateModifiedPasswords(amount*pg.Modified/100)...)
	passwords=append(passwords,pg.GenerateRandomPasswords(amount*pg.Random/100,10)...)
	passwords=shuffle(passwords)
	return passwords

}

func (pg *PasswordsGenerator) GenerateTop100Passwords(amount int) []string{
	top100passwords:=readFile("Lab4/input/top100.txt")
	for amount> len(top100passwords){
		top100passwords = append(top100passwords, top100passwords...)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(top100passwords), func(i, j int) { top100passwords[i], top100passwords[j] = top100passwords[j], top100passwords[i] })
	return top100passwords[:amount]
}

func (pg *PasswordsGenerator) GenerateTop100KPasswords(amount int) []string{
	top100Kpasswords :=readFile("Lab4/input/top100K.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(top100Kpasswords), func(i, j int) { top100Kpasswords[i], top100Kpasswords[j] = top100Kpasswords[j], top100Kpasswords[i] })
	return top100Kpasswords[:amount]
}

func (pg *PasswordsGenerator) GenerateModifiedPasswords(amount int) []string{
	ModdifiedPasswords :=readFile("Lab4/input/top100K.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ModdifiedPasswords), func(i, j int) { ModdifiedPasswords[i], ModdifiedPasswords[j] = ModdifiedPasswords[j], ModdifiedPasswords[i] })
	for i,password := range(ModdifiedPasswords[:amount]){
		password=strings.ReplaceAll(password,"a","@")
		password=strings.ReplaceAll(password,"o","0")
		password=strings.ReplaceAll(password,"l","1")
		password=strings.ReplaceAll(password,"w","vv")
		password=strings.ReplaceAll(password,"b","8")
		password=strings.ReplaceAll(password,"e","3")
		password=strings.ReplaceAll(password,"m","M")
		password=strings.ReplaceAll(password,"5","V")
		ModdifiedPasswords[i] = password
	}
	return ModdifiedPasswords[:amount]
}

func (pg *PasswordsGenerator) GenerateRandomPasswords(amount int,len int) []string{
	symbols:=strings.Split("0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz","")
	randomPasswords:=make([]string,0)
	for i:=0;i<amount;i++{
		p:=strings.Join(shuffle(symbols)[:len],"")
		randomPasswords = append(randomPasswords,p)
	}
	return randomPasswords
}

func shuffle(slice []string) []string{
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}


func readFile(path string) []string{
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}
