package auth

// TokenSource is similar to rsauth TokenSource -- it generates JWT bearer tokens.
type TokenSource interface {
	TokenString() (token string, err error)
}
