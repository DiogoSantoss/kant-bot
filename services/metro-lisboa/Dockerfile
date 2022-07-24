# Build stage
FROM golang:alpine as builder 

RUN mkdir /build
COPY . /build
WORKDIR /build

RUN go build -o metro-lisboa main.go

# Run stage
FROM alpine 

RUN mkdir /app
COPY .env /app
COPY --from=builder /build/metro-lisboa /app
WORKDIR /app

EXPOSE 8080
CMD ["./metro-lisboa"]


