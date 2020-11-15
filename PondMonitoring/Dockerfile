FROM debian:bullseye-slim

RUN apt-get update
RUN apt-get install -y python3 \
                       python3-pip \
                       python3-setuptools

RUN pip3 install requests

COPY *.py /home/
WORKDIR /home/

RUN python3 -m unittest discover -v

ENV PATH_DATA /home/data/

ENTRYPOINT ["python3", "main.py"]