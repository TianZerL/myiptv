FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /build/src
COPY server .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /build/out/myiptv .

FROM scratch
WORKDIR /app
COPY --from=build /build/out/myiptv .
COPY web web
ENTRYPOINT ["/app/myiptv"]
