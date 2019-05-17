## web
run `revel run -a  github.com/kodebot/newsfeed`

## docker

go into newsfeed folder

run `docker build -t newsfeed .`
run `docker run -p 9020:9020 newsfeed`


## config
mongo db url should be set to private IP ( this is not good as the IP can change)