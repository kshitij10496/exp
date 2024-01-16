FROM golang:latest

WORKDIR /app

COPY . ./

RUN go build -o svc-weather

CMD ["./svc-weather"]
