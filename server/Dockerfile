# build
FROM --platform=$BUILDPLATFORM docker.io/golang:bullseye AS build
RUN dpkg --add-architecture arm64 \
    && apt update -y \
    && apt install -y --no-install-recommends \
        ca-certificates \
        curl \
        gcc \
        gcc-aarch64-linux-gnu \
        libc6-dev-arm64-cross \
    && update-ca-certificates

ADD . /src
WORKDIR /src

ARG TARGETARCH
RUN if [ "$TARGETARCH" = "arm64" ]; then \
      export CC=aarch64-linux-gnu-gcc && \
      export PKG_CONFIG_PATH=/usr/lib/aarch64-linux-gnu/pkgconfig ; \
    fi; \
    GOOS=linux GOARCH=$TARGETARCH $TARGETFLAG make build

# final
FROM docker.io/debian:bullseye-slim
USER nobody
WORKDIR /workdir
COPY --from=build --chown=nobody:nobody /src/unshort.link /workdir/unshort.link
COPY --from=build /etc/ssl/certs /etc/ssl/certs
EXPOSE 8080
ENTRYPOINT ["/workdir/unshort.link"]
