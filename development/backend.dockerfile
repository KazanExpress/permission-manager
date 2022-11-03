FROM golang:1.18.8

RUN go get -u github.com/cosmtrek/air

WORKDIR /app

COPY ./go.mod go.sum ./

RUN go mod download

ENTRYPOINT air -c ./development/air.toml

