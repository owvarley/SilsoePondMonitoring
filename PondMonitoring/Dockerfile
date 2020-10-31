FROM debian:bullseye-slim

RUN apt-get update
RUN apt-get install -y python3 \
                       python3-pip \
                       python3-setuptools

RUN pip3 install requests

COPY pull_rainfall.py /home/pull_rainfall.py
WORKDIR /home/

ENTRYPOINT ["python3", "pull_rainfall.py"]