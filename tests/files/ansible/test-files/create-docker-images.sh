#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="earthdollar/cppjit-testrunner" docker-cppjit
docker build --tag="earthdollar/python-testrunner" docker-python
docker build --tag="earthdollar/go-testrunner" docker-go
