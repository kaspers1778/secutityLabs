### LAB 5.
In order to make this task was created server with standart Golang package 
**"net/http"** and front with another standart package **"html/template"**.<p>
To hash password were used **Bcrypt** algorithm, from **"golang.org/x/crypto/bcrypt"**. Because it's pretty strong and self-salting. Hashed passwords stores in the **postgress** db.<p>
Password must respond next rules:
1. Must contain 7 or more letters
2. Must contain at least 1 number
3. Must contain at least 1 special character
4. Must contain at least 1 letter in Upper register


### LAB 6.
User have sensitive information: **phone number** and **favorite color**, which he enter on sign up and can see it in his personal cabinet after he logged in successfully.<p>
After server got sensetive information on sign up it's ecrypt it with **AES** algorithm with **128bit key** in **GCM** mode using package **"crypto/aes"**. And after this stored in db. However key is stored in file system to prevent leak of encypted data and key to it in one time.<p>
We can lose our data in case that both db and key from files system were stolen. 

### LAB 7.
In order to make our server use TLS were created self-signed certificate and the private key, which stores in the file system.<p>
To create TLS connection we used golang package **"crypto/tls"**. As long as the newest version of protocol is TLS1.3 we chose to use it.<p>
And so it has acceptable only 5 ciphersuites but our package works only with 3 from it we use next ciphersuites:
1. TLS_AES_128_GCM_SHA256
1. TLS_AES_256_GCM_SHA384
1. TLS_CHACHA20_POL
