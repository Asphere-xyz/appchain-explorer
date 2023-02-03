FROM golang:1.18

WORKDIR /go/src/github.com/Ankr-network/appchain-explorer
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
        -o /usr/local/bin/appchain-explorer \
        github.com/Ankr-network/appchain-explorer

FROM alpine:3.9

COPY --from=0 /usr/local/bin/appchain-explorer /usr/local/bin/appchain-explorer
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["appchain-explorer"]
