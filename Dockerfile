FROM alpine:latest
RUN apk add --no-cache go git
RUN git clone https://github.com/rombintu/gopasswd.git
WORKDIR gopasswd
RUN cp .env.bak .env
ENV KEY="passphrasewhichneedstobe32bytes4"
ENV SECRET="secret"
RUN go build -o main main.go
CMD ["./main"]