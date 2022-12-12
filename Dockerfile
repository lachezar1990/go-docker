# syntax=docker/dockerfile:1

FROM golang:alpine

ADD . /app

WORKDIR /app

RUN go mod download
RUN go install

RUN go build -o /docker-gs-ping

EXPOSE 5555

CMD [ "/docker-gs-ping" ]