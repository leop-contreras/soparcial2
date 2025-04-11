FROM golang:1.23.5-alpine

WORKDIR /back

RUN apk add --no-cache git

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "main.go"]
