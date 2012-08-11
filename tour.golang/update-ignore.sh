#!/usr/bin/env bash

rm .gitignore
for i in *; do
	if [ -d "$i" ]; then
		ignore="$i/$i"
		if [[ $1 == '-v' || $1 == '--verbose' ]]; then
			echo "Adding $ignore to ignore list"
		fi
		echo "$i/$i" >> .gitignore
	fi
done
