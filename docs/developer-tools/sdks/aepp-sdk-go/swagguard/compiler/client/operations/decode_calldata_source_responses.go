// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/v5/swagguard/compiler/models"
)

// DecodeCalldataSourceReader is a Reader for the DecodeCalldataSource structure.
type DecodeCalldataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecodeCalldataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDecodeCalldataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDecodeCalldataSourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDecodeCalldataSourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDecodeCalldataSourceOK creates a DecodeCalldataSourceOK with default headers values
func NewDecodeCalldataSourceOK() *DecodeCalldataSourceOK {
	return &DecodeCalldataSourceOK{}
}

/*DecodeCalldataSourceOK handles this case with default header values.

Binary encoded calldata
*/
type DecodeCalldataSourceOK struct {
	Payload *models.DecodedCalldata
}

func (o *DecodeCalldataSourceOK) Error() string {
	return fmt.Sprintf("[POST /decode-calldata/source][%d] decodeCalldataSourceOK  %+v", 200, o.Payload)
}

func (o *DecodeCalldataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DecodedCalldata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecodeCalldataSourceBadRequest creates a DecodeCalldataSourceBadRequest with default headers values
func NewDecodeCalldataSourceBadRequest() *DecodeCalldataSourceBadRequest {
	return &DecodeCalldataSourceBadRequest{}
}

/*DecodeCalldataSourceBadRequest handles this case with default header values.

Invalid data
*/
type DecodeCalldataSourceBadRequest struct {
	Payload *models.Error
}

func (o *DecodeCalldataSourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /decode-calldata/source][%d] decodeCalldataSourceBadRequest  %+v", 400, o.Payload)
}

func (o *DecodeCalldataSourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecodeCalldataSourceForbidden creates a DecodeCalldataSourceForbidden with default headers values
func NewDecodeCalldataSourceForbidden() *DecodeCalldataSourceForbidden {
	return &DecodeCalldataSourceForbidden{}
}

/*DecodeCalldataSourceForbidden handles this case with default header values.

Invalid data
*/
type DecodeCalldataSourceForbidden struct {
	Payload *models.Error
}

func (o *DecodeCalldataSourceForbidden) Error() string {
	return fmt.Sprintf("[POST /decode-calldata/source][%d] decodeCalldataSourceForbidden  %+v", 403, o.Payload)
}

func (o *DecodeCalldataSourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
