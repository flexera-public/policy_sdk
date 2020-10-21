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

type (
	// flexeraRefreshTokenSource contains the logic to create new sessions using OAuth2 tokens
	flexeraRefreshTokenSource struct {
		endpoint     *url.URL
		refreshToken string
		accessToken  string
		refreshAt    time.Time
		doer         goahttp.Doer
	}

	// rsRefreshTokenSource contains the logic to create new session using OAuth tokens using the legacy RightScale login system
	rsRefreshTokenSource struct {
		endpoint     *url.URL
		refreshToken string
		accessToken  string
		refreshAt    time.Time
		doer         goahttp.Doer
	}

	// Grant represents the grant response returned by the token endpoint
	Grant struct {
		ExpiresIn   int    `json:"expires_in"`
		AccessToken string `json:"access_token"`
	}
)

var (
	flexeraHostRegexp = regexp.MustCompile(`^login\.flexera(?:test)?\.(?:com|eu)$`)
	cmHostRegexp      = regexp.MustCompile(`^(us|telstra|moo)-(\d+)\.(test.)?rightscale\.com$`)
)

// NewOAuthAuthenticator returns a authenticator that uses a oauth refresh
// token to create access tokens. The refresh token can be found in the CM
// dashboard under Settings > Account Settings > API Credentials.
func NewOAuthAuthenticator(host string, refreshToken string) (TokenSource, error) {
	if cmHostRegexp.MatchString(host) {
		endpoint, _ := url.Parse(fmt.Sprintf("https://%s/api/oauth2", host))
		return &rsRefreshTokenSource{
			endpoint:     endpoint,
			refreshToken: refreshToken,
			refreshAt:    time.Now().Add(-2 * time.Minute),
			doer:         http.DefaultClient,
		}, nil
	} else if !flexeraHostRegexp.MatchString(host) {
		return nil, fmt.Errorf("invalid authentication host: %v does not match %v or %v", host, flexeraHostRegexp, cmHostRegexp)
	}
	endpoint, _ := url.Parse(fmt.Sprintf("https://%v/oidc/token", host))
	return &flexeraRefreshTokenSource{
		endpoint:     endpoint,
		refreshToken: refreshToken,
		doer:         http.DefaultClient,
	}, nil
}

func (fts *flexeraRefreshTokenSource) TokenString() (string, error) {
	if fts.tokenValid() {
		return fts.accessToken, nil
	}

	b, err := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": fts.refreshToken,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", fts.endpoint.String(), bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := fts.doer.Do(req)
	if err != nil {
		return "", fmt.Errorf("authentication failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := ioutil.ReadAll(resp.Body)
		m := "<empty body>"
		if err != nil {
			m = "<failed to read body>"
		} else if len(b) > 0 {
			m = string(b)
		}
		return "", fmt.Errorf("authentication failed: %v: %v", resp.Status, m)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("authentication failed: failed to read response: %v", err)
	}
	var g Grant
	err = json.Unmarshal(b, &g)
	if err != nil {
		return "", fmt.Errorf("authentication failed: failed to parse response: %v", err)
	}
	if g.AccessToken == "" || g.ExpiresIn == 0 {
		return "", fmt.Errorf("authentication failed: grant missing field values: %+v", g)
	}

	fts.accessToken = g.AccessToken
	fts.refreshAt = time.Now().Add(time.Duration(g.ExpiresIn/2) * time.Second)
	return fts.accessToken, nil
}

func (fts *flexeraRefreshTokenSource) tokenValid() bool {
	return fts.accessToken != "" && !fts.refreshAt.Before(time.Now())
}

func (ts *rsRefreshTokenSource) TokenString() (string, error) {
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

func (ts *rsRefreshTokenSource) tokenValid() bool {
	return ts.accessToken != "" && !ts.refreshAt.Before(time.Now())
}
