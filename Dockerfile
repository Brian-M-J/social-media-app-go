FROM golang:1.23 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o socio cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/socio /app/socio

EXPOSE 3015

CMD ["./socio"]
