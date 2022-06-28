FROM golang:1.17 AS build
WORKDIR /build
COPY ../cmd /build/cmd
COPY ../shared /build/shared
COPY ../go.mod /build/go.mod
COPY ./go.sum /build/go.sum
RUN go build ./cmd/node

FROM golang:1.17 AS run
WORKDIR /app
COPY ./config.yaml /app/config.yaml
COPY --from=build /build/node /app/node
RUN chmod +x /app/node
CMD ["/app/node"]