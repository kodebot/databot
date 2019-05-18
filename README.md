## web
run `revel run -a  github.com/kodebot/newsfeed`

## docker

go into newsfeed folder

run `sudo docker build -t newsfeed .`
run `sudo docker run -p 9020:9020 newsfeed`


## config
mongo db url should be set to private IP ( this is not good as the IP can change)

## mongo
run `sudo docker run -d -p 27017:27107 -v ~/data:/data/db mongo --name datastore`