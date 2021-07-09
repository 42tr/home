FROM alpine

COPY ./home /tmp/home

WORKDIR /tmp/

RUN chmod +x home