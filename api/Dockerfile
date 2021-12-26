FROM cosmtrek/air as dev
WORKDIR /opt/app/api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD air

FROM golang:1.16
WORKDIR /opt/app/api
COPY go.mod go.sum ./
RUN go mod download
CMD go build .