#!/bin/bash

curl -X POST http://0.0.0.0:8080/images/upload \
  -F "file=@$1" \
  -H "Content-Type: multipart/form-data"
