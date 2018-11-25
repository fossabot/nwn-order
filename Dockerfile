FROM golang:1.11.0 as nwn-order-builder
RUN apt update \
    && apt upgrade -y \
    && rm -r /var/lib/apt/lists /var/cache/apt \
    && git clone https://github.com/Urothis/nwn-order.git \
    && cd nwn-order \
    && go mod download \
    && go build -o ./bin/order \
    && mv bin/* /usr/local/bin/

FROM ubuntu:latest
LABEL maintainer "urothis@gmail.com"
# copy go
COPY --from=nwn-order-builder /usr/local/bin/ /usr/local/bin/

# run order-cli
ENTRYPOINT [ "order" ]