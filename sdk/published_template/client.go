// Code generated by goa v3.0.6, DO NOT EDIT.
//
// PublishedTemplate client
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package publishedtemplate

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "PublishedTemplate" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
	HideEndpoint   goa.Endpoint
	UnhideEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
	ShowEndpoint   goa.Endpoint
	IndexEndpoint  goa.Endpoint
}

// NewClient initializes a "PublishedTemplate" service client given the
// endpoints.
func NewClient(create, update, hide, unhide, delete_, show, index goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
		UpdateEndpoint: update,
		HideEndpoint:   hide,
		UnhideEndpoint: unhide,
		DeleteEndpoint: delete_,
		ShowEndpoint:   show,
		IndexEndpoint:  index,
	}
}

// Create calls the "create" endpoint of the "PublishedTemplate" service.
// Create may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- "conflict" (type *ConflictError)
//	- error: internal error
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *CreateResult, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateResult), nil
}

// Update calls the "update" endpoint of the "PublishedTemplate" service.
// Update may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Hide calls the "hide" endpoint of the "PublishedTemplate" service.
// Hide may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Hide(ctx context.Context, p *HidePayload) (err error) {
	_, err = c.HideEndpoint(ctx, p)
	return
}

// Unhide calls the "unhide" endpoint of the "PublishedTemplate" service.
// Unhide may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Unhide(ctx context.Context, p *UnhidePayload) (err error) {
	_, err = c.UnhideEndpoint(ctx, p)
	return
}

// Delete calls the "delete" endpoint of the "PublishedTemplate" service.
// Delete may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// Show calls the "show" endpoint of the "PublishedTemplate" service.
// Show may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *PublishedTemplate, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PublishedTemplate), nil
}

// Index calls the "index" endpoint of the "PublishedTemplate" service.
func (c *Client) Index(ctx context.Context, p *IndexPayload) (res *PublishedTemplateList, err error) {
	var ires interface{}
	ires, err = c.IndexEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PublishedTemplateList), nil
}
