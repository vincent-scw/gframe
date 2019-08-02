FROM golang:latest
LABEL maintainer="Vincent Shen <edentidus@foxmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o receptionservice ./reception_svc/main.go

EXPOSE 80 443
CMD ["./receptionservice"]