FROM golang:1.25-alpine3.22 as builder

WORKDIR /app
COPY go.mod .
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]