#!/usr/bin/env bash

MAXPAGE=50
RESOURCES="$1"
WEBROOT="$2"
BIN="bin/"
RAW=".raw"
INDEX="index$RAW"
ABOUT="about$RAW"

_checkdir() {
    if [ -z "$1" ]; then
        echo "empty parameter"
        exit 1
    fi
}

_checkdir "$RESOURCES"
_checkdir "$WEBROOT"

rm -rf $BIN
mkdir -p $BIN

current=1
page=1
for year in $(find "$RESOURCES" -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort -r); do
    echo "processing resources for $year"
    for image in $(find "$RESOURCES/$year" -type f -exec basename {} \; | sort); do
        if [[ $current -gt $MAXPAGE ]]; then
            current=1
            page=$((page+1))
        fi
        offset=$year
        name=$(echo "$image" | cut -d "." -f 1)
        src=$offset/$image
        div="<div class='entry'><div class='imgdetail'><a href='$src'><img src='$src' loading='lazy' alt='$name'></a></div><div class='details'><b>name: </b>$name<br /><b>year: </b>$year</div><br /></div><br />";
        echo "$div" >> $BIN$page$RAW
        current=$((current+1))
    done
done

mv ${BIN}1$RAW $BIN$INDEX
links=$(find $BIN -type f -name "*$RAW" -exec basename {} \; | grep -v "index" | sort)
links=$(printf "%s\n%s\n%s" "$INDEX" "$links" "$ABOUT")
updated=$(date +%Y-%m-%d)
sed "s/{DATE}/$updated/g" about.html > ${BIN}about.raw
for p in $(find "$BIN" -type f -name "*$RAW" -exec basename {} \; | sort); do
    echo "making: $BIN$p"
    content=$(tr '\n' ' ' < "$BIN$p")
    linking=""
    page="$BIN${p//$RAW/.html}"
    for l in $links; do
        link=$(echo "$l" | cut -d "." -f 1)
        disp="$link"
        if [[ "$link" == "index" ]]; then
            disp="1"
        fi
        is_about=0
        if [[ "$link" == "about" ]]; then
            is_about=1
        fi
        if [[ "$l" != "$p" ]]; then
            disp="<a class='pager' href='$link.html'>$disp</a>"
        fi
        if [ $is_about -eq 1 ]; then
            disp="<div class='about'>$disp</div>"
        fi
        linking="$linking $disp"
    done
    sed "s|{LINKS}|$linking|g;s|{CONTENT}|$content|g" index.html > "$page"
done

rm -f $BIN*$RAW
cp ./*.css "$BIN"
# shellcheck disable=SC2086
cp -r $RESOURCES* "$BIN"
mkdir -p "$WEBROOT"
rsync -aricv --delete-after "$BIN" "$WEBROOT"
