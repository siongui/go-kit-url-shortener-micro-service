# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.17-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get -u github.com/gopherjs/gopherjs

COPY *.go ./
COPY ./frontend/ ./frontend/
COPY ./datasource/ ./datasource/

RUN gopherjs build frontend/*.go -m -o frontend/app.js
RUN go build -o /url-shortener

EXPOSE 8080

CMD [ "/url-shortener" ]
