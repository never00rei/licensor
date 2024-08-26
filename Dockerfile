FROM golang:1.19-buster as builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN cd cmd/licensor-api && go build -v -o ../../licensor


FROM debian:buster-slim

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/licensor /app/server

ENV DB_HOST=
ENV DB_PORT=
ENV DB_USER=
ENV DB_PASSWORD=
ENV DB_DATABASE=

CMD ["/app/server"]
