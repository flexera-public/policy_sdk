package auth

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // basicTokenSource builds login requests from users email and password.
// type basicTokenSource struct {
// 	accessToken string
// 	refreshAt   *time.Time
// 	username    string
// 	password    string
// 	accountID   int
// }

// // NewBasicAuthenticator returns a authenticator that uses email and password to create sessions.
// // The returned authenticator takes care of refreshing the RightScale session as needed.
// func NewBasicAuthenticator(username, password string, accountID int) Authenticator {
// 	return &basicTokenSource{
// 		refreshAt: time.Now().Add(-2 * time.Minute),
// 		username:  username,
// 		password:  password,
// 		accountID: accountID}
// }

// // buildLoginRequest builds session create requests from users email and password.
// func (b *basicTokenSource) buildLoginRequest(host string) (*http.Request, error) {
// 	if host == "" {
// 		host = "us-3.rightscale.com"
// 	}
// 	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
// 		b.username, b.password, b.accountID)
// 	authReq, err := http.NewRequest("POST", buildURL(host, "api/sessions"),
// 		bytes.NewBufferString(jsonStr))
// 	if err != nil {
// 		return authReq, fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
// 	}
// 	authReq.Header.Set("X-API-Version", "1.5")
// 	authReq.Header.Set("Content-Type", "application/json")
// 	return authReq, nil
// }

// func (ts *basicTokenSource) TokenString() (string, error) {
// 	if ts.tokenValid() {
// 		return ts.accessToken, nil
// 	}

// 	authReq, authErr := s.builder.BuildLoginRequest(s.host)
// 	if authErr != nil {
// 		return authErr
// 	}
// 	resp, err := s.client.DoHidden(authReq)
// 	if err != nil {
// 		return err
// 	}
// 	url, err := extractRedirectURL(resp)
// 	if err != nil {
// 		return err
// 	}
// 	if url != nil {
// 		authReq, authErr = s.builder.BuildLoginRequest(url.Host)
// 		if authErr != nil {
// 			return authErr
// 		}
// 		s.host = url.Host
// 		req.Host = url.Host
// 		req.URL.Host = url.Host
// 		resp, err = s.client.DoHidden(authReq)
// 	}
// 	if err != nil {
// 		return fmt.Errorf("Authentication failed: %s", err)
// 	}
// 	if err := s.refresh(resp); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (ts *basicTokenSource) tokenValid() bool {
// 	return ts.accessToken != "" && ts.refreshAt.Before(time.Now())
// }
