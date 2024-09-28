FROM alpine:latest

RUN apk add --no-cache go

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
