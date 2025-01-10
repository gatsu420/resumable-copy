FROM golang:1.22.4-alpine AS builder
RUN apk add --no-cache git gcc musl-dev
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o resumable-copy

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /src/resumable-copy .
COPY source.txt .
CMD ["sh", "-c", \
    "./resumable-copy copy \
    --src source.txt \
    --dest destination.txt \
    --resume-at 5 \
    --chunk-size 4 \
    --lag 3"]
