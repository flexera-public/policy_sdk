// Code generated by goa v3.0.6, DO NOT EDIT.
//
// Health client
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package health

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Health" service client.
type Client struct {
	CheckEndpoint  goa.Endpoint
	ReportEndpoint goa.Endpoint
}

// NewClient initializes a "Health" service client given the endpoints.
func NewClient(check, report goa.Endpoint) *Client {
	return &Client{
		CheckEndpoint:  check,
		ReportEndpoint: report,
	}
}

// Check calls the "check" endpoint of the "Health" service.
func (c *Client) Check(ctx context.Context) (err error) {
	_, err = c.CheckEndpoint(ctx, nil)
	return
}

// Report calls the "report" endpoint of the "Health" service.
func (c *Client) Report(ctx context.Context) (res map[string]string, err error) {
	var ires interface{}
	ires, err = c.ReportEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(map[string]string), nil
}
