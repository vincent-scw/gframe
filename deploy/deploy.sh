#!/bin/bash

VERSION_NUMBER=$1
echo "Travis version number is "$VERSION_NUMBER

for f in ./deploy/apps/*.yaml
do
  template=`cat $f | sed "s/{{VERSION_NUMBER}}/$VERSION_NUMBER/g"`
  echo "$template" | kubectl apply -f -
done