# redis-demo
Just for a quick demo about redis

How to create container on docker?
```SHELL
docker-compose -f docker-compose.yaml  up -d
```

How to exec container?
```SHELL
docker container exec -it redis-demo /bin/sh 
```

Connect CLI
```SHELLL
redis-cli -h localhost
```
