RESOURCES := resources/
WEBROOT   := www
DESTDIR   := bin/

build:
	./build $(RESOURCES) $(WEBROOT)

install:
	mkdir -p $(DESTDIR)
	rsync -avci $(WEBROOT) $(DESTDIR)
