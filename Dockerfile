FROM alpine:latest
RUN apk add --no-cache go git
RUN git clone https://github.com/rombintu/gopasswd.git
WORKDIR gopasswd
ENV KEY="passphrasewhichneedstobe32bytes1"
ENV SECRET="secret"
RUN go build -o main main.go
CMD ["./main"]