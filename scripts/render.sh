#!/bin/sh

cd "${0%/*}"

[ ! -d ../public ] && mkdir ../public

echo "" > ../public/index.html || touch ../public/index.html

cp -r ../src/public/* ../public/.
cp -r ../src/public/.* ../public/.
cat ../src/top2prtop.html >> ../public/index.html
cat ../extra/projects.html >> ../public/index.html
cat ../src/prbtm2bltop.html >> ../public/index.html
cat ../extra/blog.html >> ../public/index.html
cat ../src/blbtm2btm.html >> ../public/index.html

yarn run prettier -w  ${PWD%/*}/public/index.html
yarn build
