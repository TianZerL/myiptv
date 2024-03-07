FROM --platform=$BUILDPLATFORM golang:alpine AS builder
WORKDIR /build/src
COPY server .
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -tags timetzdata -o /build/bin/myiptv .

FROM scratch
WORKDIR /app
COPY --from=builder /build/bin/myiptv .
COPY web web
COPY images/favicon.ico web/
COPY LICENSE web/
ENTRYPOINT ["/app/myiptv"]
