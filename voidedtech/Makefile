DESTDIR := bin/
WEBROOT := www/

build:
	mkdir -p $(WEBROOT)
	cp *.css $(WEBROOT)
	cp *.html $(WEBROOT)
	sed -i "s/{DATE}/$(shell date +%Y-%m-%d)/g" $(WEBROOT)index.html

install:
	mkdir -p $(DESTDIR)
	rsync -iacv $(WEBROOT) $(DESTDIR)
