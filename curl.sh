#!/bin/bash

file=$1
curl -X POST http://localhost:8080/user/batch --data "@${file}" -H "Content-Type: application/json"
