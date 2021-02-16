// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth_urls

import (
	"github.com/arelate/gog_urls"
	"net/http"
)

func AddAuthHostDefaultHeaders(req *http.Request) {
	gog_urls.SetDefaultHeaders(req, authHost, gog_urls.AcceptTextHTML)
}

func AddLoginHostDefaultHeaders(req *http.Request) {
	gog_urls.SetDefaultHeaders(req, loginHost, gog_urls.AcceptTextHTML)
}
