package gog_auth_urls

import (
	"github.com/arelate/gog_urls"
	"net/url"
)

func ReCaptcha() *url.URL {
	return &url.URL{
		Scheme: gog_urls.HttpsScheme,
		Host:   reCaptchaHost,
		Path:   reCaptchaPath}
}
