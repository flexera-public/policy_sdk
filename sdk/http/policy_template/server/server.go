// Code generated by goa v3.0.6, DO NOT EDIT.
//
// PolicyTemplate HTTP server
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"context"
	"net/http"
	"regexp"

	policytemplate "github.com/rightscale/governance/front_service/gen/policy_template"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the PolicyTemplate service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	Compile      http.Handler
	Upload       http.Handler
	Update       http.Handler
	RetrieveData http.Handler
	Show         http.Handler
	Index        http.Handler
	Delete       http.Handler
	CORS         http.Handler
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

// New instantiates HTTP handlers for all the PolicyTemplate service endpoints.
func New(
	e *policytemplate.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Compile", "POST", "/api/governance/projects/{project_id}/policy_templates/compile"},
			{"Upload", "POST", "/api/governance/projects/{project_id}/policy_templates"},
			{"Update", "PUT", "/api/governance/projects/{project_id}/policy_templates/{template_id}"},
			{"RetrieveData", "POST", "/api/governance/projects/{project_id}/policy_templates/{template_id}/retrieve_data"},
			{"Show", "GET", "/api/governance/projects/{project_id}/policy_templates/{template_id}"},
			{"Index", "GET", "/api/governance/projects/{project_id}/policy_templates"},
			{"Delete", "DELETE", "/api/governance/projects/{project_id}/policy_templates/{template_id}"},
			{"CORS", "OPTIONS", "/api/governance/projects/{project_id}/policy_templates/compile"},
			{"CORS", "OPTIONS", "/api/governance/projects/{project_id}/policy_templates"},
			{"CORS", "OPTIONS", "/api/governance/projects/{project_id}/policy_templates/{template_id}"},
			{"CORS", "OPTIONS", "/api/governance/projects/{project_id}/policy_templates/{template_id}/retrieve_data"},
		},
		Compile:      NewCompileHandler(e.Compile, mux, dec, enc, eh),
		Upload:       NewUploadHandler(e.Upload, mux, dec, enc, eh),
		Update:       NewUpdateHandler(e.Update, mux, dec, enc, eh),
		RetrieveData: NewRetrieveDataHandler(e.RetrieveData, mux, dec, enc, eh),
		Show:         NewShowHandler(e.Show, mux, dec, enc, eh),
		Index:        NewIndexHandler(e.Index, mux, dec, enc, eh),
		Delete:       NewDeleteHandler(e.Delete, mux, dec, enc, eh),
		CORS:         NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "PolicyTemplate" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Compile = m(s.Compile)
	s.Upload = m(s.Upload)
	s.Update = m(s.Update)
	s.RetrieveData = m(s.RetrieveData)
	s.Show = m(s.Show)
	s.Index = m(s.Index)
	s.Delete = m(s.Delete)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the PolicyTemplate endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCompileHandler(mux, h.Compile)
	MountUploadHandler(mux, h.Upload)
	MountUpdateHandler(mux, h.Update)
	MountRetrieveDataHandler(mux, h.RetrieveData)
	MountShowHandler(mux, h.Show)
	MountIndexHandler(mux, h.Index)
	MountDeleteHandler(mux, h.Delete)
	MountCORSHandler(mux, h.CORS)
}

// MountCompileHandler configures the mux to serve the "PolicyTemplate" service
// "compile" endpoint.
func MountCompileHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/governance/projects/{project_id}/policy_templates/compile", f)
}

// NewCompileHandler creates a HTTP handler which loads the HTTP request and
// calls the "PolicyTemplate" service "compile" endpoint.
func NewCompileHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeCompileRequest(mux, dec)
		encodeResponse = EncodeCompileResponse(enc)
		encodeError    = EncodeCompileError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "compile")
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountUploadHandler configures the mux to serve the "PolicyTemplate" service
// "upload" endpoint.
func MountUploadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/governance/projects/{project_id}/policy_templates", f)
}

// NewUploadHandler creates a HTTP handler which loads the HTTP request and
// calls the "PolicyTemplate" service "upload" endpoint.
func NewUploadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeUploadRequest(mux, dec)
		encodeResponse = EncodeUploadResponse(enc)
		encodeError    = EncodeUploadError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "upload")
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountUpdateHandler configures the mux to serve the "PolicyTemplate" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/api/governance/projects/{project_id}/policy_templates/{template_id}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "PolicyTemplate" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, dec)
		encodeResponse = EncodeUpdateResponse(enc)
		encodeError    = EncodeUpdateError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountRetrieveDataHandler configures the mux to serve the "PolicyTemplate"
// service "retrieve_data" endpoint.
func MountRetrieveDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/governance/projects/{project_id}/policy_templates/{template_id}/retrieve_data", f)
}

// NewRetrieveDataHandler creates a HTTP handler which loads the HTTP request
// and calls the "PolicyTemplate" service "retrieve_data" endpoint.
func NewRetrieveDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRetrieveDataRequest(mux, dec)
		encodeResponse = EncodeRetrieveDataResponse(enc)
		encodeError    = EncodeRetrieveDataError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "retrieve_data")
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountShowHandler configures the mux to serve the "PolicyTemplate" service
// "show" endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/governance/projects/{project_id}/policy_templates/{template_id}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "PolicyTemplate" service "show" endpoint.
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
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountIndexHandler configures the mux to serve the "PolicyTemplate" service
// "index" endpoint.
func MountIndexHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/governance/projects/{project_id}/policy_templates", f)
}

// NewIndexHandler creates a HTTP handler which loads the HTTP request and
// calls the "PolicyTemplate" service "index" endpoint.
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
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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

// MountDeleteHandler configures the mux to serve the "PolicyTemplate" service
// "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePolicyTemplateOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/api/governance/projects/{project_id}/policy_templates/{template_id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "PolicyTemplate" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, dec)
		encodeResponse = EncodeDeleteResponse(enc)
		encodeError    = EncodeDeleteError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "PolicyTemplate")
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
// service PolicyTemplate.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handlePolicyTemplateOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/api/governance/projects/{project_id}/policy_templates/compile", f)
	mux.Handle("OPTIONS", "/api/governance/projects/{project_id}/policy_templates", f)
	mux.Handle("OPTIONS", "/api/governance/projects/{project_id}/policy_templates/{template_id}", f)
	mux.Handle("OPTIONS", "/api/governance/projects/{project_id}/policy_templates/{template_id}/retrieve_data", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handlePolicyTemplateOrigin applies the CORS response headers corresponding
// to the origin for the service PolicyTemplate.
func handlePolicyTemplateOrigin(h http.Handler) http.Handler {
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
