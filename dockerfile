
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/amoschen01.exaroton.me
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/amoschen01.exaroton.me /amoschen01.exaroton.me
ENTRYPOINT ./amoschen01.exaroton.me
LABEL Name=amoschen01.exaroton.me Version=0.0.1
EXPOSE 8080
