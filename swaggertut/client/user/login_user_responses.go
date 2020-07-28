// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoginUserReader is a Reader for the LoginUser structure.
type LoginUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoginUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoginUserOK creates a LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {
	return &LoginUserOK{}
}

/*LoginUserOK handles this case with default header values.

successful operation
*/
type LoginUserOK struct {
	/*date in UTC when token expires
	 */
	XExpiresAfter strfmt.DateTime
	/*calls per hour allowed by the user
	 */
	XRateLimit int32

	Payload string
}

func (o *LoginUserOK) Error() string {
	return fmt.Sprintf("[GET /user/login][%d] loginUserOK  %+v", 200, o.Payload)
}

func (o *LoginUserOK) GetPayload() string {
	return o.Payload
}

func (o *LoginUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Expires-After

	xExpiresAfter, err := formats.Parse("date-time", response.GetHeader("X-Expires-After"))
	if err != nil {
		return errors.InvalidType("X-Expires-After", "header", "strfmt.DateTime", response.GetHeader("X-Expires-After"))
	}
	o.XExpiresAfter = *(xExpiresAfter.(*strfmt.DateTime))

	// response header X-Rate-Limit
	xRateLimit, err := swag.ConvertInt32(response.GetHeader("X-Rate-Limit"))
	if err != nil {
		return errors.InvalidType("X-Rate-Limit", "header", "int32", response.GetHeader("X-Rate-Limit"))
	}
	o.XRateLimit = xRateLimit

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUserBadRequest creates a LoginUserBadRequest with default headers values
func NewLoginUserBadRequest() *LoginUserBadRequest {
	return &LoginUserBadRequest{}
}

/*LoginUserBadRequest handles this case with default header values.

Invalid username/password supplied
*/
type LoginUserBadRequest struct {
}

func (o *LoginUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /user/login][%d] loginUserBadRequest ", 400)
}

func (o *LoginUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
