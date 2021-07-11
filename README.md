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
cp .env.bak .env
```
Edit .env (KEY - 32 bytes, SECRET)
```
go run main.go
```

## Build
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
cp .env.bak .env
```
Edit .env (KEY - 32 bytes, SECRET)
```
go build -o gopasswd main.go
./gopasswd
```

## Docker build
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
sudo docker build -t gopasswd/alpine .
sudo docker run -d -p 5000:8080 gopasswd/alpine
```

### Screenshots

![alt text](/screenshots/login.png)
![alt text](/screenshots/index.png)
![alt text](/screenshots/create.png)
<img src="/screenshots/mob_login.png" alt="drawing" width="45%"/> <img src="/screenshots/mob_create.png" alt="drawing" width="45%"/>