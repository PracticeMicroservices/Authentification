FROM golang:1.18-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build  -o authentication-service ./cmd/api
RUN chmod +x /app/authentication-service

#build a tiny docker image
FROM  alpine:latest


RUN mkdir /app

COPY --from=builder /app/authentication-service /app

CMD ["/app/authentication-service"]