FROM golang:alpine AS builder

RUN apk add --no-cache git bash

ARG TARGETARCH
ARG VERSION=dev
ARG BUILD_DATE
ARG COMMIT_ID

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build \
    -ldflags "-s -w -X 'main.Version=${VERSION}' -X 'main.CommitID=${COMMIT_ID}' -X 'main.BuildDate=${BUILD_DATE}'" \
    -o /build/binary/casimg ./src

FROM alpine:latest

ARG VERSION=dev
ARG BUILD_DATE
ARG COMMIT_ID
ARG LICENSE=MIT

LABEL maintainer="casapps <casjay@casjaysdev.pro>" \
      org.opencontainers.image.vendor="casapps" \
      org.opencontainers.image.authors="casapps" \
      org.opencontainers.image.title="casimg" \
      org.opencontainers.image.base.name="casimg" \
      org.opencontainers.image.description="Self-hosted file conversion service" \
      org.opencontainers.image.url="https://github.com/casapps/casimg" \
      org.opencontainers.image.source="https://github.com/casapps/casimg" \
      org.opencontainers.image.documentation="https://github.com/casapps/casimg" \
      org.opencontainers.image.vcs-type="Git" \
      com.github.containers.toolbox="false"

LABEL org.opencontainers.image.licenses="${LICENSE}" \
      org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.schema-version="${VERSION}" \
      org.opencontainers.image.revision="${COMMIT_ID}"

RUN apk add --no-cache \
    git \
    curl \
    bash \
    tini \
    tor \
    imagemagick \
    ffmpeg \
    libreoffice \
    ghostscript

RUN mkdir -p /config /config/security /data/db /data/log /data/tor /data/backup /data/uploads /data/converted

COPY --from=builder /build/binary/casimg /usr/local/bin/casimg

COPY docker/rootfs/ /

COPY docker/Dockerfile /root/Dockerfile

RUN chmod 755 /usr/local/bin/*

ENV MODE=development \
    TZ=America/New_York \
    PORT=80 \
    ADDRESS=0.0.0.0

EXPOSE 80

STOPSIGNAL SIGRTMIN+3

HEALTHCHECK --start-period=10m --interval=5m --timeout=15s --retries=3 \
    CMD /usr/local/bin/casimg --status || exit 1

ENTRYPOINT [ "tini", "-p", "SIGTERM", "--", "/usr/local/bin/entrypoint.sh" ]
