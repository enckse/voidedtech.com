DIST := dist/
RSS  := $(DIST)rss/
SUBS := notebook crafts

.PHONY: $(SUBS)

all: build

build: $(TARGET)

release: clean main subsites

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)

main:
	mkdir -p $(RSS)
	go run www/main.go -rss $(RSS) -target $(DIST) -sites "$(SUBS)"

subsites: $(SUBS)

$(SUBS):
	cd $@ && make release
	cp -r $@/book $(DIST)$@
