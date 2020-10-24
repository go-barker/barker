#build stage (Go)
FROM golang:alpine AS backBuilder
WORKDIR /go/src/app
COPY go.* ./
COPY *.go ./
COPY pkg pkg
COPY cmd cmd
RUN apk add --no-cache git
RUN apk add build-base
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /go/bin/app
RUN ls -lah
RUN pwd

#build stage (JS)
FROM node:12-alpine AS frontBuilder
WORKDIR /app
COPY ui ui
ENV PUBLIC_URL=/ui
RUN cd ui && npm run build

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=backBuilder /go/bin/app /app
RUN mkdir /ui
COPY --from=frontBuilder /app/ui/build /ui/build
ENV GIN_MODE=release
ENTRYPOINT ./app
LABEL Name=barker
EXPOSE 3000
