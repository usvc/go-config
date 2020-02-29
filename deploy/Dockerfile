FROM golang:1.14-buster AS builder
RUN apt-get update \
  && apt-get upgrade -y
RUN apt-get install make git
WORKDIR /app
COPY . .
RUN make deps
RUN make build_production

FROM gcr.io/distroless/base-debian10
COPY --from=builder /app/bin/config_linux_amd64 /app
ENTRYPOINT ["/app"]