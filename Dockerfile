FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod init programmingpercy/hellogopher
RUN go get github.com/go-sql-driver/mysql
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hellogopher -ldflags="-w -s"

ENTRYPOINT [ "./hellogopher" ]