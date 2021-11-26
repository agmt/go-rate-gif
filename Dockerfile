FROM golang:1.15-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /rategif ./cmd/main.go



FROM alpine

WORKDIR /
COPY --from=build /rategif /rategif
COPY rategif.yml /

EXPOSE 8080
ENTRYPOINT ["/rategif"]
