package main

import (
	_ "github.com/sfomuseum/go-www-geotag-whosonfirst/writer"
	_ "github.com/whosonfirst/go-reader-github"
	_ "github.com/whosonfirst/go-writer-github"
)

import (
	"context"
	"github.com/sfomuseum/go-flags"
	sfom_app "github.com/sfomuseum/go-www-geotag-sfomuseum/app"
	geotag_app "github.com/sfomuseum/go-www-geotag/app"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()

	fl, err := geotag_app.CommonFlags()

	if err != nil {
		log.Fatalf("Failed to instantiate common flags, %v", err)
	}

	err = sfom_app.AppendSFOMuseumFlags(fl)

	if err != nil {
		log.Fatalf("Failed to append SFO Museum flags, %v", err)
	}

	flags.Parse(fl)

	err = flags.SetFlagsFromEnvVarsWithFeedback(fl, "GEOTAG", true)

	if err != nil {
		log.Fatalf("Failed to set flags from env vars, %v", err)
	}

	sfom_app.AssignSFOMuseumFlags(fl)

	mux := http.NewServeMux()

	err = sfom_app.AppendAssetHandlers(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append asset handlers, %v", err)
	}

	err = sfom_app.AppendEditorHandler(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append editor handler, %v", err)
	}

	err = geotag_app.AppendProxyTilesHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append proxy tiles handler, %v", err)
	}

	err = sfom_app.AppendWriterHandler(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}

	err = sfom_app.AppendOAuth2HandlersIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}

	s, err := geotag_app.NewServer(ctx, fl)

	if err != nil {
		log.Fatalf("Failed to create application server, %v", err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
