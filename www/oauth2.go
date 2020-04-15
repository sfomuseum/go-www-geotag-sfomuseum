package www

import (
	"github.com/aaronland/go-http-sanitize"
	// "github.com/aaronland/go-string/random"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
)

type OAuth2AccessResponse struct {
	AccessToken string `json:"access_token"`
}

func OAuth2AuthorizeHandler(cfg *oauth2.Config) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		/*
			rand_opts := random.DefaultOptions()
			state, err := random.String(rand_opts)

			if err != nil {
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}
		*/

		redir := url.URL{}
		redir.Scheme = req.URL.Scheme
		redir.Host = req.URL.Host
		redir.Path = cfg.RedirectURL

		cfg.RedirectURL = redir.String()

		auth_url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
		// auth_url := cfg.AuthCodeURL("state", state)

		http.Redirect(rsp, req, auth_url, http.StatusSeeOther)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func OAuth2AccessTokenHandler(cfg *oauth2.Config) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		code, err := sanitize.PostString(req, "code")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := req.Context()

		tok, err := cfg.Exchange(ctx, code)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(tok)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
