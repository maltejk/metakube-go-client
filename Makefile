# ====================================================================================
# Setup Project

all: generate

generate:
	rm -rf README.md pkg/client pkg/models
	swagger generate client -f doc/metakube-api-docs.json --skip-validation -c pkg/client -m pkg/models --default-scheme=https
	swagger generate markdown -f doc/metakube-api-docs.json --skip-validation --output readme.md

fetch:
	mkdir doc || true
	curl -o doc/metakube-api-docs.json https://metakube.syseleven.de/api/swagger.json

