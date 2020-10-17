
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN apk add build-base
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/app
RUN ls -lah
RUN pwd

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENV GIN_MODE=release
ENTRYPOINT ./app
LABEL Name=api Version=0.0.1
EXPOSE 3000
