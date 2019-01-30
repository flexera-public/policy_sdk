// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// PolicyTemplate client
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package policytemplate

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "PolicyTemplate" service client.
type Client struct {
	CompileEndpoint goa.Endpoint
	UploadEndpoint  goa.Endpoint
	UpdateEndpoint  goa.Endpoint
	ShowEndpoint    goa.Endpoint
	IndexEndpoint   goa.Endpoint
	DeleteEndpoint  goa.Endpoint
}

// NewClient initializes a "PolicyTemplate" service client given the endpoints.
func NewClient(compile, upload, update, show, index, delete_ goa.Endpoint) *Client {
	return &Client{
		CompileEndpoint: compile,
		UploadEndpoint:  upload,
		UpdateEndpoint:  update,
		ShowEndpoint:    show,
		IndexEndpoint:   index,
		DeleteEndpoint:  delete_,
	}
}

// Compile calls the "compile" endpoint of the "PolicyTemplate" service.
// Compile may return the following errors:
//	- "invalid_template" (type *CompilationErrors)
//	- error: internal error
func (c *Client) Compile(ctx context.Context, p *CompilePayload) (err error) {
	_, err = c.CompileEndpoint(ctx, p)
	return
}

// Upload calls the "upload" endpoint of the "PolicyTemplate" service.
// Upload may return the following errors:
//	- "invalid_template" (type *CompilationErrors)
//	- "conflict" (type *ConflictError)
//	- error: internal error
func (c *Client) Upload(ctx context.Context, p *UploadPayload) (res *PolicyTemplate, err error) {
	var ires interface{}
	ires, err = c.UploadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PolicyTemplate), nil
}

// Update calls the "update" endpoint of the "PolicyTemplate" service.
// Update may return the following errors:
//	- "invalid_template" (type *CompilationErrors)
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *PolicyTemplate, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PolicyTemplate), nil
}

// Show calls the "show" endpoint of the "PolicyTemplate" service.
// Show may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *PolicyTemplate, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PolicyTemplate), nil
}

// Index calls the "index" endpoint of the "PolicyTemplate" service.
func (c *Client) Index(ctx context.Context, p *IndexPayload) (res *PolicyTemplateList, err error) {
	var ires interface{}
	ires, err = c.IndexEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PolicyTemplateList), nil
}

// Delete calls the "delete" endpoint of the "PolicyTemplate" service.
// Delete may return the following errors:
//	- "not_found" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}