SRC    := $(shell find src/ -type f | grep -v "SUMMARY.md" | grep -v "intro.md")
UPDATE := <sub><sup>Updated:
TAG := $(shell git describe --abbrev=0)
SRC := $(shell git diff --name-only HEAD $(TAG) src/ | cut -d '/' -f 2-)

.PHONY: $(SRC)

all: setup

release: generate check
	mdbook build

generate:
	./build.sh

check: $(SRC)

$(SRC):
	grep -q "$(UPDATE) $(shell git log -n 1 --format=%as $@)</sup></sub>" $@
