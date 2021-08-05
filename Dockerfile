FROM alpine:latest
RUN apk add --no-cache go git
RUN git clone https://github.com/rombintu/gopasswd.git
WORKDIR gopasswd
RUN go build -o main main.go
ENV KEY=passphrasewhichneedstobe32bytes1
CMD ["./main"]