FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN mkdir -p output
RUN CGO_ENABLED=0 go build -tags netgo -ldflags '-extldflags "-static"' -o output/thanks ./cmd/main.go


# run
FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/log

COPY --from=builder /app/output /app

RUN chmod +x /app/thanks

CMD ./thanks
