# database image
FROM mysql:latest as todoapidb
ARG MYSQL_DATABASE
ARG MYSQL_ROOT_PASSWORD
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ENV MYSQL_DATABASE $MYSQL_DATABASE
ENV MYSQL_ROOT_PASSWORD $MYSQL_ROOT_PASSWORD
ENV MYSQL_USER $MYSQL_USER
ENV MYSQL_PASSWORD $MYSQL_PASSWORD

# api image
FROM golang:1.16.2-buster as todoapi
WORKDIR /go/src/app
COPY ../.. .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["./util/wait-for-it.sh", "database:3306", "--", "todoapi"]