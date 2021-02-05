package gogauthurls

import (
	"net/http"
	"net/url"
)

func SetLoginFormHeaders(req *http.Request, referer *url.URL) {
	const (
		urlEncoded = "application/x-www-form-urlencoded"
	)

	req.Header.Set("Origin", loginHost)
	req.Header.Set("Referer", referer.String())
	req.Header.Set("Content-Type", urlEncoded)
}
