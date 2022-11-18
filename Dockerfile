FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o api .

FROM alpine:3.16

COPY --from=builder /app/api /bin/

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ARG API_DATABASE_MIGRATIONS_PATH=/migrations
ENV API_DATABASE_MIGRATIONS_PATH=$API_DATABASE_MIGRATIONS_PATH

COPY --from=builder /app/database/migrations/ $API_DATABASE_MIGRATIONS_PATH

EXPOSE 8080

CMD ["api serve"]