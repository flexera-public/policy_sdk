// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ActionStatus client
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package actionstatus

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ActionStatus" service client.
type Client struct {
	IndexEndpoint goa.Endpoint
	ShowEndpoint  goa.Endpoint
}

// NewClient initializes a "ActionStatus" service client given the endpoints.
func NewClient(index, show goa.Endpoint) *Client {
	return &Client{
		IndexEndpoint: index,
		ShowEndpoint:  show,
	}
}

// Index calls the "index" endpoint of the "ActionStatus" service.
func (c *Client) Index(ctx context.Context, p *IndexPayload) (res *ActionStatusList, err error) {
	var ires interface{}
	ires, err = c.IndexEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ActionStatusList), nil
}

// Show calls the "show" endpoint of the "ActionStatus" service.
// Show may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *ActionStatus, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ActionStatus), nil
}
