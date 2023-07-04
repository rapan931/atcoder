#!/bin/bash

dir=`cat ./contest`
echo -e "Contest: $dir"
echo -n "Question?"
read q

mkdir -p _result/_$dir/$q
cp -p main.go _result/_$dir/$q

#read -p "Ready to commit, hit enter."
#
#git add _result/_$dir/$q/$sendfile
#git commit -m "$dir $q"

# cp -i _template/main.go ./main.go
