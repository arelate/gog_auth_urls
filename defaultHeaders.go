// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth_urls

import "net/http"

func addDefaultHeaders(req *http.Request, host string) {
	const (
		acceptHeader         = "text/html"
		acceptLanguageHeader = "en-us"
		connectionHeader     = "keep-alive"
		userAgentHeader      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/87.0.4280.141 " +
			"Safari/537.36 " +
			"Edg/87.0.664.75" // Microsoft Edge 87 UA string
	)
	if host != "" {
		req.Host = host
	}
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("Accept-Language", acceptLanguageHeader)
	req.Header.Set("Connection", connectionHeader)
	req.Header.Set("User-Agent", userAgentHeader)
}

func AddAuthHostDefaultHeaders(req *http.Request) {
	addDefaultHeaders(req, authHost)
}

func AddLoginHostDefaultHeaders(req *http.Request) {
	addDefaultHeaders(req, loginHost)
}
