SRC    := $(shell find src/ -type f | grep -v "SUMMARY.md" | grep -v "intro.md")
UPDATE := <sub><sup>Updated:

.PHONY: $(SRC)

all: setup

release: generate check
	mdbook build

generate:
	./build.sh

check: $(SRC)

$(SRC):
	grep -q "$(UPDATE)" $@
