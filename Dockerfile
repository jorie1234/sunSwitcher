#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go/src/app/app /app
ENTRYPOINT ./app
