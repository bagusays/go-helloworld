FROM golang:1.18.3-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter,netgo,nomsgpack -ldflags='-s -w -extldflags "-static"' -o engine


FROM scratch
COPY --from=builder /app/engine /