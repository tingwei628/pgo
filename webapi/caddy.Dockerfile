FROM caddy:2.6.2-builder AS builder
RUN xcaddy build \
  --with github.com/caddy-dns/cloudflare
WORKDIR /download
ADD https://github.com/yudai/gotty/releases/latest/download/gotty_linux_amd64.tar.gz gotty_linux_amd64.tar.gz
RUN tar -xf gotty_linux_amd64.tar.gz \
  && chmod a+x gotty


FROM caddy:2.6.2
WORKDIR /download
RUN apk add --update python3 py3-pip
RUN apk add libc6-compat
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
COPY --from=builder /download/gotty /usr/local/bin/gotty
RUN ln -s /usr/bin/python3 /usr/bin/python