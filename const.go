// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth_urls

import "github.com/arelate/gog_urls"

// hosts
const (
	loginHost     = "login." + gog_urls.GogHost
	authHost      = "auth." + gog_urls.GogHost
	reCaptchaHost = "www.recaptcha.net"
)

// paths
const (
	authPath           = "/auth"
	loginCheckPath     = "/login_check"
	loginTwoStepPath   = "/login/two_step"
	reCaptchaPath      = "/recaptcha/api.js"
	userDataPath       = "/userData.json"
	onLoginSuccessPath = "/on_login_success"
)
