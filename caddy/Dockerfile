FROM mini/base

RUN curl -sLO https://github.com/mholt/caddy/releases/download/v0.11.2/caddy_v0.11.2_linux_amd64.tar.gz && tar -zxvf caddy_v0.11.2_linux_amd64.tar.gz && mv caddy /usr/bin/caddy && chmod 755 /usr/bin/caddy && rm -rf caddy*
ADD Caddyfile ./Caddyfile
#RUN apk update && apk add libcap && setcap cap_net_bind_service=+ep /usr/bin/caddy
EXPOSE 443 80 8080
CMD caddy