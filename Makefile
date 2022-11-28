DIST := dist/
SUBS := notebook crafts

all: build

build: $(TARGET)

release: clean main subsites

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)

main:
	mkdir -p $(DIST)
	go run www/main.go -target $(DIST) -config www/site.json

subsites: $(SUBS)

$(SUBS):
	cd $@ && make release
	cp -r $@/book $(DIST)$@
