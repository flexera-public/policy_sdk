// Code generated by goa v3.1.3, DO NOT EDIT.
//
// AppliedPolicy endpoints
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design -o .

package appliedpolicy

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "AppliedPolicy" service endpoints.
type Endpoints struct {
	Create     goa.Endpoint
	Delete     goa.Endpoint
	Show       goa.Endpoint
	ShowStatus goa.Endpoint
	ShowLog    goa.Endpoint
	Index      goa.Endpoint
	Evaluate   goa.Endpoint
	Update     goa.Endpoint
}

// NewEndpoints wraps the methods of the "AppliedPolicy" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Create:     NewCreateEndpoint(s, a.JWTAuth),
		Delete:     NewDeleteEndpoint(s, a.JWTAuth),
		Show:       NewShowEndpoint(s, a.JWTAuth),
		ShowStatus: NewShowStatusEndpoint(s, a.JWTAuth),
		ShowLog:    NewShowLogEndpoint(s, a.JWTAuth),
		Index:      NewIndexEndpoint(s, a.JWTAuth),
		Evaluate:   NewEvaluateEndpoint(s, a.JWTAuth),
		Update:     NewUpdateEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "AppliedPolicy" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Delete = m(e.Delete)
	e.Show = m(e.Show)
	e.ShowStatus = m(e.ShowStatus)
	e.ShowLog = m(e.ShowLog)
	e.Index = m(e.Index)
	e.Evaluate = m(e.Evaluate)
	e.Update = m(e.Update)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "AppliedPolicy".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:create||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAppliedPolicy(res, view)
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "AppliedPolicy".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:delete||common:project:own||common:org:own"},
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
// service "AppliedPolicy".
func NewShowEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:show||common:project:own||common:org:own"},
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
		vres := NewViewedAppliedPolicy(res, view)
		return vres, nil
	}
}

// NewShowStatusEndpoint returns an endpoint function that calls the method
// "show_status" of service "AppliedPolicy".
func NewShowStatusEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowStatusPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:show||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ShowStatus(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAppliedPolicyStatus(res, "default")
		return vres, nil
	}
}

// NewShowLogEndpoint returns an endpoint function that calls the method
// "show_log" of service "AppliedPolicy".
func NewShowLogEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowLogPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:show_log||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ShowLog(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAppliedPolicyLog(res, "default")
		return vres, nil
	}
}

// NewIndexEndpoint returns an endpoint function that calls the method "index"
// of service "AppliedPolicy".
func NewIndexEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IndexPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:index||common:project:own||common:org:own"},
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
		vres := NewViewedAppliedPolicyList(res, view)
		return vres, nil
	}
}

// NewEvaluateEndpoint returns an endpoint function that calls the method
// "evaluate" of service "AppliedPolicy".
func NewEvaluateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EvaluatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:evaluate||common:project:own||common:org:own"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Evaluate(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "AppliedPolicy".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "GlobalSession",
			Scopes:         []string{"governance:policy_template:compile||common:project:own||common:org:own", "governance:policy_template:upload||common:project:own||common:org:own", "governance:policy_template:update||common:project:own||common:org:own", "governance:policy_template:delete||common:project:own||common:org:own", "governance:policy_template:show||common:project:own||common:org:own", "governance:policy_template:index||common:project:own||common:org:own", "governance:policy_template:retrieve_data||common:project:own||common:org:own", "governance:published_template:create||common:org:own", "governance:published_template:update||common:org:own", "governance:published_template:hide||common:org:own", "governance:published_template:unhide||common:org:own", "governance:published_template:delete||common:org:own", "governance:published_template:show||common:org:own", "governance:published_template:index||common:org:own", "governance:applied_policy:create||common:project:own||common:org:own", "governance:applied_policy:delete||common:project:own||common:org:own", "governance:applied_policy:show||common:project:own||common:org:own", "governance:applied_policy:show_log||common:project:own||common:org:own", "governance:applied_policy:index||common:project:own||common:org:own", "governance:applied_policy:evaluate||common:project:own||common:org:own", "governance:applied_policy:update||common:project:own||common:org:own", "governance:policy_aggregate:create||common:org:own", "governance:policy_aggregate:update||common:org:own", "governance:policy_aggregate:delete||common:org:own", "governance:policy_aggregate:show||common:org:own", "governance:policy_aggregate:index||common:org:own", "governance:incident:resolve||common:project:own||common:org:own", "governance:incident:show||common:project:own||common:org:own", "governance:incident:index||common:project:own||common:org:own", "governance:incident:run_action||common:project:own||common:org:own", "governance:archived_incident:show||common:project:own||common:org:own", "governance:archived_incident:index||common:project:own||common:org:own", "governance:incident_aggregate:show||common:org:own", "governance:incident_aggregate:index||common:org:own", "governance:approval_request:show||common:project:own||common:org:own", "governance:approval_request:index||common:project:own||common:org:own", "governance:approval_request:approve||common:project:own||common:org:own", "governance:approval_request:deny||common:project:own||common:org:own", "governance:action_status:show||common:project:own||common:org:own", "governance:action_status:index||common:project:own||common:org:own"},
			RequiredScopes: []string{"governance:applied_policy:update||common:project:own||common:org:own"},
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
