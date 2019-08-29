

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