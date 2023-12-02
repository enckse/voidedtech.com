DIST := dist/

release: clean
	./www/configure "$(DIST)"
	cd resume && make public
	cp -r resume/dist $(DIST)resume

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)
