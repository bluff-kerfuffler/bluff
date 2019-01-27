FROM golang:latest
EXPOSE 8080

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

# now copy your app to the proper build path
RUN mkdir -p $GOPATH/src/bluff
ADD . $GOPATH/src/bluff

# should be able to build now
WORKDIR $GOPATH/src/bluff
RUN go build -o bluff .
CMD ["/go/src/bluff/bluff"]
