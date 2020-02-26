// Code generated by goa v3.0.6, DO NOT EDIT.
//
// IncidentAggregate HTTP server
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"context"
	"net/http"
	"regexp"

	incidentaggregate "github.com/rightscale/governance/front_service/gen/incident_aggregate"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the IncidentAggregate service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	Show           http.Handler
	ShowNonCatalog http.Handler
	Index          http.Handler
	CORS           http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the IncidentAggregate service
// endpoints.
func New(
	e *incidentaggregate.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Show", "GET", "/api/governance/orgs/{org_id}/incident_aggregates/{incident_aggregate_id}"},
			{"ShowNonCatalog", "GET", "/api/governance/orgs/{org_id}/incident_aggregates/non_catalog"},
			{"Index", "GET", "/api/governance/orgs/{org_id}/incident_aggregates"},
			{"CORS", "OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates/{incident_aggregate_id}"},
			{"CORS", "OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates/non_catalog"},
			{"CORS", "OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates"},
		},
		Show:           NewShowHandler(e.Show, mux, dec, enc, eh),
		ShowNonCatalog: NewShowNonCatalogHandler(e.ShowNonCatalog, mux, dec, enc, eh),
		Index:          NewIndexHandler(e.Index, mux, dec, enc, eh),
		CORS:           NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "IncidentAggregate" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Show = m(s.Show)
	s.ShowNonCatalog = m(s.ShowNonCatalog)
	s.Index = m(s.Index)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the IncidentAggregate endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountShowHandler(mux, h.Show)
	MountShowNonCatalogHandler(mux, h.ShowNonCatalog)
	MountIndexHandler(mux, h.Index)
	MountCORSHandler(mux, h.CORS)
}

// MountShowHandler configures the mux to serve the "IncidentAggregate" service
// "show" endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleIncidentAggregateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/governance/orgs/{org_id}/incident_aggregates/{incident_aggregate_id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "IncidentAggregate" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, dec)
		encodeResponse = EncodeShowResponse(enc)
		encodeError    = EncodeShowError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "IncidentAggregate")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountShowNonCatalogHandler configures the mux to serve the
// "IncidentAggregate" service "show_non_catalog" endpoint.
func MountShowNonCatalogHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleIncidentAggregateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/governance/orgs/{org_id}/incident_aggregates/non_catalog", f)
}

// NewShowNonCatalogHandler creates a HTTP handler which loads the HTTP request
// and calls the "IncidentAggregate" service "show_non_catalog" endpoint.
func NewShowNonCatalogHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeShowNonCatalogRequest(mux, dec)
		encodeResponse = EncodeShowNonCatalogResponse(enc)
		encodeError    = EncodeShowNonCatalogError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show_non_catalog")
		ctx = context.WithValue(ctx, goa.ServiceKey, "IncidentAggregate")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountIndexHandler configures the mux to serve the "IncidentAggregate"
// service "index" endpoint.
func MountIndexHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleIncidentAggregateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/governance/orgs/{org_id}/incident_aggregates", f)
}

// NewIndexHandler creates a HTTP handler which loads the HTTP request and
// calls the "IncidentAggregate" service "index" endpoint.
func NewIndexHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeIndexRequest(mux, dec)
		encodeResponse = EncodeIndexResponse(enc)
		encodeError    = EncodeIndexError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "index")
		ctx = context.WithValue(ctx, goa.ServiceKey, "IncidentAggregate")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service IncidentAggregate.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleIncidentAggregateOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates/{incident_aggregate_id}", f)
	mux.Handle("OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates/non_catalog", f)
	mux.Handle("OPTIONS", "/api/governance/orgs/{org_id}/incident_aggregates", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleIncidentAggregateOrigin applies the CORS response headers
// corresponding to the origin for the service IncidentAggregate.
func handleIncidentAggregateOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("[.]rightscale[.]com$")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "ETag, Location")
			w.Header().Set("Access-Control-Max-Age", "900")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, POST, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Api-Version, Authorization, Client-Name, Content-Type, Csrf-Token, Prefer, If-Modified-Since, If-None-Match, If-Unmodified-Since, If-Match, X-Api-Version, X-Csrf-Token, X-Requested-With")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
