DESTDIR := bin/
WEBROOT := www/

build:
	mkdir -p $(WEBROOT)
	cp *.css $(WEBROOT)
	cp ../../*.css $(WEBROOT)
	cp *.html $(WEBROOT)

install:
	mkdir -p $(DESTDIR)
	rsync -iacv $(WEBROOT) $(DESTDIR)
