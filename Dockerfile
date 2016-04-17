FROM alpine
MAINTAINER Thinh Tran <duythinht@gmail.com>
ADD ./dist/whereami /bin/
ENV MARTINI_ENV production
EXPOSE 3000
ENTRYPOINT ["/bin/whereami"]
