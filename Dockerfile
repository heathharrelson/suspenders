ARG ALPINE_VERSION=3.13

FROM node:lts AS assets
RUN mkdir /build
WORKDIR /build
COPY ui ui
RUN npm install --prefix=ui && \
  npm run build --prefix=ui

FROM golang:1.15-alpine${ALPINE_VERSION} AS builder
RUN mkdir /build
WORKDIR /build
COPY . .
RUN apk update && \
  apk add --no-cache build-base git && \
  make suspenders

FROM alpine:${ALPINE_VERSION}
COPY --from=assets /build/ui/dist /ui/dist
COPY --from=builder /build/suspenders /suspenders
RUN \
  apk update && \
  apk add --no-cache \
  bash \
  openssl-dev \
  tzdata

EXPOSE 8080
USER nobody
CMD [ "/suspenders" ]
