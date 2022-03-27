FROM golang:1.17.5-bullseye as builder
COPY src /var/src
RUN cd /var/src && \
    go build

FROM debian:bullseye-slim

COPY --from=builder /var/src/idea /usr/local/bin/idea