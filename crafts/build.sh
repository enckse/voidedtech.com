#!/bin/bash
SRC="src/"
INDEX="$SRC/SITEINDEX.md"

_format() {
    echo "- [$1]($2)"
}

_build() {
    local f bname href
    for f in $(ls $SRC/$1); do
        bname=$(echo "$f" | rev | cut -d "." -f 2- | rev)
        echo
        echo "## $bname"
        echo
        href="$1/$f"
        echo "<a href=\"$href\"><img src=\"$href\" alt=\"$bname\" width=100 />"
        echo
        echo "- [$bname ($1)]($yr.html#$bname)" >> $INDEX
    done
}

_summary() {
    local md yr
    echo "# Summary"
    echo
    rm -f $INDEX
    _format "Intro" "./intro.md"
    for yr in $(ls $SRC | grep "^20" | grep -v "\.md" | sort -r -n); do
        md="$yr.md"
        _build "$yr" > $SRC$md
        _format "$yr" "./$md"
    done
    sort -o $INDEX $INDEX
    sed -i '1 i\# Index\n' $INDEX
    {
        echo
        echo "<sub><sup>Updated: $(date +%Y-%m-%d)</sup></sub>"
    } >> $INDEX
    _format "Index" "./SITEINDEX.md"
}

_summary > ${SRC}SUMMARY.md
