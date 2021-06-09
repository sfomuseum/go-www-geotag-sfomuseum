debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' -enable-writer -writer-uri 'whosonfirst://?writer=$(WRITER)&reader=$(READER)&update=1&source=sfomuseum'

debug-gh:
	go run -mod vendor cmd/server/main.go -server-uri 'mkcert://localhost:8080' -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' -enable-writer -writer-uri 'whosonfirst://?writer=$(WRITER)&reader=$(READER)&update=1&source=sfomuseum' -enable-oauth2 -oauth2-client-id $(CLIENTID) -oauth2-client-secret $(CLIENTSECRET) -oauth2-cookie-uri debug -oauth2-scopes 'user,repo'
