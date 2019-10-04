FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 80
CMD ["./main"]
