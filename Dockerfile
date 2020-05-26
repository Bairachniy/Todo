FROM golang:latest as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o todoserv -v ./server.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/todoserv /app/todoserv

CMD ./todoserv