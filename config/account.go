package config

// type Client struct {
// 	polc polc.Client
// 	pubc pubc.Client
// 	ic   ic.Client
// 	apc  apc.Client
// }

// func (a *Account) PolicyTemplateClient() (polc.Client, error) {
// 	if account.policyClient == nil {
// 		if err := account.validate(); err != nil {
// 			return nil, err
// 		}
// 		tokenSource := auth.NewOAuthAuthenticator(account.RefreshToken, account.ID)
// 		account.policyClient = policy.NewClient(account.Host, tokenSource, false)
// 	}
// 	return account.policyClient, nil
// }

// func NewClient(a *config.Account, ts auth.TokenSource, debug bool) Client {
// 	return &Client{
// 		polc: policytemplate.NewClient(url, ts, debug),
// 		pubc: policytemplate.NewClient(url, ts, debug),
// 		ic:   policytemplate.NewClient(url, ts, debug),
// 		apc:  policytemplate.NewClient(url, ts, debug),
// 	}
// }
