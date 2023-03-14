FROM golang:1.20-alpine

WORKDIR /app
COPY . ./
RUN go mod init github.com/kkti4216/app
RUN go mod tidy
RUN go build -o /main

RUN chmod 755 script.sh
RUN mkdir -p /var/log/app

EXPOSE 3000

CMD ["/main"]
