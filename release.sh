#!/bin/bash

now=$(date +"%Y%m%d%H%M%S")
docker buildx build --platform linux/amd64 -t epicmo/soruxgpt-saas-audit-oauth:latest --push .
docker buildx build --platform linux/amd64 -t epicmo/soruxgpt-saas-audit-oauth:latest --push .

echo "release success" $now
