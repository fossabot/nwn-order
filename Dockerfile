FROM golang:1.11.0 as nwn-order-builder
RUN apt update \
    && apt upgrade -y \
    && rm -r /var/lib/apt/lists /var/cache/apt \
    && git clone https://github.com/Urothis/nwn-order.git \
    && cd nwn-order \
    && go mod download \
    && go build -o ./bin/order \
    && mv bin/order /usr/local/bin/
FROM ubuntu:latest
LABEL maintainer "urothis@gmail.com"
COPY --from=nwn-order-builder /usr/local/bin/order /usr/local/bin/order
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get clean \
    && chmod +x ./usr/local/bin/order
ENTRYPOINT [ "order" ]