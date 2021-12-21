package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var(
	TASK1="7958401743454e1756174552475256435e59501a5c524e176f786517545e475f5245191772195019175e4317445f58425b531743565c521756174443455e595017d5b7ab5f525b5b58174058455b53d5b7aa175659531b17505e41525917435f52175c524e175e4417d5b7ab5c524ed5b7aa1b174f584517435f5217515e454443175b524343524517d5b7ab5fd5b7aa17405e435f17d5b7ab5cd5b7aa1b17435f5259174f584517d5b7ab52d5b7aa17405e435f17d5b7ab52d5b7aa1b17435f525917d5b7ab5bd5b7aa17405e435f17d5b7ab4ed5b7aa1b1756595317435f5259174f58451759524f4317545f564517d5b7ab5bd5b7aa17405e435f17d5b7ab5cd5b7aa175650565e591b17435f525917d5b7ab58d5b7aa17405e435f17d5b7ab52d5b7aa1756595317445817585919176e5842175a564e17424452175659175e5953524f1758511754585e59545e53525954521b177f565a5a5e595017535e4443565954521b177c56445e445c5e17524f565a5e5956435e58591b17444356435e44435e54565b17435244434417584517405f564352415245175a52435f5853174e5842175152525b174058425b5317445f584017435f52175552444317455244425b4319"
	TASK2="G0IFOFVMLRAPI1QJbEQDbFEYOFEPJxAfI10JbEMFIUAAKRAfOVIfOFkYOUQFI15ML1kcJFUeYhA4IxAeKVQZL1VMOFgJbFMDIUAAKUgFOElMI1ZMOFgFPxADIlVMO1VMO1kAIBAZP1VMI14ANRAZPEAJPlMNP1VMIFUYOFUePxxMP19MOFgJbFsJNUMcLVMJbFkfbF8CIElMfgZNbGQDbFcJOBAYJFkfbF8CKRAeJVcEOBANOUQDIVEYJVMNIFwVbEkDORAbJVwAbEAeI1INLlwVbF4JKVRMOF9MOUMJbEMDIVVMP18eOBADKhALKV4JOFkPbFEAK18eJUQEIRBEO1gFL1hMO18eJ1UIbEQEKRAOKUMYbFwNP0RMNVUNPhlAbEMFIUUALUQJKBANIl4JLVwFIldMI0JMK0INKFkJIkRMKFUfL1UCOB5MH1UeJV8ZP1wVYBAbPlkYKRAFOBAeJVcEOBACI0dAbEkDORAbJVwAbF4JKVRMJURMOF9MKFUPJUAEKUJMOFgJbF4JNERMI14JbFEfbEcJIFxCbHIJLUJMJV5MIVkCKBxMOFgJPlVLPxACIxAfPFEPKUNCbDoEOEQcPwpDY1QDL0NCK18DK1wJYlMDIR8II1MZIVUCOB8IYwEkFQcoIB1ZJUQ1CAMvE1cHOVUuOkYuCkA4eHMJL3c8JWJffHIfDWIAGEA9Y1UIJURTOUMccUMELUIFIlc="
	//TASK3="EFFPQLEKVTVPCPYFLMVHQLUEWCNVWFYGHYTCETHQEKLPVMSAKSPVPAPVYWMVHQLUSPQLYWLASLFVWPQLMVHQLUPLRPSQLULQESPBLWPCSVRVWFLHLWFLWPUEWFYOTCMQYSLWOYWYETHQEKLPVMSAKSPVPAPVYWHEPPLUWSGYULEMQTLPPLUGUYOLWDTVSQETHQEKLPVPVSMTLEUPQEPCYAMEWWYTYWDLUULTCYWPQLSEOLSVOHTLUYAPVWLYGDALSSVWDPQLNLCKCLRQEASPVILSLEUMQBQVMQCYAHUYKEKTCASLFPYFLMVHQLUPQLHULIVYASHEUEDUEHQBVTTPQLVWFLRYGMYVWMVFLWMLSPVTTBYUNESESADDLSPVYWCYAMEWPUCPYFVIVFLPQLOLSSEDLVWHEUPSKCPQLWAOKLUYGMQEUEMPLUSVWENLCEWFEHHTCGULXALWMCEWETCSVSPYLEMQYGPQLOMEWCYAGVWFEBECPYASLQVDQLUYUFLUGULXALWMCSPEPVSPVMSBVPQPQVSPCHLYGMVHQLUPQLWLRPOEDVMETBYUFBVTTPENLPYPQLWLRPTEKLWZYCKVPTCSTESQPBYMEHVPETCMEHVPETZMEHVPETKTMEHVPETCMEHVPETT"
)

func main(){
	//getTasks()
	//task1()
	//showCoincedences(TASK2)
	//task2()
}

func task2(){
	task2:=TASK2
	text, _ := base64.StdEncoding.DecodeString(task2)
	println(len(text))

	for a:=0;a<256;a++{
		for b:=0;b<256;b++{
			for c:=0;c<256;c++{
				w:=string(a)+string(b)+string(c)
				newText:=repeatedXOR(text,w)
				mp:=countBytes([]byte(newText))
				if mp["e"]>50 && mp["t"]>30{
					fmt.Printf("%v - %v\n",w,newText)
				}
			}
		}
	}
}

func countBytes(text []byte)map[string]int{
	m:=make(map[string]int)
	for _,b:=range text{
		m[string(b)]+=1
	}
	return m
}

func repeatedXOR(text []byte,key string) string{
	j:=0
	for i,b:=range text{
		if j == len(key){
			j = 0
		}
		text[i] = b ^ key[j]
		j++
	}
	return string(text)
}

func showCoincedences(str string){
	buf := str
	for i:=0;i<len(str);i++{
		buf=moveString(buf)
		fmt.Printf("%v - %v\n",i,countCoincedences(str,buf))
	}
}

func countCoincedences(a,b string) int{
	if len(a)!=len(b){
		return -1
	}
	same:=0
	for i,l:=range a{
		if l == rune(b[i]){
			same++
		}
	}
	return same
}

func moveString(str string) string{
	return str[len(str)-1:]+str[:len(str)-1]
}


func task1(){
	task1,_:=hex.DecodeString(TASK1)
	symbols:=[]rune("0123456789@ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	for _,s:=range symbols{
		fmt.Printf("%v - %v\n",string(s),xorSlice(task1,byte(s)))
	}
}

func xorSlice(slice []byte,key byte) string{
	for i,b:=range slice{
		slice[i] = (b ^ key)
	}
	return string(slice)
}

func getTasks(){
	base64,err:=ioutil.ReadFile("Lab1/base64.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	tasks:=base64ToString(string(base64))
	println(tasks)
}

func base64ToString(data string)string{
	sDec, _ := base64.StdEncoding.DecodeString(data)
	return string(sDec)
}

func readFile(path string) string{
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	txtline:=scanner.Text()

	file.Close()

	return txtline
}

func readFileLines(path string) []string{
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
