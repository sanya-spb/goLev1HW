#!/bin/bash

shuf -i 0-255 -n 10 -r | go run srs_01.go
