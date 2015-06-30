FROM google/golang
MAINTAINER Sakeven "sakeven.jiang@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/src/golang-mongo-sample

RUN go get -t golang-mongo-sample
RUN go install golang-mongo-sample

EXPOSE 80    
CMD ["/gopath/app/bin/golang-mongo-sample"]