FROM golang:1.13-alpine AS builder

RUN apk add --no-cache git gcc g++ make

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

RUN go build -o /app/main ./main.go

FROM alpine:latest

WORKDIR /app

RUN apk update && \
    apk add ca-certificates && \
    apk add postgresql-client

COPY --from=builder /app/main /usr/local/bin/restful

RUN chmod +x /usr/local/bin/restful

CMD ["/usr/local/bin/restful"]