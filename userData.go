package gogauthurls

import (
	"github.com/boggydigital/gogog/pkg/gogurls"
	"net/url"
)

func UserData() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   gogurls.WwwGogHost,
		Path:   userDataPath,
	}
}
