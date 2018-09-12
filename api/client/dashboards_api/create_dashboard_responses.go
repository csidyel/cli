// Code generated by go-swagger; DO NOT EDIT.

package dashboards_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderedtext/sem/api/models"
)

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDashboardOK creates a CreateDashboardOK with default headers values
func NewCreateDashboardOK() *CreateDashboardOK {
	return &CreateDashboardOK{}
}

/*CreateDashboardOK handles this case with default header values.

CreateDashboardOK create dashboard o k
*/
type CreateDashboardOK struct {
	Payload *models.V1alphaDashboard
}

func (o *CreateDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1alpha/dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alphaDashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
