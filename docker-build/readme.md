


## build http docker image


1. build image
```shell
cd ..
docker build -t vvvincentli/http-server:1.0.8 . -f ci/Dockerfile-http-server


```

2. push image

```shell
#docker login

docker push vvvincentli/http-server:1.0.8
``` 

> https://hub.docker.com/r/vvvincentli/http-server

3. run container

```shell
 docker run -p 8088:8088 -d vvvincentli/http-server:1.0.8
 
 curl http://127.0.0.1:8088/healthz
```

> i'm alive. [2021-10-14T08:06:26+08:00]


4. nsenter
```shell
docker ps -a|grep vvvincent
ps -ef|grep vvvincent
nsenter -t <pid> -n ip addr

```
