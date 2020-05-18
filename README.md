# go-www-geotag-sfomuseum

A web application, written in Go, for geotagging images in the SFO Museum collection.

## Important

This is work in progress, including the documentation. In the meantime please have a look at the [Geotagging at SFO Museum](https://millsfield.sfomuseum.org/blog/tags/geotagging) series of blog posts and the [Geotagging at SFO Museum, Part 7 – Custom Writers](https://millsfield.sfomuseum.org/blog/2020/05/01/geotagging-custom-writers/) and [Geotagging at SFO Museum, part 9 – Publishing Data](https://millsfield.sfomuseum.org/blog/2020/05/07/geotagging-publishing/) and [Geotagging at SFO Museum, part 10 – Native Applications](https://millsfield.sfomuseum.org/blog/2020/05/18/geotagging-native/) posts in particular.

## Example

```
$> cd go-www-geotag-sfomuseum
$> go build -mod vendor -p bin/server cmd/server/main.go

$> bin/server \
	-nextzen-apikey {NEXTZEN_API_KEY} \
	-enable-placeholder
	-placeholder-endpoint {PLACEHOLDER_API_KEY} \
	-enable-oembed \
	-oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}' \
	-enable-writer \
	-writer-uri 'whosonfirst://?writer={whosonfirst_writer}&reader={whosonfirst_reader}&update=1&source=sfomuseum' \
	-whosonfirst-writer-uri 'githubapi://sfomuseum-data/sfomuseum-data-collection?access_token={access_token}&prefix=data/' \
	-whosonfirst-reader-uri 'githubapi://sfomuseum-data/sfomuseum-data-collection?access_token={access_token}&prefix=data/' \
	-enable-oauth2 \
	-oauth2-scopes 'user,repo' \
	-oauth2-client-id "constant://?val={OAUTH2_CLIENT_ID}&decoder=string" \
	-oauth2-client-secret "constant://?val={OAUTH2_SECRET}&decoder=string" \
	-oauth2-cookie-uri "constant://?val=debug&decoder=string" \
	-server-uri 'mkcert://localhost:8080'
	
2020/05/05 11:42:37 Checking whether mkcert is installed. If it is not you may be prompted for your password (in order to install certificate files)
2020/05/05 11:42:40 Listening on https://localhost:8080
```

## See also

* https://github.com/sfomuseum/go-www-geotag
* https://github.com/sfomuseum/go-www-geotag-whosonfirst
* https://github.com/sfomuseum/go-http-oauth2
