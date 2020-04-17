package oauth2

import (
	"github.com/aaronland/go-http-crumb"
	"golang.org/x/oauth2"
)

type Options struct {
	AuthURL     string
	CookieURI   string
	Config      *oauth2.Config
	AuthCrumb   crumb.Crumb
	UnAuthCrumb crumb.Crumb
}
