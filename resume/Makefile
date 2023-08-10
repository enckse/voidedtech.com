NAME      :=
TARGET    :=
BIN       := dist/
TYPES     := $(BIN)personal.html $(BIN)work.html
VARS      := $(HOME)/.env/local/misc/resume.js
NORMALIZE := 8.0.1
AWESOME   := 6.4.0
FAWORKING := fa.zip

resumes: resources $(TYPES)

setup:
	mkdir -p $(BIN)

resources: setup
	@cp $(VARS) $(BIN)resume.js
	@cp src/*.css $(BIN)
	@cp -r vendor $(BIN)

$(TYPES): src/*.in
	make _resume NAME=$(shell basename $@ | cut -d "." -f 1) TARGET=$@

_resume:
	cat src/resume.html.in | sed "s/{MODE}/$(NAME)/g" > $(TARGET) 

clean:
	rm -rf $(BIN)

deps: clean setup
	wget -O vendor/normalize.css/normalize.css "https://necolas.github.io/normalize.css/$(NORMALIZE)/normalize.css"
	wget -O $(BIN)$(FAWORKING) "https://use.fontawesome.com/releases/v$(AWESOME)/fontawesome-free-$(AWESOME)-web.zip"
	cd $(BIN) && unzip $(FAWORKING)
	cp -r $(BIN)fontawesome-free-$(AWESOME)-web/* vendor/fontawesome/

public:
	make resumes VARS=$(PWD)/public.js TYPES=$(BIN)index.html
