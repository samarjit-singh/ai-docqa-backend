#!/bin/bash
# build.sh
prisma generate
go build -tags netgo -ldflags '-s -w' -o app
