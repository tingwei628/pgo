## Quick start
create and start containers \
`docker-compose -f docker-compose-webapi.yml up`

stop and remove containers \
`docker-compose -f docker-compose-webapi.yml down`

## Setup

Caddy2, Cloudflare, Go, Docker \
[your .env file]

### Docker other commands

build a docker image, pgo_webapi = image tag \
`docker build -t pgo_webapi ../ -f Dockerfile`

run \
`docker run --read-only -p 10001:9999 pgo_webapi` \
`docker run -d -p 10001:9999 pgo_webapi` // -d detach mode

delete all containers including its volumes \
`docker rm -vf $(docker ps -aq)`

stop all docker containers \
`docker stop $(docker ps -a -q)`
remove all docker containers \
`docker rm $(docker ps -a -q)`
delete all the images \
`docker rmi -f $(docker images -aq)`

docker tag local-image:tagname new-repo:tagname \
`docker tag pgo_webapi tingwei628/pgo_webapi` \
docker push new-repo:tagname \
`docker push tingwei628/pgo_webapi`

build caddy docker image \
`docker build -f caddy.Dockerfile .`


### todo
web api crud \
auth OpenID \
rate limit
