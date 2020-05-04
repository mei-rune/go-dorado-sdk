package dorado

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type HyperMetroDomain struct {
	CPSID          string `json:"CPSID"`
	CPSNAME        string `json:"CPSNAME"`
	CPTYPE         string `json:"CPTYPE"`
	DESCRIPTION    string `json:"DESCRIPTION"`
	DOMAINTYPE     string `json:"DOMAINTYPE"`
	ID             string `json:"ID"`
	NAME           string `json:"NAME"`
	REMOTEDEVICES  string `json:"REMOTEDEVICES"`
	RUNNINGSTATUS  string `json:"RUNNINGSTATUS"`
	STANDBYCPSID   string `json:"STANDBYCPSID"`
	STANDBYCPSNAME string `json:"STANDBYCPSNAME"`
	TYPE           int    `json:"TYPE"`
}

const (
	ErrHyperMetroDomainNotFound = "HyperMetroDomain ID is not found"
)

func (c *Client) GetHyperMetroDomains(ctx context.Context, query *SearchQuery) ([]HyperMetroDomain, error) {
	// HyperMetroDomain is a same value between a local device and a remote device.
	return c.LocalDevice.GetHyperMetroDomains(ctx, query)
}

func (d *Device) GetHyperMetroDomains(ctx context.Context, query *SearchQuery) ([]HyperMetroDomain, error) {
	// NOTE(whywaita): implement only GET.
	// HyperMetroDomain is a few under our usage.

	spath := "/HyperMetroDomain"

	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateRequest+": %w", err)
	}
	req = AddSearchQuery(req, query)
	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrHTTPRequestDo+": %w", err)
	}

	var hyperMetroDomains []HyperMetroDomain
	if err = decodeBody(resp, &hyperMetroDomains); err != nil {
		return nil, fmt.Errorf(ErrDecodeBody+": %w", err)
	}

	if len(hyperMetroDomains) == 0 {
		return nil, errors.New(ErrHyperMetroDomainNotFound)
	}

	return hyperMetroDomains, nil
}
