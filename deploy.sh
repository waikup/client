#!/bin/bash

GOARCH=arm GOOS=linux go build
scp wakeup-client pi@192.168.1.140:wakeup