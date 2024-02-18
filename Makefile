DIST := dist/
DATE := $(shell date +%Y-%m-%d)

release: clean
	cat index.html.in | sed 's/{DATE}/$(DATE)/g' > $(DIST)index.html

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)
