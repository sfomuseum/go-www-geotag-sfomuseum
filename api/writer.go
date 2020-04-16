package api

import (
	"github.com/sfomuseum/go-geojson-geotag"
	oauth2_www "github.com/sfomuseum/go-http-oauth2/www"
	"github.com/sfomuseum/go-www-geotag/writer"
	"net/http"
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

		token, err := oauth2_www.GetTokenContext(req)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		wr_uri = strings.Replace(wr_uri, "{access_token}", token.AccessToken, -1)

		wr, err := writer.NewWriter(ctx, wr_uri)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		geotag_f, err := geotag.NewGeotagFeatureWithReader(req.Body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		uri := geotag_f.Id

		if uri == "" {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		err = wr.WriteFeature(ctx, uri, geotag_f)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
