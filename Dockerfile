# builder image
FROM golang:1.19-alpine as builder
RUN apk add git
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/main.go

# final image
FROM alpine:3.16
COPY --from=builder /build/app .
COPY --from=builder /build/.env .

ENTRYPOINT [ "./app" ]