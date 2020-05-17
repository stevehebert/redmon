#!/bin/bash 
docker build -t redmon .
docker run --restart always --detach -t redmon
