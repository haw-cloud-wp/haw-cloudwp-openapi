FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -o openapi .

FROM ubuntu:latest AS runtime
RUN mkdir /app
WORKDIR /app
COPY --from=build /app/openapi /app/
RUN apt update && apt install --reinstall -y ca-certificates
EXPOSE 80/tcp
ENTRYPOINT ["/app/openapi"]