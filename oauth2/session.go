package oauth2

import "github.com/ory-am/fosite/handler/oidc/strategy"

type Session struct {
	Subject                  string `json:"sub"`
	*strategy.DefaultSession `json:"idToken"`
}
