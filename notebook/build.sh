#!/bin/bash

_parse() {
    local rel title
    rel=${2/src\//\./}
    title=$(head -n 1 "$2")
    _format "$1" "$title" "$rel"
}

_format() {
    echo "$1- [$2]($3)"
}

_summary() {
    local src dir
    echo "# Summary"
    echo
    _format "" "Intro" "./intro.md"
    for src in $(find src/ -mindepth 2 -type f -name "intro.md" | sort); do
        _parse "" "$src"
        dir=$(dirname "$src")
        for f in $(find "$dir" -type f -name "*.md" | grep -v "intro.md" | sort); do
            _parse "    " "$f"
        done
    done
}

_summary > src/SUMMARY.md
