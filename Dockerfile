FROM golang as builder

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server .

FROM alpine:3.10.3
COPY --from=builder /app/server ./server
EXPOSE 8080

CMD ./server