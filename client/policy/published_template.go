// Copyright (c) 2018 RightScale, Inc. - see LICENSE

package policy

// import (
// 	"context"
// 	"net/http"
// 	"net/url"
// 	"time"

// 	"github.com/pkg/errors"
// 	"github.com/rightscale/right_pt/auth"
// 	"github.com/rightscale/right_pt/client_wrappers"
// 	pubclient "github.com/rightscale/right_pt/sdk/http/published_template/client"
// 	"github.com/rightscale/right_pt/sdk/published_template"
// 	"goa.design/goa"
// 	goahttp "goa.design/goa/http"
// )

// // CreatePublishedTemplate a published template
// func (c *client) CreatePublishedTemplate(ctx context.Context, templateHref string) (*publishedtemplate.CreateResult, error) {
// 	token := c.getToken()
// 	p := &publishedtemplate.CreatePayload{
// 		Token:        token,
// 		OrgID:        c.orgID,
// 		APIVersion:   apiVersion,
// 		TemplateHref: templateHref,
// 	}
// 	resIF, err := c.create(ctx, p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, ok := resIF.(*publishedtemplate.CreateResult)
// 	if !ok {
// 		return nil, errors.New("error interpreting published template create result")
// 	}
// 	return res, nil
// }

// // DeletePublishedTemplate a published template
// func (c *client) DeletePublishedTemplate(ctx context.Context, templateID string) error {
// 	token := c.getToken()
// 	p := &publishedtemplate.DeletePayload{
// 		Token:      token,
// 		OrgID:      c.orgID,
// 		APIVersion: apiVersion,
// 		TemplateID: templateID,
// 	}
// 	_, err := c.delete(ctx, p)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // ShowPublishedTemplate a published template
// func (c *client) ShowPublishedTemplate(ctx context.Context, templateID, view string) (*publishedtemplate.PublishedTemplate, error) {
// 	token := c.getToken()
// 	p := &publishedtemplate.ShowPayload{
// 		Token:      token,
// 		OrgID:      c.orgID,
// 		APIVersion: apiVersion,
// 		TemplateID: templateID,
// 	}
// 	if view != "" {
// 		p.View = &view
// 	}
// 	pubIF, err := c.show(ctx, p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pub, ok := pubIF.(*publishedtemplate.PublishedTemplate)
// 	if !ok {
// 		return nil, errors.New("error interpreting published template")
// 	}
// 	return pub, nil
// }

// // getToken generates an access token for us
// func (c *client) getToken() *string {
// 	t, err := c.ts.TokenString()
// 	if err != nil {
// 		return nil
// 	}
// 	return &t
// }

// func strPtr(s string) *string { return &s }
