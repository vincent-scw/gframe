FROM golang:latest
LABEL maintainer="Vincent Shen <edentidus@foxmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./admin_svc .
COPY ./events .
COPY ./kafkactl .
COPY ./redisctl .
RUN go build -o admin_svc ./admin_svc/main.go

EXPOSE 80 443
CMD ["./admin_svc"]