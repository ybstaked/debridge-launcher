swaggerui_repo    := https://github.com/swagger-api/swagger-ui
swaggerui_version := 3.24.0

spec/swaggerui:
	mkdir $@
	curl -Ls $(swaggerui_repo)/archive/v$(swaggerui_version).tar.gz \
		| tar -C $@ -xz --transform "s|^swagger-ui-$(swaggerui_version)|.|"

spec/swaggerui.go: spec/swaggerui
	nix-shell -p statik --run "statik --src $</dist --dest $(PWD) -p $(patsubst %/,%,$(dir $@)) -m"
	mv -f $(dir $@)statik.go $@

.PHONY: clean
clean:
	rm -rf spec/swaggerui.go
	rm -rf spec/swaggerui
