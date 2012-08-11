#!/usr/bin/env bash

for i in *; do
	if [ -d "$i" ]; then
		if [[ $1 == '-v' || $1 == '--verbose' ]]; then
			echo "Cleaning $i ..."
		fi
		cd $i
		go clean
		cd ..
	fi
done
