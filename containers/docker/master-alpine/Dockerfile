FROM alpine:3.3

RUN \
  apk add --update go git make gcc musl-dev                && \
  git clone https://github.com/Tzunami/go-earthdollar && \
  (cd go-earthdollar && make ged)                            && \
  cp go-earthdollar/build/bin/geth /geth                      && \
  apk del go git make gcc musl-dev                         && \
  rm -rf /go-earthdollar && rm -rf /var/cache/apk/*

EXPOSE 8811
EXPOSE 20203

ENTRYPOINT ["/geth"]
