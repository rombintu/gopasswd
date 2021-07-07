FROM golang:latest
RUN git clone https://github.com/rombintu/gopasswd.git
WORKDIR gopasswd
ENV KEY="passphrasewhichneedstobe32bytes1"
CMD ["go", "run", "main.go"]