package main

import (
	_ "github.com/sfomuseum/go-www-geotag-whosonfirst/writer"
	_ "github.com/whosonfirst/go-reader-github"
	_ "github.com/whosonfirst/go-writer-github"
)

import (
	"context"
	sfom_app "github.com/sfomuseum/go-www-geotag-sfomuseum/app"
	sfom_flags "github.com/sfomuseum/go-www-geotag-sfomuseum/flags"	
	"github.com/sfomuseum/go-www-geotag/app"
	"github.com/sfomuseum/go-www-geotag/flags"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()

	fl, err := flags.CommonFlags()

	if err != nil {
		log.Fatalf("Failed to instantiate common flags, %v", err)
	}

	err = sfom_flags.AppendSFOMuseumFlags(fl)

	if err != nil {
		log.Fatalf("Failed to append SFO Museum flags, %v", err)
	}
	
	flags.Parse(fl)
	fl.Set("enable-map-layers", "true")

	mux := http.NewServeMux()

	err = sfom_app.AppendAssetHandlers(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append asset handlers, %v", err)
	}

	err = sfom_app.AppendEditorHandler(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append editor handler, %v", err)
	}

	err = app.AppendProxyTilesHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append proxy tiles handler, %v", err)
	}

	err = app.AppendWriterHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}

	err = sfom_app.AppendOAuth2HandlersIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}
	
	s, err := app.NewServer(ctx, fl)

	if err != nil {
		log.Fatalf("Failed to create application server, %v", err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
