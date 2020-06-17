package dorado

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// PortGroup is group of Port (ex Ethernet, FiberChannel...)
type PortGroup struct {
	DESCRIPTION string `json:"DESCRIPTION"`
	ID          int    `json:"ID,string"`
	NAME        string `json:"NAME"`
	TYPE        int    `json:"TYPE"`
}

// Error const
const (
	ErrPortGroupNotFound = "PortGroup is not found"
)

// GetPortGroups get port groups by query
func (d *Device) GetPortGroups(ctx context.Context, query *SearchQuery) ([]PortGroup, error) {
	spath := "/portgroup"

	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateRequest+": %w", err)
	}
	req = AddSearchQuery(req, query)

	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrHTTPRequestDo+": %w", err)
	}

	portgroups := []PortGroup{}
	if err = decodeBody(resp, &portgroups); err != nil {
		return nil, fmt.Errorf(ErrDecodeBody+": %w", err)
	}

	if len(portgroups) == 0 {
		return nil, errors.New(ErrPortGroupNotFound)
	}

	return portgroups, nil
}

// GetPortGroup get port group by id
func (d *Device) GetPortGroup(ctx context.Context, portgroupID int) (*PortGroup, error) {
	spath := fmt.Sprintf("/portgroup/%d", portgroupID)

	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateRequest+": %w", err)
	}
	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrHTTPRequestDo+": %w", err)
	}

	portgroup := &PortGroup{}
	if err = decodeBody(resp, portgroup); err != nil {
		return nil, fmt.Errorf(ErrDecodeBody+": %w", err)
	}

	return portgroup, nil
}

// GetPortGroupsAssociate get port group that associated by mapping view id
func (d *Device) GetPortGroupsAssociate(ctx context.Context, mappingviewID int) ([]PortGroup, error) {
	spath := "/portgroup/associate"

	req, err := d.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, fmt.Errorf(ErrCreateRequest+": %w", err)
	}
	param := &AssociateParam{
		ASSOCIATEOBJID:   strconv.Itoa(mappingviewID),
		ASSOCIATEOBJTYPE: TypeMappingView,
	}
	req = AddAssociateParam(req, param)
	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrHTTPRequestDo+": %w", err)
	}

	portgroups := []PortGroup{}
	if err = decodeBody(resp, &portgroups); err != nil {
		return nil, fmt.Errorf(ErrDecodeBody+": %w", err)
	}

	return portgroups, nil
}

// IsAddToMappingViewPortGroup check to associated mapping view
func (d *Device) IsAddToMappingViewPortGroup(ctx context.Context, mappingViewID, portgroupID int) (bool, error) {
	portgroups, err := d.GetPortGroupsAssociate(ctx, mappingViewID)
	if err != nil {
		return false, fmt.Errorf("failed to get portgroups: %w", err)
	}

	for _, p := range portgroups {
		if p.ID == portgroupID {
			return true, nil
		}
	}

	return false, nil
}
