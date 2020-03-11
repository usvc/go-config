FROM golang:1.14-buster AS builder
RUN apt-get update \
  && apt-get upgrade -y
RUN apt-get install -y make git upx
WORKDIR /app
COPY . .
RUN make deps
RUN make build_production
RUN make compress

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/bin/config_linux_amd64 /config
ENTRYPOINT ["/config"]
