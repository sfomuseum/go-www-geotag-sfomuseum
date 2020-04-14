package app

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-www-geotag-sfomuseum"
	geotag_app "github.com/sfomuseum/go-www-geotag/app"
	"github.com/sfomuseum/go-www-geotag/flags"
	"net/http"
)

func AppendAssetHandlers(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	err := geotag_app.AppendAssetHandlers(ctx, fs, mux)

	if err != nil {
		return err
	}

	err = sfomuseum.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	return nil
}

func AppendEditorHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path, err := flags.StringVar(fs, "path-editor")

	if err != nil {
		return err
	}

	handler, err := geotag_app.NewEditorHandler(ctx, fs)

	if err != nil {
		return err
	}

	editor_opts := sfomuseum.DefaultEditorOptions()
	handler = sfomuseum.AppendResourcesHandler(handler, editor_opts)

	handler, err = geotag_app.AppendCrumbHandler(ctx, fs, handler)

	if err != nil {
		return err
	}

	mux.Handle(path, handler)
	return nil
}
