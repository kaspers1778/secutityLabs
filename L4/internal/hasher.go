package internal

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sync"
)

type Hasher struct {
	Passwords []string
}

func (h *Hasher) HashMD5() []string {
	hashedPasswords:=make([]string,0)
	for _, password := range h.Passwords {
		hashedPasswords = append(hashedPasswords, fmt.Sprintf("%x",md5.Sum([]byte(password))))
	}
	return hashedPasswords
}

func putPasswords(input *chan string,passwords []string ){
	for _,p:=range passwords{
		*input <- p
	}
	close(*input)
}

func getHashes(output *chan string,hashes *[]string){
	for h:=range *output{
		*hashes = append(*hashes,h)
	}
}

func hash(input *chan string,output *chan string,group *sync.WaitGroup) {
	for password := range *input {
		p,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err!=nil{
			log.Print(err)
		}
		*output<-fmt.Sprintf("%x",p)
	}
	group.Done()
}

func (h *Hasher) HashBCrypt() []string {
	hashedPasswords:=make([]string,0)
	passwords:=make(chan string,1)
	hashes:=make(chan string,1)
	go putPasswords(&passwords,h.Passwords)
	var wg sync.WaitGroup
	for j:=0;j<10;j++{
		wg.Add(1)
		go hash(&passwords,&hashes,&wg)
	}
	go getHashes(&hashes,&hashedPasswords)
	wg.Wait()
	//println(len(hashedPasswords))
	close(hashes)
	return hashedPasswords
}
