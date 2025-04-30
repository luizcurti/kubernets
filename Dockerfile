FROM golang:1.15
WORKDIR /appdocker build -t luizcurti/hello-go .
COPY . .
RUN go build -o server .
EXPOSE 8000
CMD ["./server"]