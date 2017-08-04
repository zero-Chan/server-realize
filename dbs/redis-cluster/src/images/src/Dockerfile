FROM redis:3.0.7

ENV TZ Asia/Shanghai

COPY redis.conf /usr/local/etc/redis/redis.conf

ENTRYPOINT ["redis-server"]

CMD ["/usr/local/etc/redis/redis.conf"]
