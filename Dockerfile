FROM golang:1.23-alpine as builder

LABEL maintainer="tapiaw38 Singh <tapiaw38@gmail.com>"

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

RUN apk add --no-cache build-base

COPY . /app/

EXPOSE 8083

COPY entrypoint.sh .

COPY .env .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]