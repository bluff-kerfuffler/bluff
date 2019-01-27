FROM ubuntu:latest

RUN curl https://getcaddy.com | bash -s personal
ADD Caddyfile ./Caddyfile
CMD ["/usr/local/bin/caddy -conf Caddyfile -agree"]