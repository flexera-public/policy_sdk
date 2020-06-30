package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"

	goahttp "goa.design/goa/v3/http"
)

// refreshTokenSource contains the logic to create new session using OAuth tokens
type refreshTokenSource struct {
	endpoint     *url.URL
	refreshToken string
	accessToken  string
	refreshAt    time.Time
	doer         goahttp.Doer
}

// Grant from CM API 1.5
type Grant struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

// NewOAuthAuthenticator returns a authenticator that uses a oauth refresh
// token to create access tokens. The refresh token can be found in the CM
// dashboard under Settings > Account Settings > API Credentials.
func NewOAuthAuthenticator(host string, refreshToken string) (TokenSource, error) {
	if !validHost(host) {
		return nil, fmt.Errorf("invalid authentication host, must be a form like us-3.rightscale.com")
	}
	endpoint, _ := url.Parse(fmt.Sprintf("https://%s/api/oauth2", host))
	return &refreshTokenSource{
		endpoint:     endpoint,
		refreshToken: refreshToken,
		refreshAt:    time.Now().Add(-2 * time.Minute),
		doer:         http.DefaultClient,
	}, nil
}

func (ts *refreshTokenSource) TokenString() (string, error) {
	if ts.tokenValid() {
		return ts.accessToken, nil
	}
	body := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, ts.refreshToken)
	req, err := http.NewRequest("POST", ts.endpoint.String(), bytes.NewBufferString(body))
	if err != nil {
		return "", fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	req.Header.Set("X-API-Version", "1.5")
	req.Header.Set("Content-Type", "application/json")
	resp, err := ts.doer.Do(req)
	if err != nil {
		return "", fmt.Errorf("Authentication failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		var msg string
		if err != nil {
			msg = " - <failed to read body>"
		}
		if len(body) > 0 {
			msg = " - " + string(body)
		}
		return "", fmt.Errorf("Authentication failed: %s%s", resp.Status, msg)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Authentication failed (failed to read response): %s", err)
	}
	var g Grant
	err = json.Unmarshal(jsonBytes, &g)
	if err != nil {
		return "", fmt.Errorf("Authentication failed (failed to load response JSON): %s", err)
	}
	if g.AccessToken == "" || g.ExpiresIn == 0 {
		return "", fmt.Errorf("Authentication failed: cloud not parse grant: %+v", g)
	}

	ts.accessToken = g.AccessToken
	ts.refreshAt = time.Now().Add(time.Duration(g.ExpiresIn/2) * time.Second)
	return ts.accessToken, nil
}

func (ts *refreshTokenSource) tokenValid() bool {
	return ts.accessToken != "" && ts.refreshAt.Before(time.Now())
}

func validHost(host string) bool {
	hostRe := regexp.MustCompile(`(us|telstra|moo)-(\d+)\.(test.)?rightscale\.com`)
	return hostRe.MatchString(host)
}
