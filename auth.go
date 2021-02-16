// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth_urls

import (
	"github.com/arelate/gog_urls"
	"net/url"
)

func auth(host string) *url.URL {
	authURL := &url.URL{
		Scheme: gog_urls.HttpsScheme,
		Host:   host,
		Path:   authPath,
	}

	setAuthQuery(authURL)

	return authURL
}

func setAuthQuery(url *url.URL) {
	q := url.Query()
	q.Set("client_id", "46755278331571209")
	q.Set("redirect_uri", OnLoginSuccess().String())
	q.Set("response_type", "code")
	q.Set("layout", "default")
	q.Set("brand", "gog")
	q.Set("gog_lc", "en-US")
	url.RawQuery = q.Encode()
}

func AuthHost() *url.URL {
	return auth(authHost)
}

func LoginHost() *url.URL {
	return auth(loginHost)
}
