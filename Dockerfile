FROM golang:1.17

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /server

EXPOSE 5000

ENTRYPOINT [ "/server" ]
