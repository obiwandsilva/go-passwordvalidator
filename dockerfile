FROM golang:1.15-alpine3.12 AS builder

WORKDIR /go-passwordvalidator
COPY . .
RUN go mod download
RUN GOOS=linux go build -o bin/passwordvalidatorservice cmd/passwordvalidator/main.go

FROM alpine:3.12.3
WORKDIR /app
COPY --from=builder /go-passwordvalidator/bin/passwordvalidatorservice .
EXPOSE 7000
CMD ./passwordvalidatorservice
