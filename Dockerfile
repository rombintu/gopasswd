FROM alpine:latest
RUN apk add --no-cache go git
RUN git clone https://github.com/rombintu/gopasswd.git
WORKDIR gopasswd
RUN go build -o main main.go
ENV KEY=passphrasewhichneedstobe32bytes1
ENV CREDS="host=127.0.0.1 user=gopasswd password=gopasswd dbname=gopasswd port=5432 sslmode=disable"
CMD ["./main"]
