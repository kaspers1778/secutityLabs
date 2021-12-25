package internal

import (
	"L567/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func Encrypt(data string) []byte {
	key,err:=ioutil.ReadFile(config.KEY_FILE_PATH)
	if err!=nil{
		return nil
	}
	c,err:=aes.NewCipher(key)
	if err!=nil{
		return nil
	}
	gcm,err:=cipher.NewGCM(c)
	if err!=nil{
		return nil
	}
	nonce:=make([]byte,gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	encryptedData:=gcm.Seal(nonce,nonce,[]byte(data),nil)
	return encryptedData
}

func Decrypt(encryptedData []byte) string{
	key,err:=ioutil.ReadFile(config.KEY_FILE_PATH)
	if err!=nil{
		return ""
	}
	c,err:=aes.NewCipher(key)
	if err!=nil{
		return ""
	}
	gcm,err:=cipher.NewGCM(c)
	if err!=nil{
		return ""
	}
	nonceSize:=gcm.NonceSize()
	nonce,enctyptedData:=encryptedData[:nonceSize],encryptedData[nonceSize:]
	data,err:=gcm.Open(nil,nonce,enctyptedData,nil)
	if err!=nil{
		return ""
	}
	return string(data)
}