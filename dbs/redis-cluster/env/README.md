1. modify 'docker-compose.yml' as your env ip/port.
2. $ docker-compose up -d
3. create a new redis-cluster:
$ docker run -it --network=host zero-chan/redis-cluster-manager:1 create --replicas 1 127.0.0.1:15379 127.0.0.1:15380
