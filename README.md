# GOPASSWD

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