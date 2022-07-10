RESOURCES := resources/
WEBROOT   := www
DESTDIR   := bin/

build:
	./configure $(RESOURCES) $(WEBROOT)

install:
	mkdir -p $(DESTDIR)
	rsync -avci $(WEBROOT) $(DESTDIR)
