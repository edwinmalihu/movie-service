FROM golang:1.19.1-alpine3.15 AS build_base

WORKDIR /tmp/api

RUN apk add git
COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.9 

WORKDIR /app

RUN chgrp -R 0 /app && \
    chmod -R g=u /app

COPY --from=build_base /tmp/api/main api

CMD ["./api"]
