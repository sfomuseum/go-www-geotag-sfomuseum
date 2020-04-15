package app

import (
	"context"
	"errors"
	"flag"
	"github.com/sfomuseum/go-www-geotag-sfomuseum"
	"github.com/sfomuseum/go-www-geotag-sfomuseum/www"
	geotag_app "github.com/sfomuseum/go-www-geotag/app"
	"github.com/sfomuseum/go-www-geotag/flags"
	"golang.org/x/oauth2"
	_ "log"
	"net/http"
	"strings"
	"sync"
)

var oauth2_init sync.Once

var oauth2_opts *www.OAuth2Options
var oauth2_err error

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

	opts, err := oauth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	handler = www.EnsureOAuth2TokenHandler(opts, handler)

	mux.Handle(path, handler)
	return nil
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

	opts, err := oauth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	return www.OAuth2AuthorizeHandler(opts)
}

func NewOAuth2TokenHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	opts, err := oauth2OptionsWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	return www.OAuth2AccessTokenHandler(opts)
}

func oauth2OptionsWithFlagSet(ctx context.Context, fs *flag.FlagSet) (*www.OAuth2Options, error) {

	oauth2_func := func() {

		client_id, err := flags.StringVar(fs, "oauth2-client-id")

		if err != nil {
			oauth2_err = err
			return
		}

		client_secret, err := flags.StringVar(fs, "oauth2-client-secret")

		if err != nil {
			oauth2_err = err
			return
		}

		auth_url, err := flags.StringVar(fs, "oauth2-auth-url")

		if err != nil {
			oauth2_err = err
			return
		}

		token_url, err := flags.StringVar(fs, "oauth2-token-url")

		if err != nil {
			oauth2_err = err
			return
		}

		str_scopes, err := flags.StringVar(fs, "oauth2-scopes")

		if err != nil {
			oauth2_err = err
			return
		}

		scopes := strings.Split(str_scopes, ",")

		path_token, err := flags.StringVar(fs, "path-oauth2-token")

		if err != nil {
			oauth2_err = err
			return
		}

		oauth2_cfg := &oauth2.Config{
			ClientID:     client_id,
			ClientSecret: client_secret,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  auth_url,
				TokenURL: token_url,
			},
			RedirectURL: path_token,
		}

		oauth2_opts = &www.OAuth2Options{
			Config: oauth2_cfg,
			// cookie stuff goes here...
		}
	}

	oauth2_init.Do(oauth2_func)

	if oauth2_err != nil {
		return nil, oauth2_err
	}

	if oauth2_opts == nil {
		return nil, errors.New("Failed to construct OAuth2 options")
	}

	return oauth2_opts, nil
}
