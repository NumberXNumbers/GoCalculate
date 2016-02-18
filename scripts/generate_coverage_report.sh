#!/bin/bash
set -ev
cd $TRAVIS_BUILD_DIR
touch cover.cov
echo "mode: count" >> cover.cov
DIRS=$(go list ./...)
for dir in $DIRS; do
    go test $dir -coverprofile=cover.out -covermode=count
    if [ $? != 0 ]
    then
        exit 2
    fi
    if [ -f cover.out ]
    then
        sed -i '1d' cover.out
        cat cover.out >> cover.cov
        rm cover.out
    fi
done
