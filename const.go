// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogauthurls

import "github.com/arelate/gogurls"

// hosts
const (
	loginHost     = "login." + gogurls.GogHost
	authHost      = "auth." + gogurls.GogHost
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
