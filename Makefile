BUILD_ARTIFACTS = ".build-artifiacts"
BINARY_NAME = "simple-webapp"

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

generate-assets: ## generates static assets
	statics -i=assets/templates -o=assets/templates.go  -pkg=assets -group=Templates 	-ignore=\.gitignore -prefix=assets
	statics -i=assets/css       -o=assets/css.go        -pkg=assets -group=Styles 		-ignore=\.gitignore -prefix=assets
	statics -i=assets/icons     -o=assets/icons.go      -pkg=assets -group=Icons 		-ignore=\.gitignore -prefix=assets
	statics -i=assets/js 	    -o=assets/javascript.go -pkg=assets -group=Javascript 	-ignore=\.gitignore -prefix=assets
	
build: generate-assets ##
	go build -o $(BUILD_ARTIFACTS)/$(BINARY_NAME) cmd/main.go

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help