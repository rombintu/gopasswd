# GOPASSWD

## Description
**Gopasswd** is a simple password manager  
Run **Gopasswd** on your PC or Server  
Store encrypted passwords on your server  
Import passwords from evil browsers and cloud services  

## RUN
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
export CREDS="host=127.0.0.1 user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
export KEY=passphrasewhichneedstobe32bytes0
go run main.go
```

## Build
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
go build -o gopasswd main.go
export CREDS="host=127.0.0.1 user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
export KEY=passphrasewhichneedstobe32bytes0
./gopasswd
```

## Docker build
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
sudo docker build -t gopasswd/alpine .
sudo docker run -d -p 80:8080 gopasswd/alpine
```

## Docker-compose 
```
sudo docker-compose up
```

### Screenshots

![alt text](/screenshots/login.png)
![alt text](/screenshots/index.png)
![alt text](/screenshots/create.png)
<img src="/screenshots/mob_login.png" alt="drawing" width="45%"/> <img src="/screenshots/mob_create.png" alt="drawing" width="45%"/>