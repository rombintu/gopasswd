#!/bin/bash

sudo docker ps -a | cut -f 1 -d " " | xargs docker rm