FROM golang:1.12.4

RUN apt-get update && apt-get install -y unzip --no-install-recommends && \
    apt-get autoremove -y && apt-get clean -y && \
    wget -O dep.zip https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip && \
    echo '96c191251164b1404332793fb7d1e5d8de2641706b128bf8d65772363758f364  dep.zip' | sha256sum -c - && \
    unzip -d /usr/bin dep.zip && rm dep.zip

ADD . /go/src/github.com/kodebot/newsfeed

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

ENTRYPOINT revel run github.com/kodebot/newsfeed dev 9000

EXPOSE 9000