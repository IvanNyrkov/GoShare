#!/usr/bin/env bash
docker build -t ivannyrkov/goshare .;
docker run --rm -p 80:80 \
    --env CONFIG_FILE=config-docker.json \
    -t -i ivannyrkov/goshare .;