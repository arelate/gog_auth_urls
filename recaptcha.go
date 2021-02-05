package gogauthurls

import (
	"github.com/boggydigital/gogog/pkg/gogurls"
	"net/url"
)

func ReCaptcha() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   reCaptchaHost,
		Path:   reCaptchaPath}
}
