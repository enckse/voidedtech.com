RESOURCES := resources/
WEBROOT   := www
DESTDIR   := bin/

build:
	./configure $(RESOURCES) $(WEBROOT)

install:
	mkdir -p $(DESTDIR)
	mv $(WEBROOT) $(DESTDIR)
