DIST := dist/

release: clean
	./www/configure "$(DIST)"

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)
