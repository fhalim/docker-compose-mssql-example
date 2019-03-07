FROM golang:1.11.0 AS build

WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /build/myapp

FROM alpine:3.8
COPY --from=build /build/myapp /usr/local/bin/myapp
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/myapp"]
