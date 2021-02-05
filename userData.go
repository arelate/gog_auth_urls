package gogauthurls

import (
	"github.com/arelate/gogurls"
	"net/url"
)

func UserData() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   gogurls.WwwGogHost,
		Path:   userDataPath,
	}
}
