#!/usr/bin/env bash
set -ev
touch coverage.txt

for d in $(go list ./...); do
    go test -coverprofile=profile.out -covermode=count $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
