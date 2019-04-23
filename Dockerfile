FROM golang:latest

COPY . /gomysql/

WORKDIR /gomysql/

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on

RUN go build -a -i -o gomysql main.go

CMD ["./gomysql"]