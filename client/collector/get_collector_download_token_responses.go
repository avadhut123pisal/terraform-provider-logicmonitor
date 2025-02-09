// Code generated by go-swagger; DO NOT EDIT.

package collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"terraform-provider-logicmonitor/models"
)

// GetCollectorDownloadTokenReader is a Reader for the GetCollectorDownloadToken structure.
type GetCollectorDownloadTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectorDownloadTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectorDownloadTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCollectorDownloadTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectorDownloadTokenOK creates a GetCollectorDownloadTokenOK with default headers values
func NewGetCollectorDownloadTokenOK() *GetCollectorDownloadTokenOK {
	return &GetCollectorDownloadTokenOK{}
}

/* GetCollectorDownloadTokenOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCollectorDownloadTokenOK struct {
	Payload interface{}
}

func (o *GetCollectorDownloadTokenOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}/downloadToken][%d] getCollectorDownloadTokenOK  %+v", 200, o.Payload)
}
func (o *GetCollectorDownloadTokenOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetCollectorDownloadTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectorDownloadTokenDefault creates a GetCollectorDownloadTokenDefault with default headers values
func NewGetCollectorDownloadTokenDefault(code int) *GetCollectorDownloadTokenDefault {
	return &GetCollectorDownloadTokenDefault{
		_statusCode: code,
	}
}

/* GetCollectorDownloadTokenDefault describes a response with status code -1, with default header values.

Error
*/
type GetCollectorDownloadTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get collector download token default response
func (o *GetCollectorDownloadTokenDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectorDownloadTokenDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}/downloadToken][%d] getCollectorDownloadToken default  %+v", o._statusCode, o.Payload)
}
func (o *GetCollectorDownloadTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCollectorDownloadTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
