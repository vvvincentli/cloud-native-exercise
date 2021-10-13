


## build http docker image


1. build image
```shell
cd ..
docker build -t vvvincentli/http-server:1.0.0 . -f ci/Dockerfile-http-server


```

2. push image

```shell
#docker login

docker push vvvincentli/http-server:1.0.0
``` 

3. run container

```shell
 docker run -p 8088:8088 -d vvvincentli/http-server:1.0.0
```

4. nsenter
```shell
docker ps -a|grep vvvincent
ps -ef|grep vvvincent
nsenter -t <pid> -n ip addr

```
