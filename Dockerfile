FROM golang:1.23.2-bookworm AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o main .

FROM debian:12.7-slim AS runner

ARG version="dev"
ARG revision="dev"

ENV MODE="production"
ENV GIT_REF=$version

LABEL org.opencontainers.image.authors="Jack Gledhill"
LABEL org.opencontainers.image.description=""
LABEL org.opencontainers.image.documentation="https://github.com/Jack-Gledhill/robojack"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.revision=$revision
LABEL org.opencontainers.image.source="https://github.com/Jack-Gledhill/robojack"
LABEL org.opencontainers.image.title="RoboJack"
LABEL org.opencontainers.image.url="https://github.com/Jack-Gledhill/robojack"
LABEL org.opencontainers.image.version=$version

EXPOSE 8080

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /app
COPY --from=builder /app/ /app

ENTRYPOINT [ "./main" ]