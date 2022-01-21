FROM ubuntu:latest

WORKDIR /cmd

ADD ./proxy /cmd
ADD ./config.yaml /cmd
COPY ./docs /cmd/docs

CMD ["/cmd/proxy"]