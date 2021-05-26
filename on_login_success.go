package gog_auth_urls

import (
	"github.com/arelate/gog_urls"
	"net/url"
)

func OnLoginSuccess() *url.URL {
	return &url.URL{
		Scheme: gog_urls.HttpsScheme,
		Host:   gog_urls.WwwGogHost,
		Path:   onLoginSuccessPath,
	}
}
