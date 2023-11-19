FROM golang:1.18

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /go/bin/app ./cmd

EXPOSE 50050

CMD ["/go/bin/app"]