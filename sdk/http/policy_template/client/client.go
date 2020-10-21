// Code generated by goa v3.1.3, DO NOT EDIT.
//
// PolicyTemplate client HTTP transport
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the PolicyTemplate service endpoint HTTP clients.
type Client struct {
	// Compile Doer is the HTTP client used to make requests to the compile
	// endpoint.
	CompileDoer goahttp.Doer

	// Upload Doer is the HTTP client used to make requests to the upload endpoint.
	UploadDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// RetrieveData Doer is the HTTP client used to make requests to the
	// retrieve_data endpoint.
	RetrieveDataDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Index Doer is the HTTP client used to make requests to the index endpoint.
	IndexDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the PolicyTemplate service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CompileDoer:         doer,
		UploadDoer:          doer,
		UpdateDoer:          doer,
		RetrieveDataDoer:    doer,
		ShowDoer:            doer,
		IndexDoer:           doer,
		DeleteDoer:          doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Compile returns an endpoint that makes HTTP requests to the PolicyTemplate
// service compile server.
func (c *Client) Compile() goa.Endpoint {
	var (
		encodeRequest  = EncodeCompileRequest(c.encoder)
		decodeResponse = DecodeCompileResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCompileRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CompileDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "compile", err)
		}
		return decodeResponse(resp)
	}
}

// Upload returns an endpoint that makes HTTP requests to the PolicyTemplate
// service upload server.
func (c *Client) Upload() goa.Endpoint {
	var (
		encodeRequest  = EncodeUploadRequest(c.encoder)
		decodeResponse = DecodeUploadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUploadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UploadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "upload", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the PolicyTemplate
// service update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "update", err)
		}
		return decodeResponse(resp)
	}
}

// RetrieveData returns an endpoint that makes HTTP requests to the
// PolicyTemplate service retrieve_data server.
func (c *Client) RetrieveData() goa.Endpoint {
	var (
		encodeRequest  = EncodeRetrieveDataRequest(c.encoder)
		decodeResponse = DecodeRetrieveDataResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRetrieveDataRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RetrieveDataDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "retrieve_data", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the PolicyTemplate
// service show server.
func (c *Client) Show() goa.Endpoint {
	var (
		encodeRequest  = EncodeShowRequest(c.encoder)
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Index returns an endpoint that makes HTTP requests to the PolicyTemplate
// service index server.
func (c *Client) Index() goa.Endpoint {
	var (
		encodeRequest  = EncodeIndexRequest(c.encoder)
		decodeResponse = DecodeIndexResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildIndexRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.IndexDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "index", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the PolicyTemplate
// service delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("PolicyTemplate", "delete", err)
		}
		return decodeResponse(resp)
	}
}
