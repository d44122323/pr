FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY back ./back
COPY front ./front

WORKDIR /app/back

RUN go build -o server

EXPOSE 8080

CMD ["./server"]
