DIST := dist/
SUBS := notebook

.PHONY: $(SUBS)

release: clean main subsites

clean:
	rm -rf $(DIST)
	mkdir -p $(DIST)

main:
	./www/configure "$(DIST)"

subsites: $(SUBS)

$(SUBS):
	cd $@ && ./configure && mdbook build
	cp -r $@/book $(DIST)$@
