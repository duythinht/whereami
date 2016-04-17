FROM alpine
# Update package index and install go + git
RUN apk add --update go git

# Set up GOPATH
RUN mkdir /go
ENV GOPATH /go
ENV GO15VENDOREXPERIMENT 1
WORKDIR /go
ADD ./ src/duythinht/whoami
RUN go install duythinht/whoami
EXPOSE 3000
RUN ls /go/bin
CMD ['/go/bin/whoami']
