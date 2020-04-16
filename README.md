# go-www-geotag-sfomuseum

Work in progress.

## Example

```
go run -mod vendor cmd/server/main.go \
	-server-uri 'mkcert://localhost:8080' \
	-nextzen-apikey {NEXTZEN_API_KEY} \
	-enable-placeholder \
	-placeholder-endpoint {PLACEHOLDER_ENDPOINT} \
	-enable-oembed \
	-oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' \
	-enable-writer \
	-writer-uri 'whosonfirst://?writer=githubapi%3A%2F%2Fsfomuseum-data%2Fsfomuseum-data-collection%3Faccess_token%3D%7Baccess_token%7D%26prefix%3Ddata&reader=githubapi%3A%2F%2Fsfomuseum-data%2Fsfomuseum-data-collection%3Faccess_token%3D%7Baccess_token%7D%26prefix%3Ddata&update=1&source=sfomuseum' \
	-crumb-dsn debug \
	-enable-oauth2 \
	-oauth2-client-id {CLIENT_ID} \
	-oauth2-client-secret {CLIENT_SECRET} \
	-oauth2-cookie-dsn debug \
	-oauth2-scopes 'user,repo'
2020/04/16 16:03:57 Listening on https://localhost:8080
```

## See also

* https://github.com/sfomuseum/go-www-geotag
* https://github.com/sfomuseum/go-www-geotag-whosonfirst
