# GOPASSWD

## RUN
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
export KEY="passphrasewhichneedstobe32bytes1"
go run main.go
```

## Docker
```
git clone https://github.com/rombintu/gopasswd.git
cd gopasswd
sudo docker build -t gopasswd/docker .
sudo docker run -d -p 5000:8080 gopasswd/docker
```