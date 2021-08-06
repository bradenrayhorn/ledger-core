FROM golang:1.16.4

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -v -t -d ./...
RUN go build -v .

CMD go test -v -coverprofile=./reports/coverage.txt -covermode=atomic -coverpkg=./... ./...

