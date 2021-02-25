// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PrometheusReader is a Reader for the Prometheus structure.
type PrometheusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusOK creates a PrometheusOK with default headers values
func NewPrometheusOK() *PrometheusOK {
	return &PrometheusOK{}
}

/* PrometheusOK describes a response with status code 200, with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 201.
*/
type PrometheusOK struct {
}

func (o *PrometheusOK) Error() string {
	return fmt.Sprintf("[GET /metrics/prometheus][%d] prometheusOK ", 200)
}

func (o *PrometheusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
