# go mod vendor
# Run beatstep-demo:
# docker run --rm -ti --device=/dev/snd/seq beatstep-demo
FROM golang:alpine-3.8 AS build
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
	apk add --no-cache portmidi-dev build-base
WORKDIR /go/src/github.com/nogoegst/beatstep
COPY . .
RUN go build -v -o beatstep-demo ./cmd/beatstep-demo

FROM alpine:3.8
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
	apk add --no-cache portmidi
COPY --from=build /go/src/github.com/nogoegst/beatstep/beatstep-demo /

ENTRYPOINT ["/beatstep-demo"] 
