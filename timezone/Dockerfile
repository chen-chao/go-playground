FROM golang:latest

ENV GOPATH=/go

WORKDIR /go/src/go-time-docker

COPY . .

RUN go get ./... \
	&& go build -o go-time-docker .

ENTRYPOINT ["go-time-docker"]
