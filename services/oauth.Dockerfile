FROM golang:latest
LABEL maintainer="Vincent Shen <edentidus@foxmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o oauthservice ./oauth

EXPOSE 80 443
CMD ["./oauthservice"]