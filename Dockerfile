FROM golang:1.18.3-alpine

WORKDIR /app

COPY go.mod ./

COPY . .
COPY web/ /
COPY .secret /

RUN go build -o /mock-server ./cmd/mock-server

EXPOSE 8000

CMD ["/mock-server"]
