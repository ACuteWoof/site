#!/bin/sh

cd "${0%/*}"

[ ! -d ../extra ] && mkdir ../extra

echo "" > ../extra/blog.html || touch ../extra/blog.html

TEXT=""

while IFS="|" read -r DATE TITLE DESCRIPTION LINK; do
	TEXT="<div class='blogpost'><h3>$TITLE</h3><div class='!text-xs'>$DATE</div><p>$DESCRIPTION</p><a href='/blog/$LINK' target='_blank'>Read</a></div>$TEXT"
done < ../src/blogposts.txt
echo $TEXT > ../extra/blog.html
