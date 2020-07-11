FROM golang:1.14-alpine
WORKDIR /app
ADD . /app
WORKDIR /app/main_api
RUN go build -o main .
CMD ["/app/main_api/main"]