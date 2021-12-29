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

// MiscGetCollectorDownloadTokenReader is a Reader for the MiscGetCollectorDownloadToken structure.
type MiscGetCollectorDownloadTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MiscGetCollectorDownloadTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMiscGetCollectorDownloadTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMiscGetCollectorDownloadTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMiscGetCollectorDownloadTokenOK creates a MiscGetCollectorDownloadTokenOK with default headers values
func NewMiscGetCollectorDownloadTokenOK() *MiscGetCollectorDownloadTokenOK {
	return &MiscGetCollectorDownloadTokenOK{}
}

/* MiscGetCollectorDownloadTokenOK describes a response with status code 200, with default header values.

successful operation
*/
type MiscGetCollectorDownloadTokenOK struct {
	Payload interface{}
}

func (o *MiscGetCollectorDownloadTokenOK) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}/downloadToken][%d] miscGetCollectorDownloadTokenOK  %+v", 200, o.Payload)
}
func (o *MiscGetCollectorDownloadTokenOK) GetPayload() interface{} {
	return o.Payload
}

func (o *MiscGetCollectorDownloadTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMiscGetCollectorDownloadTokenDefault creates a MiscGetCollectorDownloadTokenDefault with default headers values
func NewMiscGetCollectorDownloadTokenDefault(code int) *MiscGetCollectorDownloadTokenDefault {
	return &MiscGetCollectorDownloadTokenDefault{
		_statusCode: code,
	}
}

/* MiscGetCollectorDownloadTokenDefault describes a response with status code -1, with default header values.

Error
*/
type MiscGetCollectorDownloadTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the misc get collector download token default response
func (o *MiscGetCollectorDownloadTokenDefault) Code() int {
	return o._statusCode
}

func (o *MiscGetCollectorDownloadTokenDefault) Error() string {
	return fmt.Sprintf("[GET /setting/collector/collectors/{id}/downloadToken][%d] miscGetCollectorDownloadToken default  %+v", o._statusCode, o.Payload)
}
func (o *MiscGetCollectorDownloadTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MiscGetCollectorDownloadTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
