#!/usr/bin/env bash

for src in */*.go; do
	gofmt -w $src
done
