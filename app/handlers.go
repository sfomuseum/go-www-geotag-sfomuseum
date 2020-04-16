package app

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-flags"
	oauth2_flags "github.com/sfomuseum/go-http-oauth2/flags"
	oauth2_www "github.com/sfomuseum/go-http-oauth2/www"
	sfom_api "github.com/sfomuseum/go-www-geotag-sfomuseum/api"
	sfom_www "github.com/sfomuseum/go-www-geotag-sfomuseum/www"
	geotag_app "github.com/sfomuseum/go-www-geotag/app"
	_ "log"
	"net/http"
)

func AppendAssetHandlers(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	err := geotag_app.AppendAssetHandlers(ctx, fs, mux)

	if err != nil {
		return err
	}

	err = sfom_www.AppendAssetHandlers(mux)

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

	editor_opts := sfom_www.DefaultEditorOptions()
	handler = sfom_www.AppendResourcesHandler(handler, editor_opts)

	opts, err := oauth2_flags.OAuth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	handler = oauth2_www.EnsureOAuth2TokenHandler(opts, handler)

	mux.Handle(path, handler)
	return nil
}

func AppendWriterHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path, err := flags.StringVar(fs, "path-writer")

	if err != nil {
		return err
	}

	handler, err := NewWriterHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path, handler)
	return nil
}

func NewWriterHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	writer_uri, err := flags.StringVar(fs, "writer-uri")

	if err != nil {
		return nil, err
	}

	handler, err := sfom_api.WriterHandler(writer_uri)

	if err != nil {
		return nil, err
	}

	handler, err = geotag_app.AppendCrumbHandler(ctx, fs, handler)

	if err != nil {
		return nil, err
	}

	opts, err := oauth2_flags.OAuth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	handler = oauth2_www.EnsureOAuth2TokenHandler(opts, handler)

	return handler, nil
}

func AppendOAuth2HandlersIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enabled, err := flags.BoolVar(fs, "enable-oauth2")

	if err != nil {
		return err
	}

	if !enabled {
		return nil
	}

	path_auth, err := flags.StringVar(fs, "path-oauth2-auth")

	if err != nil {
		return err
	}

	path_token, err := flags.StringVar(fs, "path-oauth2-token")

	if err != nil {
		return err
	}

	auth_handler, err := NewOAuth2AuthorizeHandler(ctx, fs)

	if err != nil {
		return err
	}

	token_handler, err := NewOAuth2TokenHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path_auth, auth_handler)
	mux.Handle(path_token, token_handler)
	return nil
}

func NewOAuth2AuthorizeHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	opts, err := oauth2_flags.OAuth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	return oauth2_www.OAuth2AuthorizeHandler(opts)
}

func NewOAuth2TokenHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	opts, err := oauth2_flags.OAuth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	return oauth2_www.OAuth2AccessTokenHandler(opts)
}
