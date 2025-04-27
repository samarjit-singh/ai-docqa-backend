#!/bin/bash
set -e

echo "Installing Prisma CLI..."
npm install -g prisma

echo "Generating Prisma client..."
npx prisma generate

echo "Building Go app..."
go build -tags netgo -ldflags '-s -w' -o app
