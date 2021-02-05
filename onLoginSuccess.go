package gogauthurls

import (
	"github.com/boggydigital/gogog/pkg/gogurls"
	"net/url"
)

func OnLoginSuccess() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   gogurls.WwwGogHost,
		Path:   onLoginSuccessPath,
	}
}
