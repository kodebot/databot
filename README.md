## web
run `revel run -a  github.com/kodebot/newsfeed`

## docker

go into newsfeed folder

run `sudo docker build -t newsfeed .`
run `sudo docker run -p 9020:9020 newsfeed` or `sudo docker run -it -p 9020:9020 newsfeed` for interactive connection


## config
mongo db url should be set to private IP ( this is not good as the IP can change) -- THIS IS NOT WORKING
set the mongo db url to the container's ip address using the follwing command and use it
`docker inspect --format '{{ .NetworkSettings.IPAddress }}' <container name or id>`

## mongo
run `sudo docker run --name datastore -d -p 27017:27017 -v ~/data:/data/db mongo`