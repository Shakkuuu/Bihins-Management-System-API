FROM golang:1.19.3

RUN rm -rf /api
RUN mkdir /api

COPY go.mod /api/go.mod
COPY go.sum /api/go.sum

WORKDIR /api
RUN ls
RUN go mod download

COPY . /api

EXPOSE 8080

CMD go run main.go
