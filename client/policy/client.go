package policy

import (
	"context"
	"net/http"
	"time"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"

	"github.com/rightscale/policy_sdk/auth"
	"github.com/rightscale/policy_sdk/sdk/applied_policy"
	apclient "github.com/rightscale/policy_sdk/sdk/http/applied_policy/client"
	iclient "github.com/rightscale/policy_sdk/sdk/http/incident/client"
	ptclient "github.com/rightscale/policy_sdk/sdk/http/policy_template/client"
	"github.com/rightscale/policy_sdk/sdk/incident"
	"github.com/rightscale/policy_sdk/sdk/policy_template"
)

const apiVersion = "1.0"

type (
	// Client is a thin wrapper around the generated front service clients in the sdk.
	Client interface {
		CreateAppliedPolicy(ctx context.Context, p *appliedpolicy.CreatePayload) (*appliedpolicy.AppliedPolicy, error)
		DeleteAppliedPolicy(ctx context.Context, id string) error
		ShowAppliedPolicy(ctx context.Context, id string, view string) (*appliedpolicy.AppliedPolicy, error)
		IndexAppliedPolicies(ctx context.Context, names []string, view, etag string) (*appliedpolicy.AppliedPolicyList, error)
		ShowAppliedPolicyLog(ctx context.Context, id string, etag string) (*appliedpolicy.AppliedPolicyLog, error)
		ShowAppliedPolicyStatus(ctx context.Context, id string) (*appliedpolicy.AppliedPolicyStatus, error)

		ShowIncident(ctx context.Context, id string, view string) (*incident.Incident, error)
		IndexIncidents(ctx context.Context, appliedPolicyID string, state []string, view string, etag string) (*incident.IncidentList, error)
		IndexEscalations(ctx context.Context, id string) (*incident.Escalations, error)

		CompilePolicyTemplate(ctx context.Context, filename string, source string) error
		UploadPolicyTemplate(ctx context.Context, filename string, source string) (*policytemplate.PolicyTemplate, error)
		UpdatePolicyTemplate(ctx context.Context, id, filename string, source string) (*policytemplate.PolicyTemplate, error)
		ShowPolicyTemplate(ctx context.Context, id string, view string) (*policytemplate.PolicyTemplate, error)
		RetrieveData(ctx context.Context, templateID string, names []string, options []*policytemplate.ConfigurationOptionCreateType, credentials map[string]string) ([]*policytemplate.Data, error)
		DeletePolicyTemplate(ctx context.Context, id string) error

		// CreatePublishedTemplate(ctx context.Context, orgID uint, templateHref string) (*publishedtemplate.CreateResult, error)
		// ShowPublishedTemplate(ctx context.Context, orgID uint, id string, view string) (*publishedtemplate.PublishedTemplate, error)
		// DeletePublishedTemplate(ctx context.Context, orgID uint, id string) error
	}

	client struct {
		host      string
		projectID uint
		ts        auth.TokenSource
		ape       appliedPolicyEndpoints
		ie        incidentEndpoints
		pte       policyTemplateEndpoints
	}

	// appliedPolicyEndpoints implements the applied policy pkg client wrapper
	appliedPolicyEndpoints struct {
		create     goa.Endpoint
		delete     goa.Endpoint
		show       goa.Endpoint
		index      goa.Endpoint
		showLog    goa.Endpoint
		showStatus goa.Endpoint
	}

	// policyTemplateEndpoints implements the applied policy pkg client wrapper
	policyTemplateEndpoints struct {
		compile      goa.Endpoint
		upload       goa.Endpoint
		update       goa.Endpoint
		show         goa.Endpoint
		delete       goa.Endpoint
		retrieveData goa.Endpoint
	}

	// incidentEndpoints implements the applied policy pkg client wrapper
	incidentEndpoints struct {
		show        goa.Endpoint
		index       goa.Endpoint
		escalations goa.Endpoint
	}
)

// NewClient returns a new client for RightScale Policy service.
//   host should be the API host, such as governance-3.rightscale.com
func NewClient(host string, projectID uint, ts auth.TokenSource, debug bool) Client {
	var doer goahttp.Doer = &http.Client{Timeout: time.Duration(300) * time.Second}
	// if debug {
	// 	doer = clientwrappers.LogResponseWrap(doer)
	// }
	// doer = clientwrappers.NewDefaultBackoffDoerV2(doer)
	apc := apclient.NewClient("https", host, doer, goahttp.RequestEncoder, goahttp.ResponseDecoder, false)
	apcCustom := apclient.NewClient("https", host, doer, goahttp.RequestEncoder, showLogResponseDecoder, false)
	ic := iclient.NewClient("https", host, doer, goahttp.RequestEncoder, goahttp.ResponseDecoder, false)
	ptc := ptclient.NewClient("https", host, doer, goahttp.RequestEncoder, goahttp.ResponseDecoder, false)

	return &client{
		host:      host,
		projectID: projectID,
		ts:        ts,
		ape: appliedPolicyEndpoints{
			create:     apc.Create(),
			delete:     apc.Delete(),
			show:       apc.Show(),
			index:      apc.Index(),
			showLog:    apcCustom.ShowLog(),
			showStatus: apc.ShowStatus(),
		},
		pte: policyTemplateEndpoints{
			compile:      ptc.Compile(),
			upload:       ptc.Upload(),
			update:       ptc.Update(),
			delete:       ptc.Delete(),
			show:         ptc.Show(),
			retrieveData: ptc.RetrieveData(),
		},
		ie: incidentEndpoints{
			show:        ic.Show(),
			index:       ic.Index(),
			escalations: ic.IndexEscalations(),
		},
	}
}

// getToken generates an Bearer access token for us
func (c *client) getToken() (*string, error) {
	tok, err := c.ts.TokenString()
	if err != nil {
		return nil, err
	}
	return &tok, nil
}
