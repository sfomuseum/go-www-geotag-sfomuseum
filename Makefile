CWD=$(shell pwd)

go-bindata:
	mkdir -p cmd/go-bindata
	mkdir -p cmd/go-bindata-assetfs
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/master/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/master/cmd/go-bindata-assetfs/main.go

bake:
	@make bake-assets

bake-templates:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

bake-assets:	
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	rm -f static/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg www -o www/assets.go static static/javascript

debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' -enable-writer -writer-uri 'whosonfirst://?writer=$(WRITER)&reader=$(READER)&update=1&source=sfomuseum' -crumb-dsn debug

debug-gh:
	go run -mod vendor cmd/server/main.go -server-uri 'mkcert://localhost:8080' -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' -enable-writer -writer-uri 'whosonfirst://?writer=$(WRITER)&reader=$(READER)&update=1&source=sfomuseum' -crumb-dsn debug -enable-oauth2 -oauth2-client-id $(CLIENTID) -oauth2-client-secret $(CLIENTSECRET) -oauth2-cookie-uri debug -oauth2-scopes 'user,repo'
