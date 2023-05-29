DIST := dist/

release: clean
	./www/configure "$(DIST)"
	cd notebook && ./configure && mdbook build
	cp -r notebook/book $(DIST)notebook

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)
