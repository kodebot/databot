FROM golang:1.12.4

RUN apt-get update && apt-get install -y unzip --no-install-recommends && \
    apt-get autoremove -y && apt-get clean -y && \
    wget -O dep.zip https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip && \
    echo '96c191251164b1404332793fb7d1e5d8de2641706b128bf8d65772363758f364  dep.zip' | sha256sum -c - && \
    unzip -d /usr/bin dep.zip && rm dep.zip

ENV env=DEV
# DEV mongodb connection string for http cache and export store - change this as appropriate
ENV DATABOT_CACHE_DB_CON_STR=mongodb://localhost:27017
ENV DATABOT_EXPORT_TO_DB_CON_STR=mongodb://localhost:27017

ADD . /go/src/github.com/kodebot/databot

ENTRYPOINT go run /go/src/github.com/kodebot/databot/cmd/databot/main.go -stderrthreshold=INFO -feedconfigpath=/go/src/github.com/kodebot/databot/testdata/feeds/ready/

EXPOSE 9025