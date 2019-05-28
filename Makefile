CDIST=client/dist
SVG_DIR=client/src/img
PNG_DIR=client/src/img/intermediates
PNGS= $(patsubst $(SVG_DIR)/%.svg, $(PNG_DIR)/%.png,$(wildcard $(SVG_DIR)/*.svg))

all: client

robo: #TODO
	go build .

.PHONY: all client $(CDIST)/bundle.js clean png

client: $(CDIST)/index.html $(CDIST)/bundle.js
$(CDIST)/index.html: client/src/index.html $(CDIST)
	cp $< $@
$(CDIST):
	mkdir $(CDIST)

$(CDIST)/bundle.js: png
	(cd client && npm run build)

client/node_modules: client/package.json
	(cd client && npm install)

png: $(PNGS)

$(PNG_DIR)/%.png: $(SVG_DIR)/%.svg $(PNG_DIR)
	inkscape -z $< -e $@

$(PNG_DIR):
	mkdir $(PNG_DIR)

clean:
	rm -r $(PNG_DIR)
	rm -r $(CDIST)
