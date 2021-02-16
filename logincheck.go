// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth_urls

import (
	"github.com/arelate/gog_urls"
	"net/url"
)

func LoginCheck() *url.URL {
	return &url.URL{
		Scheme: gog_urls.HttpsScheme,
		Host:   loginHost,
		Path:   loginCheckPath,
	}
}
