package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	goahttp "goa.design/goa/v3/http"
)

type (
	// clientCredentialsTokenSource contains the logic to create new sessions using an OAuth2 client ID and secret
	clientCredentialsTokenSource struct {
		endpoint     *url.URL
		clientID     string
		clientSecret string
		accessToken  string
		refreshAt    time.Time
		doer         goahttp.Doer
	}
)

// NewOAuthClientCredentialsAuthenticator returns an authenticator that uses an OAuth2 client ID and secret
// to create access tokens.
func NewOAuthClientCredentialsAuthenticator(host, clientID, clientSecret string) (TokenSource, error) {
	if !flexeraHostRegexp.MatchString(host) {
		return nil, fmt.Errorf("invalid authentication host: %v does not match %v", host, flexeraHostRegexp)
	}
	endpoint, _ := url.Parse(fmt.Sprintf("https://%v/oidc/token", host))
	return &clientCredentialsTokenSource{
		endpoint:     endpoint,
		clientID:     clientID,
		clientSecret: clientSecret,
		doer:         http.DefaultClient,
	}, nil
}

func (ccts *clientCredentialsTokenSource) TokenString() (string, error) {
	if ccts.tokenValid() {
		return ccts.accessToken, nil
	}

	b, err := json.Marshal(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     ccts.clientID,
		"client_secret": ccts.clientSecret,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", ccts.endpoint.String(), bytes.NewBuffer(b))
	if err != nil {
		return "", fmt.Errorf("authentication failed (failed to build request): %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ccts.doer.Do(req)
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

	ccts.accessToken = g.AccessToken
	ccts.refreshAt = time.Now().Add(time.Duration(g.ExpiresIn/2) * time.Second)
	return ccts.accessToken, nil
}

func (ccts *clientCredentialsTokenSource) tokenValid() bool {
	return ccts.accessToken != "" && !ccts.refreshAt.Before(time.Now())
}
