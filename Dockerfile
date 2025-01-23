FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/AltTube-Go .

# deployment image
FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/AltTube-Go ./

CMD [ "./AltTube-Go" ]

EXPOSE 8072
