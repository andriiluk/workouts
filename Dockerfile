FROM golang:1.18 as builder
WORKDIR /app
# RUN apk --no-cache add ca-certificates git tzdata
COPY ./ /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server .

FROM scratch
WORKDIR /app
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/* /app/
ENV TZ=Europe/Kiev
EXPOSE 8080
CMD ["/app/server", "run" ]