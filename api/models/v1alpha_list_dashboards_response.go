// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1alphaListDashboardsResponse v1alpha list dashboards response
// swagger:model v1alphaListDashboardsResponse
type V1alphaListDashboardsResponse struct {

	// dashboards
	Dashboards []*V1alphaDashboard `json:"dashboards"`

	// next page token
	NextPageToken string `json:"next_page_token,omitempty"`

	// total size
	TotalSize int32 `json:"total_size,omitempty"`
}

// Validate validates this v1alpha list dashboards response
func (m *V1alphaListDashboardsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alphaListDashboardsResponse) validateDashboards(formats strfmt.Registry) error {

	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {
		if swag.IsZero(m.Dashboards[i]) { // not required
			continue
		}

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alphaListDashboardsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alphaListDashboardsResponse) UnmarshalBinary(b []byte) error {
	var res V1alphaListDashboardsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
