FROM golang:1.14.1-alpine3.11 as builder

RUN apk add make

WORKDIR /src
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
RUN go mod download 
COPY . .
RUN make build

FROM alpine:3.11.6
WORKDIR /app
COPY --from=builder /src/bin/main /app/main
EXPOSE 8080

CMD /app/main