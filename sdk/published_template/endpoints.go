// Code generated by goa v3.0.6, DO NOT EDIT.
//
// PublishedTemplate endpoints
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package publishedtemplate

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "PublishedTemplate" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Update goa.Endpoint
	Hide   goa.Endpoint
	Unhide goa.Endpoint
	Delete goa.Endpoint
	Show   goa.Endpoint
	Index  goa.Endpoint
}

// NewEndpoints wraps the methods of the "PublishedTemplate" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Create: NewCreateEndpoint(s, a.JWTAuth),
		Update: NewUpdateEndpoint(s, a.JWTAuth),
		Hide:   NewHideEndpoint(s, a.JWTAuth),
		Unhide: NewUnhideEndpoint(s, a.JWTAuth),
		Delete: NewDeleteEndpoint(s, a.JWTAuth),
		Show:   NewShowEndpoint(s, a.JWTAuth),
		Index:  NewIndexEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "PublishedTemplate" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Hide = m(e.Hide)
	e.Unhide = m(e.Unhide)
	e.Delete = m(e.Delete)
	e.Show = m(e.Show)
	e.Index = m(e.Index)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "PublishedTemplate".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:create"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Create(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "PublishedTemplate".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:update"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Update(ctx, p)
	}
}

// NewHideEndpoint returns an endpoint function that calls the method "hide" of
// service "PublishedTemplate".
func NewHideEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*HidePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:hide"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Hide(ctx, p)
	}
}

// NewUnhideEndpoint returns an endpoint function that calls the method
// "unhide" of service "PublishedTemplate".
func NewUnhideEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UnhidePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:unhide"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Unhide(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "PublishedTemplate".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:delete"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Delete(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "PublishedTemplate".
func NewShowEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:show"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPublishedTemplate(res, view)
		return vres, nil
	}
}

// NewIndexEndpoint returns an endpoint function that calls the method "index"
// of service "PublishedTemplate".
func NewIndexEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IndexPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile", "governance:policy_template:upload", "governance:policy_template:update", "governance:policy_template:delete", "governance:policy_template:show", "governance:policy_template:index", "governance:policy_template:retrieve_data", "governance:published_template:create", "governance:published_template:update", "governance:published_template:hide", "governance:published_template:unhide", "governance:published_template:delete", "governance:published_template:show", "governance:published_template:index", "governance:applied_policy:create", "governance:applied_policy:delete", "governance:applied_policy:show", "governance:applied_policy:show_log", "governance:applied_policy:index", "governance:applied_policy:evaluate", "governance:policy_aggregate:create", "governance:policy_aggregate:update", "governance:policy_aggregate:delete", "governance:policy_aggregate:show", "governance:policy_aggregate:index", "governance:incident:resolve", "governance:incident:show", "governance:incident:index", "governance:archived_incident:show", "governance:archived_incident:index", "governance:incident_aggregate:show", "governance:incident_aggregate:index", "governance:approval_request:show", "governance:approval_request:index", "governance:approval_request:approve", "governance:approval_request:deny"},
			RequiredScopes: []string{"governance:published_template:index"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Index(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedPublishedTemplateList(res, view)
		return vres, nil
	}
}
