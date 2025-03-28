FROM golang:1.24.1-alpine3.21 AS build-stage
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app ./...

FROM genotek/r-base AS release-stage
COPY --from=build-stage /usr/local/bin/app /usr/local/bin/app
COPY do.sh /usr/local/bin/
ENV GIN_MODE=release
EXPOSE 8080
CMD ["app"]
