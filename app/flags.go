package app

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-flags"
	oauth2_flags "github.com/sfomuseum/go-http-oauth2/flags"
	wof_app "github.com/sfomuseum/go-www-geotag-whosonfirst/app"
	"gocloud.dev/runtimevar"
	_ "log"
)

func AppendSFOMuseumFlags(fs *flag.FlagSet) error {

	err := oauth2_flags.AppendOAuth2Flags(fs)

	if err != nil {
		return err
	}

	err = wof_app.AppendWhosOnFirstFlags(fs)

	if err != nil {
		return err
	}

	fs.Bool("enable-wk-webview", false, "Enable support to operate in a WKWebView context.")
	fs.String("wk-webview-handler", "", "The name of the WKScriptMessageHandler to dispatch geotag updates to.")

	return nil
}

func AssignSFOMuseumFlags(fs *flag.FlagSet) error {

	fs.Set("enable-map-layers", "true")

	fs.Set("oauth2-auth-url", "https://github.com/login/oauth/authorize")
	fs.Set("oauth2-token-url", "https://github.com/login/oauth/access_token")

	err := wof_app.AssignWhosOnFirstFlags(fs)

	if err != nil {
		return err
	}

	enable_oauth2, err := flags.BoolVar(fs, "enable-oauth2")

	if err != nil {
		return err
	}

	if enable_oauth2 {

		id_uri, err := flags.StringVar(fs, "oauth2-client-id")

		if err != nil {
			return err
		}

		secret_uri, err := flags.StringVar(fs, "oauth2-client-secret")

		if err != nil {
			return err
		}

		cookie_uri, err := flags.StringVar(fs, "oauth2-cookie-uri")

		if err != nil {
			return err
		}

		ctx := context.Background()

		client_id, err := runtimeStringVar(ctx, id_uri)

		if err != nil {
			return err
		}

		client_secret, err := runtimeStringVar(ctx, secret_uri)

		if err != nil {
			return err
		}

		var oauth2_cookie string

		if cookie_uri == "debug" {

			oauth2_cookie = "constant://?val=debug&decoder=string"

		} else {

			cookie, err := runtimeStringVar(ctx, cookie_uri)

			if err != nil {
				return err
			}

			oauth2_cookie = cookie
		}

		fs.Set("oauth2-client-id", client_id)
		fs.Set("oauth2-client-secret", client_secret)
		fs.Set("oauth2-cookie-uri", oauth2_cookie)
	}

	return nil
}

func runtimeStringVar(ctx context.Context, uri string) (string, error) {

	v, err := runtimevar.OpenVariable(ctx, uri)

	if err != nil {
		return "", err
	}

	latest, err := v.Latest(ctx)

	if err != nil {
		return "", err
	}

	return latest.Value.(string), nil
}
