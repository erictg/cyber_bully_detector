FROM golang:alpine

ENV GOPATH /app
ENV PATH $PATH:$GOPATH/bin
ENV GOBIN $GOPATH/bin
ENV DOCKER_INSIDE yee
RUN echo $PATH

RUN apk update
RUN apk add git
RUN apk add curl

RUN mkdir /app
RUN mkdir /app/bin

#golang dep replaces govendor
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . /app/src/bitcamp

WORKDIR /app/src/bitcamp

RUN dep ensure

RUN go install -v ./worker/main.go

CMD $GOBIN/main
