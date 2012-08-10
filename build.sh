#!/usr/bin/env bash

for i in *; do
	if [ -d $i ]; then
		cd "$i"
		go build
		cd ..
	fi
done
