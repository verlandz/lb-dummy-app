# cmd: sudo docker build --tag verlandz/lb-dummy-app:1.0 .
FROM golang:1.11
EXPOSE 8080

WORKDIR $GOPATH
COPY . $PWD

CMD ["sh","-c","go run $PWD/main.go"]