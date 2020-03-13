CDIST=client/dist
SVG_DIR=client/src/img
PNG_DIR=client/src/img/intermediates
PNGS= $(patsubst $(SVG_DIR)/%.svg, $(PNG_DIR)/%.png,$(wildcard $(SVG_DIR)/*.svg))
PWD= $(shell pwd)

all: client

robo: #TODO
	go build .

.PHONY: all client $(CDIST)/bundle.js clean png docker-inkscape-image

client: $(CDIST)/index.html $(CDIST)/bundle.js
$(CDIST)/index.html: client/src/index.html | $(CDIST)
	cp $< $@

$(CDIST):
	mkdir $(CDIST)

$(CDIST)/bundle.js: png
	(cd client && npm run build)

client/node_modules: client/package.json
	(cd client && npm install)

png: $(PNGS)

$(PNG_DIR)/%.png: $(SVG_DIR)/%.svg | $(PNG_DIR) docker-inkscape-image
	docker run -it --mount "type=bind,src=$(PWD)/client,target=/home/client" --rm tjbearse/robo/inkscape -z $< -e $@

docker-inkscape-image:
	docker build -t tjbearse/robo/inkscape -f inkscape.Dockerfile .

$(PNG_DIR):
	mkdir $(PNG_DIR)

clean:
	rm -r $(PNG_DIR)
	rm -r $(CDIST)
