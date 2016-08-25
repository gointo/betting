FROM golang:1.7.0-alpine
MAINTAINER Maxime Vaude <maxime.vaude@gmail.com> (@mvaude)

RUN apk update && apk add --update --no-cache git && rm -rf /var/cache/apk/*

ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /goapath/src/betting
COPY . /gopath/src/betting
RUN go get betting

CMD ["betting", "https://userstream.twitter.com/1.1/user.json"]
ENTRYPOINT ["/gopath/bin/betting"]
