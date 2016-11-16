FROM golang:1.7.3
MAINTAINER Valentin Pichard <valentin.pichard@corp.ovh.com>

# Add local sources
ADD . /go/src/github.com/w3st3ry/tekleader
WORKDIR /go/src/github.com/w3st3ry/tekleader

# Get vendoring tool and sync dependencies
RUN go get -u github.com/kardianos/govendor
RUN govendor sync

RUN go install

ENTRYPOINT ["tekleader"]

CMD ["leader"]
