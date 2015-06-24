FROM google/golang
MAINTAINER Golfen Guo "golfen.guo@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/src/golang-mongo-sample

RUN go get -t golang-mongo-sample
RUN go install golang-mongo-sample

EXPOSE 80
CMD ["/gopath/app/bin/golang-mongo-sample"]