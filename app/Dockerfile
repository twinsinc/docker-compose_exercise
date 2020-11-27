FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

FROM golang:1.13 as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY . /app/
RUN go build -o go-app

FROM scratch
EXPOSE 8080
COPY --from=certs /etc/ssl/certs/* /etc/ssl/certs/
COPY --from=builder /app/go-app /
ENV IMAGES_URL http://localhost:3000
ENTRYPOINT ["/go-app"]
