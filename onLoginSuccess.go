package gogauthurls

import (
	"github.com/arelate/gogurls"
	"net/url"
)

func OnLoginSuccess() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   gogurls.WwwGogHost,
		Path:   onLoginSuccessPath,
	}
}
