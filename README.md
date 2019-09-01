

## run test harness

`cd cmd\databot-testharness`

`go run main.go`

>test harness runs in port 9022 by default. browse to http://localhost:9022 to use the test harness

## run the main program once
`go run main.go -runonce -logtostderr=true -stderrthreshold=INFO -feedconfigpath=../../testdata/feeds/ready/`

`feedconfigpath` must be relative to current location

## run scheduler
`go run main.go -logtostderr=true -stderrthreshold=INFO -feedconfigpath=../../testdata/feeds/ready/`

`feedconfigpath` must be relative to current location


# run in docker - locally

`docker build -t databot .`


# run in docker - production

go to the server that hosts docker
clone the project files from github

then run the following to build with prod file

`docker build -t databot -f Dockerfile.PROD .`

to run the container

`docker run -t -d -p:9025:9025 databot`