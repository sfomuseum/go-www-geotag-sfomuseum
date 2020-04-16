package www

import (
	"context"
	"encoding/json"
	"github.com/aaronland/go-http-cookie"
	"github.com/aaronland/go-http-sanitize"
	"golang.org/x/oauth2"
	_ "log"
	"net/http"
	"net/url"
)

const CONTEXT_TOKEN_KEY string = "token"

type OAuth2Options struct {
	SigninURL    string
	CookieName   string
	CookieSecret string
	CookieSalt   string
	Config       *oauth2.Config
}

func EnsureOAuth2TokenHandler(opts *OAuth2Options, next http.Handler) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ck, err := cookie.NewAuthCookie(opts.CookieName, opts.CookieSecret, opts.CookieSalt)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		str_token, err := ck.Get(req)

		if err != nil && err != http.ErrNoCookie {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		// set internal redirect URL here...

		if err != nil {
			http.Redirect(rsp, req, opts.Config.RedirectURL, 303)
			return
		}

		if str_token == "" {
			http.Redirect(rsp, req, opts.Config.RedirectURL, 303)
			return
		}

		var token *oauth2.Token

		err = json.Unmarshal([]byte(str_token), &token)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		req, err = SetTokenContext(req, token)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(rsp, req)
	}

	h := http.HandlerFunc(fn)
	return h
}

func OAuth2AuthorizeHandler(opts *OAuth2Options) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		cfg := opts.Config

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

func OAuth2AccessTokenHandler(opts *OAuth2Options) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		cfg := opts.Config
		code, err := sanitize.PostString(req, "code")

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

		ck, err := cookie.NewAuthCookie(opts.CookieName, opts.CookieSecret, opts.CookieSalt)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		enc_token, err := json.Marshal(tok)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		str_token := string(enc_token)

		http_cookie := &http.Cookie{
			Value:    str_token,
			SameSite: http.SameSiteStrictMode,
			Expires:  tok.Expiry,
		}

		err = ck.SetCookie(rsp, http_cookie)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		// get internal redirect URL here... (see above)

		redir_url := "/"

		http.Redirect(rsp, req, redir_url, 303)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func SetTokenContext(req *http.Request, token *oauth2.Token) (*http.Request, error) {

	ctx := req.Context()
	ctx = context.WithValue(ctx, CONTEXT_TOKEN_KEY, token)

	return req.WithContext(ctx), nil
}

func GetTokenContext(req *http.Request) (*oauth2.Token, error) {

	ctx := req.Context()
	v := ctx.Value(CONTEXT_TOKEN_KEY)

	if v == nil {
		return nil, nil
	}

	token := v.(*oauth2.Token)
	return token, nil
}
