auth OpenID

caddy2
rate limit

docker build -t pgo_webapi ../ -f Dockerfile

pgo_webapi = image tag

docker run --read-only -p 10001:9999 pgo_webapi

--read-only // volume is read only too
10001 => external port
9999 => internal port inside docker

delete all containers including its volumes
docker rm -vf $(docker ps -aq)

delete all the images
docker rmi -f $(docker images -aq)