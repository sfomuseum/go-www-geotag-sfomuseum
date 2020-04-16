package app

import (
	"flag"
	oauth2_flags "github.com/sfomuseum/go-http-oauth2/flags"
)

func AppendSFOMuseumFlags(fs *flag.FlagSet) error {

	return oauth2_flags.AppendOAuth2Flags(fs)
}

func AssignSFOMuseumFlags(fs *flag.FlagSet) {

	fs.Set("enable-map-layers", "true")

	fs.Set("oauth2-auth-url", "https://github.com/login/oauth/authorize")
	fs.Set("oauth2-token-url", "https://github.com/login/oauth/access_token")
}
