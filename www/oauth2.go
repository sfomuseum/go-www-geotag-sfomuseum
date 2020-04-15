package www

import (
	"github.com/aaronland/go-http-sanitize"
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
		
		scheme := "http"

		// because this: https://github.com/golang/go/issues/28940#issuecomment-441749380
		
		if req.TLS != nil {
			scheme = "https"
		}
		
		redir := url.URL{}
		redir.Scheme = scheme
		redir.Host = req.Host
		redir.Path = cfg.RedirectURL

		redir_url := redir.String()
		cfg.RedirectURL = redir_url

		auth_url := cfg.AuthCodeURL("state", oauth2.AccessTypeOnline)
		
		http.Redirect(rsp, req, auth_url, http.StatusSeeOther)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func OAuth2AccessTokenHandler(cfg *oauth2.Config) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		code, err := sanitize.GetString(req, "code")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		if code == "" {
			http.Error(rsp, "Missing code", http.StatusBadRequest)
			return
		}
		
		ctx := req.Context()

		tok, err := cfg.Exchange(ctx, code, oauth2.AccessTypeOnline)

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
