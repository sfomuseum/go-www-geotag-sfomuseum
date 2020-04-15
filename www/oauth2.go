package www

import (
	"bytes"
	"encoding/json"
	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-string/random"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type OAuth2AccessResponse struct {
	AccessToken string `json:"access_token"`
}

type OAuth2Options struct {
	ClientId     string
	ClientSecret string
	Scope        string
	RedirectURL  string
	State        string
	Endpoint     oauth2.Endpoint
}

func OAuth2AuthorizeHandler(opts *OAuth2Options) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		rand_opts := random.DefaultOptions()
		state, err := random.String(rand_opts)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		redir := url.URL{}
		redir.Scheme = req.URL.Scheme
		redir.Host = req.URL.Host
		redir.Path = opts.RedirectURL

		redir_params := url.Values{}
		redir_params.Set("state", state)

		redir.RawQuery = redir_params.Encode()
		redir_url := redir.String()

		params := url.Values{}
		params.Set("client_id", opts.ClientId)
		params.Set("client_secret", opts.ClientSecret)
		params.Set("scope", opts.Scope)
		params.Set("state", state)
		params.Set("redirect_uri", redir_url)

		auth_req, err := http.NewRequest(http.MethodPost, opts.Endpoint.AuthURL, nil)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		auth_req.URL.RawQuery = params.Encode()
		auth_url := auth_req.URL.String()

		http.Redirect(rsp, req, auth_url, http.StatusSeeOther)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}

func OAuth2AccessTokenHandler(opts *OAuth2Options) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		code, err := sanitize.PostString(req, "code")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		state, err := sanitize.PostString(req, "state")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		params := url.Values{}
		params.Set("client_id", opts.ClientId)
		params.Set("client_secret", opts.ClientSecret)
		params.Set("state", state)
		params.Set("code", code)

		post_body := bytes.NewBufferString(params.Encode())

		auth_req, err := http.NewRequest(http.MethodPost, opts.Endpoint.TokenURL, post_body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		auth_req.Header.Set("accept", "application/json")

		client := http.Client{}
		auth_res, err := client.Do(auth_req)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		defer auth_res.Body.Close()

		var access OAuth2AccessResponse

		dec := json.NewDecoder(auth_res.Body)
		err = dec.Decode(&access)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
