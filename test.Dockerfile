FROM golang:1.17

WORKDIR /go/src/app
COPY . /go/src/app
COPY ./wait-for-it.sh /wait-for-it.sh
RUN chmod 555 /wait-for-it.sh

CMD /wait-for-it.sh db:5432 -t 30; go test -v -coverprofile=./reports/coverage.txt -covermode=atomic -coverpkg=./... ./...

