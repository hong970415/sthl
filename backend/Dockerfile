FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o sthl-app

FROM alpine

WORKDIR /app

COPY --from=builder /app/sthl-app sthl-app

CMD ["./sthl-app"]