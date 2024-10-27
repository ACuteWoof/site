#!/bin/sh

cd "${0%/*}"


[ ! -d ../extra ] && mkdir ../extra

echo "" > ../extra/projects.html || touch ../extra/projects.html

TEXT=""

while IFS="|" read -r TITLE DESCRIPTION LINK IMAGE; do
	TEXT="$TEXT<div class='project'><img src='$IMAGE' /><div class='text'><span class='text-lg font-bold'>$TITLE</span><span>$DESCRIPTION</span><span><a href='$LINK'>View</a></span></div></div>"
done < ../src/projects.txt
echo $TEXT > ../extra/projects.html
