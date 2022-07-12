RESOURCES := resources/
WEBROOT   := www
DESTDIR   := bin/

build:
	./build.sh $(RESOURCES) $(WEBROOT)

install:
	mkdir -p $(DESTDIR)
	rsync -avci $(WEBROOT) $(DESTDIR)
