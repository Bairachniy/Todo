FROM golang:1.14.2-onbuild

RUN /
WORKDIR /Todorun

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o todoserv


FROM alpine:3.10.3

WORKDIR /askmechat

COPY --from=builder /app/cmd/askmechatprocessor /askmechat/askmechatprocessor

EXPOSE 8090

CMD ./askmechatprocessor