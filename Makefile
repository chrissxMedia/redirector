all: response.min.html

response.min.html: response.html
	npx html-minifier --collapse-whitespace --remove-comments --remove-tag-whitespace --minify-css true --minify-js true -o $@ $<

clean:
	rm -f response.min.html

.PHONY: all clean
