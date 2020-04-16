package www

import (
	"context"
	"encoding/json"
	"github.com/aaronland/go-http-cookie"
	"github.com/aaronland/go-http-sanitize"
	"golang.org/x/oauth2"
	"log"
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

		token, err := GetTokenFromCookie(opts, req)

		if err != nil && err != http.ErrNoCookie {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Redirect(rsp, req, opts.Config.RedirectURL, 303)
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

		token, err := GetTokenFromCookie(opts, req)

		if err != nil && err != http.ErrNoCookie {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if token != nil {
			http.Redirect(rsp, req, "/", 303)
			return
		}

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

		log.Println(redir_url)

		auth_url := cfg.AuthCodeURL("state", oauth2.AccessTypeOnline)

		log.Println(auth_url)
		http.Redirect(rsp, req, auth_url, http.StatusSeeOther)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func OAuth2AccessTokenHandler(opts *OAuth2Options) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		cfg := opts.Config

		code, err := RequestString(req, "code")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := req.Context()

		tok, err := cfg.Exchange(ctx, code, oauth2.AccessTypeOnline)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		err = SetCookieWithToken(opts, rsp, tok)

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

func GetTokenFromCookie(opts *OAuth2Options, req *http.Request) (*oauth2.Token, error) {

	ck, err := cookie.NewAuthCookie(opts.CookieName, opts.CookieSecret, opts.CookieSalt)

	if err != nil {
		return nil, err
	}

	str_token, err := ck.Get(req)

	if err != nil && err != http.ErrNoCookie {
		return nil, err
	}

	if str_token == "" {
		return nil, http.ErrNoCookie
	}

	var token *oauth2.Token

	err = json.Unmarshal([]byte(str_token), &token)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func SetCookieWithToken(opts *OAuth2Options, rsp http.ResponseWriter, tok *oauth2.Token) error {

	ck, err := cookie.NewAuthCookie(opts.CookieName, opts.CookieSecret, opts.CookieSalt)

	if err != nil {
		return err
	}

	enc_token, err := json.Marshal(tok)

	if err != nil {
		return err
	}

	str_token := string(enc_token)

	// https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00#section-4.1.1

	http_cookie := &http.Cookie{
		Value:    str_token,
		SameSite: http.SameSiteLaxMode,
		// SameSite: http.SameSiteDefaultMode,
		Expires: tok.Expiry,
		Path:    "/",
	}

	return ck.SetCookie(rsp, http_cookie)
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

func RequestString(req *http.Request, param string) (string, error) {

	value, err := sanitize.PostString(req, param)

	if err != nil {
		return "", err
	}

	if value == "" {

		value, err = sanitize.GetString(req, param)

		if err != nil {
			return "", err
		}
	}

	return value, nil
}
