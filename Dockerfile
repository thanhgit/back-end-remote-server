FROM gcr.io/bitnami-containers/minideb-extras:jessie-r14

MAINTAINER Thanh Nguyen <thanh29695@gmail.com>

ADD . /app/

USER bitnami

WORKDIR /app

EXPOSE 7000

CMD ["./remote-server"]