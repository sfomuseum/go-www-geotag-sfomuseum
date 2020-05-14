package api

import (
	"fmt"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-geojson-geotag"
	oauth2_www "github.com/sfomuseum/go-http-oauth2/www"
	"github.com/sfomuseum/go-www-geotag/writer"
	_ "log"
	"net/http"
	"net/url"
	"strings"
)

func WriterHandler(wr_uri string) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "PUT":
			// pass
		default:
			http.Error(rsp, "Method not allowed.", http.StatusMethodNotAllowed)
			return
		}

		defer req.Body.Close()

		ctx := req.Context()

		token, err := oauth2_www.GetOAuth2TokenContext(req)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		wr_u, err := url.Parse(wr_uri)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if token != nil {

			switch wr_u.Scheme {
			case "whosonfirst":

				wr_q := wr_u.Query()

				wr_reader := wr_q.Get("reader")
				wr_writer := wr_q.Get("writer")

				wr_reader, err = url.QueryUnescape(wr_reader)

				if err != nil {
					http.Error(rsp, err.Error(), http.StatusInternalServerError)
					return
				}

				wr_writer, err = url.QueryUnescape(wr_writer)

				if err != nil {
					http.Error(rsp, err.Error(), http.StatusInternalServerError)
					return
				}

				wr_reader = strings.Replace(wr_reader, "{access_token}", token.AccessToken, -1)
				wr_writer = strings.Replace(wr_writer, "{access_token}", token.AccessToken, -1)

				wr_reader = url.QueryEscape(wr_reader)
				wr_writer = url.QueryEscape(wr_writer)

				wr_q.Set("reader", wr_reader)
				wr_q.Set("writer", wr_writer)

				wr_uri = fmt.Sprintf("whosonfirst://?%s", wr_q.Encode())

			default:
				// pass
			}
		}

		wr, err := writer.NewWriter(ctx, wr_uri)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		uid, err := sanitize.GetString(req, "id")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		geotag_f, err := geotag.NewGeotagFeatureWithReader(req.Body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, err = writer.SetIOWriterWithContext(ctx, rsp)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")

		err = wr.WriteFeature(ctx, uid, geotag_f)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
