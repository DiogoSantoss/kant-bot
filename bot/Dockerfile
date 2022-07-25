# Build stage
FROM golang:alpine as builder 

RUN mkdir /build
COPY . /build
WORKDIR /build

RUN go build -o bot main.go

# Run stage
FROM alpine 

RUN mkdir /app
COPY .env /app
COPY --from=builder /build/bot /app
WORKDIR /app

CMD ["./bot"]


