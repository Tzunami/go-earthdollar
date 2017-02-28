FROM ubuntu:14.04

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update
RUN apt-get upgrade -y

RUN apt-get install -qy curl git python2.7 python-pip python-dev

# this is a workaround, to make sure that docker's cache is invalidated whenever the git repo changes
ADD https://api.github.com/repos/earthdollar/pyearthdollar/git/refs/heads/develop unused.txt 

RUN git clone --branch develop --recursive https://github.com/Tzunami/pyearthdollar.git

RUN cd pyearthdollar && curl https://bootstrap.pypa.io/bootstrap-buildout.py | python

RUN cd pyearthdollar && bin/buildout

#default port for incoming requests
EXPOSE 20203

WORKDIR /pyearthdollar 

ENTRYPOINT ["bin/python", "tests/test_vm.py"]
