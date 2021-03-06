// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package swagger

import (
	api "code.gitea.io/gitea/modules/structs"
)

// OAuth2Application
// swagger:response OAuth2Application
type swaggerResponseOAuth2Application struct {
	// in:body
	Body api.OAuth2Application `json:"body"`
}
