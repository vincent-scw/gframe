FROM golang:latest
LABEL maintainer="Vincent Shen <edentidus@foxmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o brokerservice ./broker_svc

EXPOSE 80 443
CMD ["./brokerservice"]