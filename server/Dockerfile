# build
FROM --platform=$BUILDPLATFORM docker.io/golang:alpine AS build
RUN apk --no-cache add curl build-base gcc
ADD . /src
WORKDIR /src

ARG TARGETARCH
RUN export PATH=${PATH}:$HOME/go/bin \
    && echo $PATH \
    && GOOS=linux GOARCH=$TARGETARCH make build

# final
FROM docker.io/alpine
COPY --from=build /src/unshort.link /src/unshort.link
EXPOSE 8080
ENTRYPOINT ["/src/unshort.link"]
