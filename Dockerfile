FROM golang:1.18.0-alpine3.15 as builder

LABEL author="echelonfoundation"

RUN apk add --update-cache \
    git \
    gcc \
    musl-dev \
    linux-headers \
    make \
    wget

RUN git clone https://github.com/echelonfoundation/echelon.git /echelon && \
    #chmod -R 755 /echelon && \
    chmod -R 755 /echelon
WORKDIR /echelon
RUN make install

# final image
FROM golang:1.18.0-alpine3.15

RUN mkdir -p /data

VOLUME ["/data"]

COPY --from=builder /go/bin/echelond /usr/local/bin/echelond

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["echelond"]