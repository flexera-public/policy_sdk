// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ArchivedIncident client
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package archivedincident

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ArchivedIncident" service client.
type Client struct {
	ShowEndpoint             goa.Endpoint
	IndexEndpoint            goa.Endpoint
	IndexEscalationsEndpoint goa.Endpoint
	IndexResolutionsEndpoint goa.Endpoint
}

// NewClient initializes a "ArchivedIncident" service client given the
// endpoints.
func NewClient(show, index, indexEscalations, indexResolutions goa.Endpoint) *Client {
	return &Client{
		ShowEndpoint:             show,
		IndexEndpoint:            index,
		IndexEscalationsEndpoint: indexEscalations,
		IndexResolutionsEndpoint: indexResolutions,
	}
}

// Show calls the "show" endpoint of the "ArchivedIncident" service.
// Show may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *ArchivedIncident, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ArchivedIncident), nil
}

// Index calls the "index" endpoint of the "ArchivedIncident" service.
func (c *Client) Index(ctx context.Context, p *IndexPayload) (res *ArchivedIncidentList, err error) {
	var ires interface{}
	ires, err = c.IndexEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ArchivedIncidentList), nil
}

// IndexEscalations calls the "index_escalations" endpoint of the
// "ArchivedIncident" service.
// IndexEscalations may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) IndexEscalations(ctx context.Context, p *IndexEscalationsPayload) (res *Escalations, err error) {
	var ires interface{}
	ires, err = c.IndexEscalationsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Escalations), nil
}

// IndexResolutions calls the "index_resolutions" endpoint of the
// "ArchivedIncident" service.
// IndexResolutions may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- "unprocessable_entity" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) IndexResolutions(ctx context.Context, p *IndexResolutionsPayload) (res *Resolutions, err error) {
	var ires interface{}
	ires, err = c.IndexResolutionsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Resolutions), nil
}
