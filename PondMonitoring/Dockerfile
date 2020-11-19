FROM alpine:latest

RUN apk update \
    && apk add --update python3 \
    && apk add --update py3-pip \
    && pip3 install requests

COPY *.py /home/
WORKDIR /home/

ENV PATH_DATA /home/data/
ENV HOUR_TO_CAPTURE 23

ENTRYPOINT ["python3", "main.py"]