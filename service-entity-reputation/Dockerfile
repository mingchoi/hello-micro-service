# build stage
FROM golang:1.13 AS build
WORKDIR /src
ADD . /src
RUN go get ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app service-entity-reputation/cmd/server/main.go

#  runtime stage
FROM alpine
WORKDIR /
COPY --from=build /src/app /app
ENTRYPOINT /app