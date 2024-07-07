FROM golang:1.22.5 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/main.go

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/src/infrastructure/secondary/repositories/files/seed/data/portfolios.clients-portfolios.json ./data/portfolios.clients-portfolios.json

ENV MONGO_DB_HOST=localhost \
    MONGO_DB_PORT=27017 \
    MONGO_DB_DATABASE=portfolios \
    REDIS_HOST=localhost \
    PORT=3000 \
    FILE_PATH=./data/portfolios.clients-portfolios.json

EXPOSE 3000

CMD ["./main"]