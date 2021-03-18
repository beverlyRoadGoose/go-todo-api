# database image
FROM mysql:latest as todoapidb

# api image
FROM golang:1.16.2-buster as todoapi
WORKDIR /go/src/app
COPY ../.. .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["./util/wait-for-it.sh", "database:3306", "--", "todoapi"]