// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogauthurls

import (
	"github.com/boggydigital/gogog/pkg/gogurls"
	"net/url"
)

func LoginCheck() *url.URL {
	return &url.URL{
		Scheme: gogurls.HttpsScheme,
		Host:   loginHost,
		Path:   loginCheckPath,
	}
}
