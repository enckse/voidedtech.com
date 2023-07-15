DIST := dist/

release: clean
	./www/configure "$(DIST)"
	cd notebook && ./configure && mdbook build
	cp -r notebook/book $(DIST)notebook
	cd resume && make public
	cp -r resume/dist $(DIST)resume

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)
