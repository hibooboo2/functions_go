package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// DeleteCallsCallLogReader is a Reader for the DeleteCallsCallLog structure.
type DeleteCallsCallLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCallsCallLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteCallsCallLogAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteCallsCallLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteCallsCallLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCallsCallLogAccepted creates a DeleteCallsCallLogAccepted with default headers values
func NewDeleteCallsCallLogAccepted() *DeleteCallsCallLogAccepted {
	return &DeleteCallsCallLogAccepted{}
}

/*DeleteCallsCallLogAccepted handles this case with default header values.

Log delete request accepted
*/
type DeleteCallsCallLogAccepted struct {
}

func (o *DeleteCallsCallLogAccepted) Error() string {
	return fmt.Sprintf("[DELETE /calls/{call}/log][%d] deleteCallsCallLogAccepted ", 202)
}

func (o *DeleteCallsCallLogAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCallsCallLogNotFound creates a DeleteCallsCallLogNotFound with default headers values
func NewDeleteCallsCallLogNotFound() *DeleteCallsCallLogNotFound {
	return &DeleteCallsCallLogNotFound{}
}

/*DeleteCallsCallLogNotFound handles this case with default header values.

Does not exist.
*/
type DeleteCallsCallLogNotFound struct {
	Payload *models.Error
}

func (o *DeleteCallsCallLogNotFound) Error() string {
	return fmt.Sprintf("[DELETE /calls/{call}/log][%d] deleteCallsCallLogNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCallsCallLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCallsCallLogDefault creates a DeleteCallsCallLogDefault with default headers values
func NewDeleteCallsCallLogDefault(code int) *DeleteCallsCallLogDefault {
	return &DeleteCallsCallLogDefault{
		_statusCode: code,
	}
}

/*DeleteCallsCallLogDefault handles this case with default header values.

Unexpected error
*/
type DeleteCallsCallLogDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete calls call log default response
func (o *DeleteCallsCallLogDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCallsCallLogDefault) Error() string {
	return fmt.Sprintf("[DELETE /calls/{call}/log][%d] DeleteCallsCallLog default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCallsCallLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
