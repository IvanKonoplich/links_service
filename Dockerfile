FROM golang:alpine AS builder

WORKDIR /

ADD go.mod .

COPY . .

RUN go build -o links cmd/main.go

FROM alpine

WORKDIR /

COPY --from=builder config/config.toml config/

COPY --from=builder links .

EXPOSE 8000

CMD ["./links"]