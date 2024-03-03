FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o queer-hunt-backend .

EXPOSE 8080

CMD ["./queer-hunt-backend"]