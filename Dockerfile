FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

EXPOSE 6565

CMD ["/app/main"]