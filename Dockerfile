FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/main /main
CMD ["/main"]
