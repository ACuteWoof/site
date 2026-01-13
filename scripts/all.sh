#!/bin/sh

cd "${0%/*}"

./projects.sh && ./blog.sh && ./render.sh
npx prettier -w ../public/*
