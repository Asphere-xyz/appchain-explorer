FROM golang:1.18 AS build_backend
WORKDIR /build
COPY ./cmd /build/cmd
COPY ./shared /build/shared
COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum
RUN go build ./cmd/node

FROM node:16-alpine AS build_frontend
WORKDIR /build
RUN apk update && apk add python3 alpine-sdk
COPY ./web /build
RUN cd /build
RUN yarn
RUN yarn build

FROM golang:1.18 AS run
WORKDIR /explorer
RUN apt update && apt install -y gettext
COPY --from=build_backend /build/node /explorer/node
COPY --from=build_frontend /build/build /explorer/web/build
COPY ./docker/entrypoint.sh /explorer/entrypoint.sh
RUN chmod +x /explorer/entrypoint.sh
RUN chmod +x /explorer/node
CMD /explorer/entrypoint.sh