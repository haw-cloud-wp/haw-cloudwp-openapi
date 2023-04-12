FROM golang:1.19 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN CGO_ENABLED=0 GOOS=linux go build -o openapi .

FROM scratch AS runtime
COPY --from=build /app/openapi ./
EXPOSE 3001/tcp
ENTRYPOINT ["./openapi"]
