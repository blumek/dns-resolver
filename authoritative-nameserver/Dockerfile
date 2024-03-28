FROM golang:1.22.1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /authoritative_nameserver

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /authoritative_nameserver /authoritative_nameserver
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/authoritative_nameserver"]