DIST := dist/

all: build

build: $(TARGET)

release: clean main subsites

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)

main:
	mkdir -p $(DIST)
	go run www/main.go -target $(DIST) -config www/site.json

subsites:
	cd notebook && make release
	cd crafts && make release
	cp -r notebook/book $(DIST)notebook	
	cp -r crafts/book $(DIST)crafts
