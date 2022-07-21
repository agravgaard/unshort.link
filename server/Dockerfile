# build
FROM --platform=$BUILDPLATFORM docker.io/golang:bullseye AS build
RUN dpkg --add-architecture arm64 \
    && apt update -y \
    && apt install -y --no-install-recommends \
        curl \
        gcc \
        gcc-aarch64-linux-gnu \
        libc6-dev-arm64-cross
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
COPY --from=build /src/unshort.link /src/unshort.link
EXPOSE 8080
ENTRYPOINT ["/src/unshort.link"]
