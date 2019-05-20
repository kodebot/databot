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

## nginx

run `sudo docker run --name kodebot-nginx --mount type=bind,source=/home/kodebot/nginx/www,target=/usr/share/nginx/html,readonly --mount type=bind,source=/home/kodebot/nginx/conf.d,target=/etc/nginx/conf.d,readonly -p 80:80 -d nginx`

```
When the container is created, you can mount a local directory on the Docker host to a directory in the container. The NGINX image uses the default NGINX configuration, which uses /usr/share/nginx/html as the container’s root directory and puts configuration files in /etc/nginx. For a Docker host with content in the local directory /var/www and configuration files in /var/nginx/conf, run the command:

$ docker run --name mynginx2 --mount type=bind source=/var/www,target=/usr/share/nginx/html,readonly --mount source=/var/nginx/conf,target=/etc/nginx/conf,readonly -p 80:80 -d nginx
Any change made to the files in the local directories /var/www and /var/nginx/conf on the Docker host are reflected in the directories /usr/share/nginx/html and /etc/nginx in the container. The readonly option means these directories can be changed only on the Docker host, not from within the container.
```

nginx logs are available in the following directories inside the container 

`/var/log/nginx/access.log`
`/var/log/nginx/error.log`

create a config file inside the conf.d directory with the following content (file extension must be .conf)

use the private IP address of the server

```
server {
    listen 80;
    server_name thegiftapp.kodebot.com;
    location / {
        proxy_pass http://10.0.0.4:9000;
    }
}
server {
    listen 80;
    server_name api.newsfeed.kodebot.com;
    location / {
        proxy_pass http://10.0.0.4:9020;
    }
}
```

