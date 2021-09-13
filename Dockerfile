FROM golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lunchandgo .

FROM scratch
EXPOSE 8080
WORKDIR /app
COPY --from=builder /app/lunchandgo /usr/bin/
ENTRYPOINT ["lunchandgo"]