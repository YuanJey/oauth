package oauth

type AuthProvider interface {
	GetAuthorizationURL(redirectURI string, state string) string
	ExchangeCodeForToken(code string) (string, error)
	GetUserInfo(accessToken string) (map[string]interface{}, error)
}
