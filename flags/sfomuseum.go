package flags

import (
	"flag"
)

func AppendSFOMuseumFlags(fs *flag.FlagSet) error {

	fs.Bool("enable-oauth2", false, "...")

	fs.String("oauth2-client-id", "", "...")
	fs.String("oauth2-client-secret", "", "...")

	fs.String("oauth2-scopes", "", "...")

	fs.String("path-oauth2-auth", "/signin", "...")
	fs.String("path-oauth2-token", "/auth", "...")

	fs.String("oauth2-auth-url", "https://github.com/login/oauth/authorize", "...")
	fs.String("oauth2-token-url", "https://github.com/login/oauth/access_token", "...")

	return nil
}
