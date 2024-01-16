FROM golang:latest

WORKDIR /app

COPY . ./

RUN go build -o svc-calculator

CMD ["./svc-calculator"]
