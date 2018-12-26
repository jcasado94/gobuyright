FROM golang:latest

WORKDIR /go/src/github.com/jcasado94/gobuyright
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gobuyright"]
