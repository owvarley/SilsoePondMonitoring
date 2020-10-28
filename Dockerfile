FROM debian:bullseye-slim

RUN apt-get update
RUN apt-get install -y python3 \
					   python3-pip \
					   python3-setuptools

ENTRYPOINT ["/bin/bash"]