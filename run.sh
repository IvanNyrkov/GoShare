#!/usr/bin/env bash
docker build -t ivannyrkov/goshare .;
docker run --rm -p 80:80 -t -i ivannyrkov/goshare .;